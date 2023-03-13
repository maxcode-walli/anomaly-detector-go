package p

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go/v4"
)

var (
	client *firestore.Client
)

func getUser(ctx context.Context, userId string) User {
	var user User
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "impactful-shard-374913"})
	if err != nil {
		log.Panic("Error creating firebase app")
		return user
	}

	client, err = app.Firestore(ctx)
	if err != nil {
		log.Panic("Error initializing firestore client")
		return user
	} else {
		err := client.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {

			account := client.Collection("users").Where("uid", "==", userId).Documents(ctx)

			accountDoc, err := account.Next()
			if err != nil {
				log.Panic(fmt.Sprintf("no user found with Uid: %s", userId))
				return err
			}

			var userData User
			if err := accountDoc.DataTo(&userData); err != nil {
				log.Panic("could not cast firestore document to User struct")
				return err
			}

			user.Email = userData.Email
			user.Job = userData.Job

			return nil
		})
		if err != nil {
			log.Panic("Error transaction - read user from db")
			return user
		}
	}
	return user
}
