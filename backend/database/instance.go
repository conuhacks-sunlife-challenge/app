package Database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type DatabaseInstance struct {
    client *mongo.Client
    production *mongo.Database
    ctx context.Context
}

