package data

import (
	"encoding/json"
	"strings"

	"github.com/Shopify/sarama"

	"economy_system/internal/biz"
	"economy_system/internal/conf"
)

type KafkaProducerImpl struct {
	client sarama.SyncProducer
	topic  []string
}

func NewKafkaProducer(c *conf.Client) biz.KafkaProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	addrs := strings.Split(c.Kafka.Addr, ",")
	client, err := sarama.NewSyncProducer(addrs, config)
	if err != nil {
		panic(err)
	}
	topics := strings.Split(c.Kafka.Topic, ",")
	return &KafkaProducerImpl{client: client, topic: topics}
}

func (p *KafkaProducerImpl) SendMsg(content interface{}) error {
	msgList, err := p.buildMsgList(content)
	if err != nil {
		return err
	}
	err = p.client.SendMessages(msgList)
	if err != nil {
		return err
	}
	return nil
}

func (p *KafkaProducerImpl) buildMsgList(content interface{}) ([]*sarama.ProducerMessage, error) {
	data, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}

	msgList := make([]*sarama.ProducerMessage, 0, len(p.topic))
	for _, t := range p.topic {
		msg := &sarama.ProducerMessage{
			Topic: t,
			Value: sarama.ByteEncoder(data),
		}
		msgList = append(msgList, msg)
	}

	return msgList, nil
}
