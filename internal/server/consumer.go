package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/Shopify/sarama"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"

	"economy_system/internal/conf"
)

func (cs *ConsumerServer) Setup(session sarama.ConsumerGroupSession) error {
	close(cs.ready)
	return nil
}

func (cs *ConsumerServer) Cleanup(session sarama.ConsumerGroupSession) error {
	session.Commit()
	return nil
}

func (cs *ConsumerServer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		ctx := session.Context()
		err := cs.handle(ctx, message)
		if err != nil {
			cs.log.WithContext(ctx).Errorf("handle message err: %+v", err)
			session.ResetOffset(claim.Topic(), claim.Partition(), claim.HighWaterMarkOffset(), "")
		} else {
			session.Commit()
		}
	}

	return nil
}

type ConsumerServer struct {
	client  sarama.ConsumerGroup
	topic   string
	groupID string
	handle  func(ctx context.Context, message *sarama.ConsumerMessage) error
	log     *log.Helper
	ready   chan bool
}

func NewConsumerServer(c *conf.Server, logger log.Logger) *ConsumerServer {
	defaultHandle := func(ctx context.Context, message *sarama.ConsumerMessage) error {
		fmt.Printf("Partition: %d, Offset: %d, key: %s, value: %v",
			message.Partition, message.Offset, string(message.Key), string(message.Value))
		return nil
	}
	config := sarama.NewConfig()
	version, err := sarama.ParseKafkaVersion(c.Kafka.Version)
	if err != nil {
		panic(err)
	}
	config.Version = version
	config.Consumer.Offsets.Initial = sarama.OffsetNewest

	switch c.Kafka.Assignor {
	case "sticky":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		panic(errors.Newf(400, "UnSupport_Assignor", "UnSupport assignor %s", c.Kafka.Assignor))
	}

	if c.Kafka.Oldest {
		config.Consumer.Offsets.Initial = sarama.OffsetOldest
	}

	addrs := strings.Split(c.Kafka.Addr, ",")
	client, err := sarama.NewConsumerGroup(addrs, c.Kafka.Group, config)
	if err != nil {
		panic(err)
	}
	return &ConsumerServer{
		client:  client,
		topic:   c.Kafka.Topic,
		groupID: c.Kafka.Group,
		handle:  defaultHandle,
		log:     log.NewHelper(logger),
	}
}

func (cs *ConsumerServer) RegisterHandler(handle func(ctx context.Context, message *sarama.ConsumerMessage) error) {
	cs.handle = handle
}

func (cs *ConsumerServer) Start(ctx context.Context) error {
	iCtx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Wait()
		for {
			if err := cs.client.Consume(iCtx, strings.Split(cs.topic, ","), cs); err != nil {
				panic(err)
			}
			if iCtx.Err() != nil {
				return
			}
			cs.ready = make(chan bool)
		}
	}()

	<-cs.ready // Await till the consumer has been set up
	cs.log.WithContext(iCtx).Info("Sarama consumer up and running!...")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-iCtx.Done():
		cs.log.WithContext(iCtx).Info("terminating: context cancelled")
	case <-sigterm:
		cs.log.WithContext(iCtx).Info("terminating: via signal")
	}
	cancel()
	wg.Wait()
	if err := cs.client.Close(); err != nil {
		panic(err)
	}
	return nil
}

func (cs *ConsumerServer) Stop(ctx context.Context) error {
	err := cs.client.Close()
	if err != nil {
		cs.log.WithContext(ctx).Errorf("Stop consumer err: %+v", err)
	}
	return err
}
