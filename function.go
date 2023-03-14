package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// Set up a Google Cloud Pub/Sub client
	// ctx := context.Background()
	// client, err := pubsub.NewClient(ctx, "impactful-shard-374913")
	// if err != nil {
	// 	fmt.Println("Failed to create client: ", err)
	// 	return
	// }
	// // Set up a subscription to the topic
	// sub := client.Subscription("SubForAnomalyDetector")
	// fmt.Println("Listening ...")

	// err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
	// 	fmt.Println(">>> Received message: ", string(msg.Data))
	// 	eventType := msg.Attributes["pigeon.eventType"]

	// 	if eventType != "walli.TransactionUpdatedEventV1" {
	// 		log.Println("Missing or inccorrect attribute value for pigeon.eventType")
	// 		msg.Ack()
	// 		return
	// 	}

	// 	msg.Ack()
	// 	var transaction Transaction
	// 	_ = json.Unmarshal(msg.Data, &transaction)

	// 	user := getUser(ctx, transaction.UserID)

	// 	log.Println("Sending to AnomalyDetector:")
	// 	log.Println(fmt.Sprintf("\tTransaction uuid: %s\n\tAmount: %d\n\tEmail: %s\n\tJob: %s", transaction.Uuid, transaction.TransactionAmount.Amount, user.Email, user.Job))

	// 	predict, error := calculateAnomalyScore(ctx, transaction, user.Job, user.Email)

	// 	if error != nil {
	// 		log.Panic(fmt.Printf("Error in anomaly detector\n\terror: %v\n", error))
	// 		return
	// 	} else {

	// 		log.Println(fmt.Printf("Step 2 - Prediction results for: { Amount: %d, Email: %s, Job:%s }", transaction.TransactionAmount.Amount, user.Email, user.Job))

	// 		for key, element := range predict.ModelScores[0].Scores {
	// 			fmt.Println("\t(", key, " => ", element, ")")
	// 			score := element
	// 			label := getLabel(score)

	// 			transactionScore := struct {
	// 				TransactionID     string
	// 				UserID            string
	// 				ExternalAccountID string
	// 				Label             string
	// 				Score             float32
	// 			}{
	// 				TransactionID:     transaction.Uuid,
	// 				UserID:            transaction.UserID,
	// 				ExternalAccountID: transaction.ExternalAccountID,
	// 				Label:             label,
	// 				Score:             score,
	// 			}

	// 			attributes := map[string]string{"pigeon.eventType": "walli.TransactionAnomalyScoreCalculatedEventV1"}

	// 			publishMessage("TransactionScores", transactionScore, attributes)
	// 			log.Println("Step 3 - Published Message")
	// 		}
	// 	}
	// })

	// fmt.Println("Listening ...")
	// if err != nil {
	// 	fmt.Println("Failed to start subscription: ", err)
	// 	return
	// }

	// fmt.Println("Listening ...")
}

type PubSubMessage struct {
	Data       []byte            `json:"data"`
	Attributes map[string]string `json:"attributes"`
}

func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	var transaction Transaction
	_ = json.Unmarshal(m.Data, &transaction)
	log.Println(fmt.Sprintf("Received: %s", m.Data))

	eventType := m.Attributes["pigeon.eventType"]

	if eventType != "walli.TransactionUpdatedEventV1" {
		return fmt.Errorf("Missing or inccorrect attribute value for pigeon.eventType %s", eventType)
	}

	user := getUser(ctx, transaction.UserID)
	log.Println("Got user: ", user)

	predict, error := calculateAnomalyScore(ctx, transaction, user.Job, user.Email)

	if error != nil {
		log.Panic("Error in anomaly detector")
		return error
	} else {
		for key, element := range predict.ModelScores[0].Scores {
			fmt.Println("Key:", key, "=>", "Element:", element)
		}
	}

	return nil
}

func getLabel(score float32) string {
	if score >= 800 {
		return "High risk"
	}
	if score >= 400 && score < 800 {
		return "Unusual"
	}
	return "Legit"
}
