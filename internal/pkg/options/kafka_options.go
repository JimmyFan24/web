package options

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
)

type KafkaOptions struct {
	Address string
	Topic string
	MsgChan chan *sarama.ProducerMessage
	Client sarama.SyncProducer
}

func NewKafkaOptions () *KafkaOptions{
	logrus.Info("4.创建新的Options，这里是kafka")
	return &KafkaOptions{
		Address: "",
		Topic:"",
		MsgChan: make(chan *sarama.ProducerMessage),
	}
}
func(o *KafkaOptions) AddFlags(fs *pflag.FlagSet){
	fs.StringVar(&o.Address,"kafka-address","","kafka connect address")
	fs.StringVar(&o.Topic,"kafka-topic","","kafka topic")

}
/*func (o *KafkaOptions)NewKafkaClient ()error{
	address := o.Address
	cli,err :=  db.NewKafkaClient(address)
	if err != nil {
		return err
	}
	o.Client = cli
	return nil
}*/