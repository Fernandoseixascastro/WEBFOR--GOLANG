package db

import "context"

func Insert(collection string, date any) (err error) {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database("Webform").Collection(collection)

	_, err = c.InsertOne(context.Background(), date)

	return err
}