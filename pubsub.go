package main

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/pubsub"
)

func publishMessage(topicId string, transactionScore any, attributes map[string]string) (error, error) {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, "impactful-shard-374913")
	if err != nil {
		return fmt.Errorf("pubsub.NewClient: %v", err), nil
	}
	defer client.Close()
	message, _ := json.Marshal(transactionScore)

	topic := client.Topic(topicId)
	result := topic.Publish(ctx, &pubsub.Message{
		Data:       message,
		Attributes: attributes,
	})

	_, err = result.Get(ctx)
	if err != nil {
		return fmt.Errorf("error publishing to Topic %s: %v", topicId, err), nil
	}
	return nil, nil
}
