package server

import(
	"context"
	"errors"
	"fmt"
	"github.com/gogo/protobuf/version"
	"github.com/prometheus/common/log"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/zsais/go-gin-prometheus"
)

// GenericLogAgentServer contains state for a iam api server.
// type GenericLogAgentServer gin.Engine.

type GenericLogAgentServer struct {
	middlewares[] string
	mode string
	SecureServingInfo *SecureServingInfo
	InsecureServingInfo *InsecureServingInfo
	// ShutdownTimeout is the timeout used for server shutdown. This specifies the timeout before server
	// gracefully shutdown returns.
	ShutdownTimeout time.Duration

	*gin.Engine

	healthz bool
	enableMetrics bool
	enableProfiling bool
	insecureserver,secureserver *http.Server

}

func initGenericLogAgentServer(s *GenericLogAgentServer){
	s.SetUp()
	s.InstallAPIs()
}

// Setup do some setup work for gin engine.

func (s *GenericLogAgentServer)SetUp(){
	gin.SetMode(s.mode)
	gin.DebugPrintRouteFunc = func(httpMethod,absolutePath, handlerName string, nuHandlers int) {
		logrus.Infof("%-6s %-s --> %s (%d handlers)", httpMethod, absolutePath, handlerName, nuHandlers)
	}

}

func (s *GenericLogAgentServer)InstallAPIs(){
	// install healthz handler
	if s.healthz {
		s.GET("/healthz", func(context *gin.Context) {

			context.JSON(http.StatusOK, map[string]string{"status": "ok"})
		})
	}

	// install metric handler
	if s.enableMetrics {
		prometheus := ginprometheus.NewPrometheus("gin")
		prometheus.Use(s.Engine)
	}
	s.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK,version.Get())
	})
}

// Run spawns the http server. It only returns when the port cannot be listened on initially.

func (s *GenericLogAgentServer)Run()error{
	// For scalability, use custom HTTP configuration mode here
	s.insecureserver = &http.Server{
		Addr:s.InsecureServingInfo.Address,
		Handler: s,
	}

	s.secureserver =&http.Server{
		Addr: s.SecureServingInfo.Address(),
		Handler: s,
	}

	var eg errgroup.Group

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below

	eg.Go(func() error {
		logrus.Infof("Start to listening the incoming requests on http address: %s", s.InsecureServingInfo.Address)
			if err :=s.insecureserver.ListenAndServe();err != nil && !errors.Is(err,http.ErrServerClosed){
				logrus.Fatal(err.Error())
				return err
			}
		logrus.Infof("Server on %s stopped", s.InsecureServingInfo.Address)
		return nil
	})

	eg.Go(func() error {
		key, cert := s.SecureServingInfo.CertKey.KeyFile, s.SecureServingInfo.CertKey.CertFile
		if cert == "" || key == "" || s.SecureServingInfo.BindPort == 0 {
			return nil
		}

		logrus.Infof("Start to listening the incoming requests on https address: %s", s.SecureServingInfo.Address())

		if err := s.secureserver.ListenAndServeTLS(cert, key); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logrus.Fatal(err.Error())

			return err
		}

		logrus.Infof("Server on %s stopped", s.SecureServingInfo.Address())

		return nil
	})
	// Ping the server to make sure the router is working.
	ctx,cancel :=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	if s.healthz {
		if err := s.ping(ctx); err != nil {
			return err
		}
	}

	if err := eg.Wait(); err != nil {
		logrus.Fatal(err.Error())
	}

	return nil
}
// ping pings the http server to make sure the router is working.

func (s *GenericLogAgentServer)ping (ctx context.Context)error{
	url := fmt.Sprintf("http://%s/healthz",s.InsecureServingInfo.Address)
	if strings.Contains(s.InsecureServingInfo.Address, "0.0.0.0") {
		url = fmt.Sprintf("http://127.0.0.1:%s/healthz", strings.Split(s.InsecureServingInfo.Address, ":")[1])
	}
	for  {
		req,err:= http.NewRequestWithContext(ctx,http.MethodGet,url,nil)

		if err != nil {
			return err
		}
		// Ping the server by sending a GET request to `/healthz`.
		// nolint: gosec
		resp, err := http.DefaultClient.Do(req)
		if err == nil && resp.StatusCode == http.StatusOK {
			logrus.Info("The router has been deployed successfully.")

			resp.Body.Close()

			return nil
		}

		// Sleep for a second to continue the next ping.
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(1 * time.Second)

		select {
		case <-ctx.Done():
			log.Fatal("can not ping http server within the specified time interval.")
		default:
		}
	}
}

// Close graceful shutdown the api server.
func (s *GenericLogAgentServer) Close() {
	// The context is used to inform the server it has 10 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := s.secureserver.Shutdown(ctx); err != nil {
		log.Warnf("Shutdown secure server failed: %s", err.Error())
	}

	if err := s.insecureserver.Shutdown(ctx); err != nil {
		log.Warnf("Shutdown insecure server failed: %s", err.Error())
	}
}