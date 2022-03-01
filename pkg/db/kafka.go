package db

import (
	"github.com/Shopify/sarama"
	"github.com/sirupsen/logrus"
)

func NewKafkaClient(addr string)( sarama.SyncProducer,error){
	logrus.Info("创建kafka连接....")
	config :=sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	//新选举出一个分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	//成功交付的信息在success channel里面返回
	config.Producer.Return.Successes = true
	//Address := o.Address
	cli,err := sarama.NewSyncProducer([]string{addr},config)
	if err!= nil{
		logrus.Error("kafka producer closed,err:%v",err)
		return nil,err
	}
	logrus.Info("kafka connect success...")
	return cli,nil

}
