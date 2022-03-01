package server

import (
	"github.com/gin-gonic/gin"
	"net"
	"strconv"
)



// Config is a structure used to configure a GenericAPIServer.
// Its members are sorted roughly in order of importance for composers.
// SecureServingInfo holds configuration of the TLS server.

type Config struct {
	SecureServing   *SecureServingInfo
	InsecureServing *InsecureServingInfo
	//Jwt             *JwtInfo
	Mode            string
	Middlewares     []string
	Healthz         bool
	EnableProfiling bool
	EnableMetrics   bool
}

// NewConfig returns a Config struct with the default values.
func NewConfig() *Config {
	return &Config{
		Healthz:         true,
		Mode:            gin.ReleaseMode,
		Middlewares:     []string{},
		EnableProfiling: true,
		EnableMetrics:   true,

	}
}

type SecureServingInfo struct {
	BindAddress string
	BindPort int
	CertKey CertKey
}
// Address join host IP address and host port number into a address string, like: 0.0.0.0:8443.
func (s *SecureServingInfo) Address() string {
	return net.JoinHostPort(s.BindAddress, strconv.Itoa(s.BindPort))
}

// CertKey contains configuration items related to certificate.
type CertKey struct {
	// CertFile is a file containing a PEM-encoded certificate, and possibly the complete certificate chain
	CertFile string
	// KeyFile is a file containing a PEM-encoded private key for the certificate specified by CertFile
	KeyFile string
}
type InsecureServingInfo struct {
	Address string
}
// CompletedConfig is the completed configuration for GenericAPIServer.
type CompletedConfig struct {
	*Config
}
// Complete fills in any fields not set that are required to have valid data and can be derived
// from other fields. If you're going to `ApplyOptions`, do that first. It's mutating the receiver.
func (c *Config) Complete() CompletedConfig {
	return CompletedConfig{c}
}
// New returns a new instance of GenericAPIServer from the given config.
func (c CompletedConfig) New() (*GenericLogAgentServer, error) {
	s := &GenericLogAgentServer{
		SecureServingInfo:   c.SecureServing,
		InsecureServingInfo: c.InsecureServing,
		mode:                c.Mode,
		healthz:             c.Healthz,
		enableMetrics:       c.EnableMetrics,
		enableProfiling:     c.EnableProfiling,
		middlewares:         c.Middlewares,
		Engine:              gin.New(),
	}

	initGenericLogAgentServer(s)

	return s, nil
}