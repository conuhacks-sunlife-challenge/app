package Database

import (
    "fmt"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "golang.org/x/crypto/bcrypt"
)

var dummyHash = "$2a$10$dXJwaS5tYWluZy5jb20udG9kYXkqYEBAJCQkJCDAwMDAwMDAw"

type User struct {
    Email string `bson:"_id"`
    Password []byte `bson:"password"`
    FirstName string `bson:"first_name"`
    LastName string `bson:"last_name"`
    RecentFailedLoginAttempts string `bson:"recent_failed_login_attempts"`
}


// TODO: Change input to be one object
func (db DatabaseInstance) AddUser(email, password, first_name, last_name string) error {

    usersCollection := db.production.Collection("users")
    duplicateUser, err := db.CheckUser(email)

    if err != nil {
        return err
    }

    if duplicateUser != nil {
        return fmt.Errorf("User already exists! Please check for existing user before adding. Email: %s", duplicateUser.Email)
    }

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    user := User {
        Email: email,
        Password: hashedPassword,
        FirstName: first_name,
        LastName: last_name,
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

func (db DatabaseInstance) Authenticate(email, password string) (bool, error) {

    usersCollection := db.production.Collection("users")

    filter := bson.D{{"email", email}}

    var user User

    err := usersCollection.FindOne(db.ctx, filter).Decode(&user)

    if err == mongo.ErrNoDocuments {
        _ = bcrypt.CompareHashAndPassword([]byte(dummyHash), []byte(password))
        return false, nil
    } else if err != nil {
        return false, err
    }

    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

    if err != nil {
        return false, nil
    }

    return true, nil
}

