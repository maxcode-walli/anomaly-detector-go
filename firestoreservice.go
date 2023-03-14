package main

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

func getUser(ctx context.Context, userId string) (a AccountUser) {
	userData := AccountUser{
		Uid:           userId,
		EmailVerified: false,
	}
	app, err := firebase.NewApp(ctx, &firebase.Config{ProjectID: "impactful-shard-374913"})
	if err != nil {
		log.Panic("Error creating firebase app")
		return userData
	}
	client, err = app.Firestore(ctx)
	if err != nil {
		log.Panic("Error initializing firestore client")
		return userData
	} else {
		var user, e = client.Doc("users/" + userId).Get(ctx)
		if e != nil {
			log.Panic("cannot read firestore user")
		}

		m := user.Data()
		log.Println(fmt.Printf("Have read: \tuser.data():%#v", m))

		err := client.RunTransaction(ctx, func(ctx context.Context, transaction *firestore.Transaction) error {
			ref := client.Collection("users").Doc(userId)
			us, err := transaction.Get(ref)

			if err != nil {
				log.Panic(fmt.Printf("User with ref %s not found", userId))
				return nil
				//try reading without transaction : us, err := client.Collection("users").Doc(userId).Get(ctx)

			} else {
				var dbUser AccountUser
				us.DataTo(&dbUser)
				userData.Email = dbUser.Email
				userData.Job = dbUser.Job

				return nil
			}
		})
		if err != nil {
			log.Panic("Error in updating user data from Firestore")
		} else {
			log.Println(fmt.Printf("Success updating user data:\t[email=%s]\tjob=%s]", userData.Email, userData.Job))
		}

		return userData
	}
}
