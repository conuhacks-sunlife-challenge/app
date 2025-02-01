package Database

import (
    "context"
    "fmt"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseInstance struct {
    client *mongo.Client
}

func Connect(user, pass string) DatabaseInstance {
    opts := options.Client().ApplyURI("mongodb+srv://cluster0.dexa4.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetAuth(options.Credential{
        AuthMechanism: "SCRAM-SHA-1", // or "SCRAM-SHA-1"
        AuthSource:    "admin",         // the database where the user is defined
        Username:      "user1",
        Password:      "user1password",
    })

    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
        panic(err)
    }

    return DatabaseInstance{
        client: client,
    }
}

func (db DatabaseInstance) Disconnect() {
    db.Disconnect()
}

func (db DatabaseInstance) Ping() {
    if err := db.client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
        panic(err)
    }

    fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

