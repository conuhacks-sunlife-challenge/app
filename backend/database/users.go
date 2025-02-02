package Database

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

type User struct {
    Email string `bson:"email"`
    Password string `bson:"password"`
    FirstName string `bson:"first_name"`
    LastName string `bson:"last_name"`
}


// TODO: Change input to be one object
func (db DatabaseInstance) AddUser(email, password, firstname, lastname string) (bool, error) {

    usersCollection := db.production.Collection("users")
    duplicateUser, err := db.CheckUser(email)

    if err != nil {
        return false, err
    }

    if duplicateUser != nil {
        return false, fmt.Errorf("User already exists! Please check for existing user before adding. Email: %s", duplicateUser.Email)
    }

    user := User{
        Email: email,
        Password: password,
        FirstName: firstname,
        LastName: lastname,
    }

    usersCollection.InsertOne(db.ctx, user)


    return true, nil
}

func (db DatabaseInstance) CheckUser(email string) (*User, error) {
    usersCollection := db.production.Collection("users")

    filter := bson.D{{"email", email}}
    cursor, err := usersCollection.Find(db.ctx, filter)
    if err != nil {
        return nil, err
    }

    var results []User
    err = cursor.All(db.ctx, &results)
    if err != nil {
        return nil, err
    }

    if len(results) > 1 {
        return nil, errors.New("Duplicate user with the same email!")
    }
    if len(results) > 0 {
        return &results[0], nil
    }

    return nil, nil
}
