package Database

import (
    "context"
    "fmt"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func Connect(user, pass string) DatabaseInstance {
    opts := options.Client().ApplyURI("mongodb+srv://cluster0.dexa4.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0").SetAuth(options.Credential{
        AuthMechanism: "SCRAM-SHA-1", // or "SCRAM-SHA-1"
        AuthSource:    "admin",         // the database where the user is defined
        Username:      user,
        Password:      pass,
    })
    ctx := context.TODO()

    client, err := mongo.Connect(context.TODO(), opts)
    if err != nil {
        panic(err)
    }

    prod := client.Database("snappy_production")

    return DatabaseInstance{
        client: client,
        production: prod,
        ctx: ctx,
    }
}

func (db DatabaseInstance) Disconnect() {
    db.Disconnect()
}

func (db DatabaseInstance) ForceDropRestart() {
    db.production.Drop(db.ctx)
    db.production.CreateCollection(db.ctx, "users")
}

func (db DatabaseInstance) Ping() {
    if err := db.production.RunCommand(db.ctx, bson.D{{"ping", 1}}).Err(); err != nil {
        panic(err)
    }

    fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}

