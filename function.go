// Package p contains a Pub/Sub Cloud Function.
package p

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	frauddetector "github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
)

// PubSubMessage is the payload of a Pub/Sub event. Please refer to the docs for
// additional information regarding Pub/Sub events.
type PubSubMessage struct {
	Data []byte `json:"data"`
}

// HelloPubSub consumes a Pub/Sub message.
func HelloPubSub(ctx context.Context, m PubSubMessage) error {
	var transaction Transaction
	_ = json.Unmarshal(m.Data, &transaction)
	log.Println(fmt.Sprintf("Received: %s", m.Data))

	//Step 1 - Get User
	//Step 2 - Calculate anomaly score
	//Step 3 - Send score in topic TransactionScores

	user := getUser(ctx, transaction.UserID)
	log.Println("Step1 - Get user:")
	log.Println(user)

	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion("eu-west-1"))

	if err != nil {
		log.Panic("ERROR config")
	}

	client := frauddetector.NewFromConfig(cfg)

	entityId := "anomaly_detector_1"
	entityType := "anomaly_detector_1"
	detectorId := "anomaly_detector"
	timestamp := "2023-03-13T11:00:00Z"
	detectorVersionId := "1"

	entities := []types.Entity{{
		EntityId:   &entityId,
		EntityType: &entityType,
	}}
	eventVariables := map[string]string{
		"amount":             strconv.Itoa(int(transaction.TransactionAmount.Amount)),
		"job":                user.Job,
		"user_email_address": user.Email,
	}

	eventInput := frauddetector.GetEventPredictionInput{
		DetectorId:                     &detectorId,
		Entities:                       entities,
		EventId:                        &transaction.Uuid,
		EventTimestamp:                 &timestamp,
		EventTypeName:                  &entityType,
		EventVariables:                 eventVariables,
		DetectorVersionId:              &detectorVersionId,
		ExternalModelEndpointDataBlobs: map[string]types.ModelEndpointDataBlob{},
	}

	pred, err := client.GetEventPrediction(ctx, &eventInput)

	if err != nil {
		return err
	}

	log.Println("Step2 - Get Prediction:")

	for key, element := range pred.ModelScores[0].Scores {
		fmt.Println("Key:", key, "=>", "Element:", element)
	}

	log.Println("Step3 - Publish Message: (TODO)")

	return nil
}
