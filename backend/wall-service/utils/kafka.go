package utils

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
	"google.golang.org/protobuf/proto"

	stats "messenger/stats-service/stats"
)

const (
	TopicCommentCreated = "comment_created"
	TopicPostViewed     = "post_viewed"
	TopicPostLiked      = "post_liked"
)

func NewPublisher(cfg *TKafkaConfig) (*TPublisher, error) {
	writers := make(map[string]*kafka.Writer)

	for _, topic := range []string{
		TopicCommentCreated,
		TopicPostViewed,
		TopicPostLiked,
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

func (p *TPublisher) PublishCommentCreated(ev *stats.PostCommented) error {
	msg, err := proto.Marshal(ev)
	if err != nil {
		return fmt.Errorf("marshal CommentCreated: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()
	return p.writers[TopicCommentCreated].WriteMessages(ctx, kafka.Message{Value: msg})
}

func (p *TPublisher) PublishPostViewed(ev *stats.PostViewed) error {
	msg, err := proto.Marshal(ev)
	if err != nil {
		return fmt.Errorf("marshal PostViewed: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()
	return p.writers[TopicPostViewed].WriteMessages(ctx, kafka.Message{Value: msg})
}

func (p *TPublisher) PublishPostLiked(ev *stats.PostLiked) error {
	msg, err := proto.Marshal(ev)
	if err != nil {
		return fmt.Errorf("marshal PostLiked: %w", err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), p.timeout)
	defer cancel()
	return p.writers[TopicPostLiked].WriteMessages(ctx, kafka.Message{Value: msg})
}
