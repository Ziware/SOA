package utils

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	stats "messenger/stats-service/stats"
)

const (
	TopicUserCreated = "user_created"
)

func NewPublisher(cfg *TKafkaConfig) (*TPublisher, error) {
	writers := make(map[string]*kafka.Writer)

	for _, topic := range []string{
		TopicUserCreated,
	} {
		writers[topic] = &kafka.Writer{
			Addr:     kafka.TCP(cfg.Brokers...),
			Topic:    topic,
			Balancer: &kafka.Hash{},
		}
	}

	return &TPublisher{
		writers: writers,
		timeout: cfg.Timeout,
	}, nil
}

func (p *TPublisher) Close() error {
	var firstErr error
	for _, w := range p.writers {
		if err := w.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (p *TPublisher) PublishUserCreated(ev *stats.UserCreated) error {
	msg, err := proto.Marshal(ev)
	if err != nil {
		return fmt.Errorf("marshal UserCreated: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()
	return p.writers[TopicUserCreated].WriteMessages(ctx, kafka.Message{Value: msg})
}
