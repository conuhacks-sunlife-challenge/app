package Database

type User struct {
    Email string `bson:"email"`
    Password string `bson:"password"`
    FirstName string `bson:"first_name"`
    LastName string `bson:"last_name"`
}


// TODO: Change input to be one object
func (db DatabaseInstance) AddUser(email, password, firstname, lastname string) error {

    usersCollection := db.production.Collection("users")

    user := User{
        Email: email,
        Password: password,
        FirstName: firstname,
        LastName: lastname,
    }

    usersCollection.InsertOne(db.ctx, user)


    return nil
}
