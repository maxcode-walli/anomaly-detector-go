package main

import (
	"context"
	"log"
	"strconv"

	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector/types"
)

func calculateAnomalyScore(ctx context.Context, transaction Transaction, job string, email string) (p *frauddetector.GetEventPredictionOutput, e error) {
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion("eu-west-1"))

	if err != nil {
		log.Panic("ERROR reading aws config")
	}

	client := frauddetector.NewFromConfig(cfg)

	entityId := "anomaly_detector_1"
	entityType := "anomaly_detector_1"
	detectorId := "anomaly_detector"
	timestamp := transaction.BookingDate
	if len(timestamp) <= 10 {
		timestamp = timestamp + "T13:54:33Z"
	}
	detectorVersionId := "1"

	entities := []types.Entity{{
		EntityId:   &entityId,
		EntityType: &entityType,
	}}
	eventVariables := map[string]string{
		"amount":             strconv.Itoa(int(transaction.TransactionAmount.Amount)),
		"job":                job,
		"user_email_address": email,
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
		log.Println("error GetEventPrediction")
		return nil, err
	}
	return pred, nil
}
