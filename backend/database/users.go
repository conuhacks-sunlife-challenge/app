package Database

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
    Email string `bson:"_id"`
    Password string `bson:"password"`
    FirstName string `bson:"first_name"`
    LastName string `bson:"last_name"`
    RecentFailedLoginAttempts string `bson:"recent_failed_login_attempts"`
}


// TODO: Change input to be one object
func (db DatabaseInstance) AddUser(user User) error {

    usersCollection := db.production.Collection("users")
    duplicateUser, err := db.CheckUser(user.Email)

    if err != nil {
        return err
    }

    if duplicateUser != nil {
        return fmt.Errorf("User already exists! Please check for existing user before adding. Email: %s", duplicateUser.Email)
    }

    usersCollection.InsertOne(db.ctx, user)


    return nil
}

func (db DatabaseInstance) CheckUser(email string) (*User, error) {
    usersCollection := db.production.Collection("users")

    filter := bson.D{{"_id", email}}
    var user User
    err := usersCollection.FindOne(db.ctx, filter).Decode(&user)

    if err != nil {
        if err == mongo.ErrNoDocuments {
            return nil, nil
        }
        return nil, fmt.Errorf("error finding user: %v", err)
    }

    return &user, nil
}

// TODO: make secure
func (db DatabaseInstance) Authenticate(email, password string) (bool, error) {

    usersCollection := db.production.Collection("users")

    filter := bson.D{{"email", email}}
    cursor, err := usersCollection.Find(db.ctx, filter)
    if err != nil {
        return false, err
    }

    var results []User
    err = cursor.All(db.ctx, &results)
    if err != nil {
        return false, err
    }

    if len(results) > 1 {
        return false, errors.New("Duplicate user with the same email!")
    }

    if len(results) == 0 {
        return false, nil
    }

    user := results[0]
    if user.Password == password {
        return true, nil
    }

    return false, nil
}

