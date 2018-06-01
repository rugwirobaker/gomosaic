package engine

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

//App ...
type App struct{}

//init initializes google app engine appliation instance.
// it is the entry point of the application.
func init() {}

func initFirebase() *firebase.App {
	opt := option.WithCredentialsFile("path/to/serviceAccountKey.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	return app
}

func initFirestore() {}
