package models

import (
	"furrble.com/backend/forms"
	"gopkg.in/mgo.v2/bson"
)

//Address Structure
type AddressType struct {
	house string
	line1 string
	line2 string
	city string
	state string
	country string
	pincode int
	lat float64
	lon float64
}

//User structure
type Pet struct {
	ID          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string        `json:"name" bson:"name"`
	Animal      string		 `json:"animal" bson:"animal"`
	Breed       string        `json:"email" bson:"email"`
	SecondBreed string
	DOB         string
	weight      int
	male        bool
	neutered    bool
	trackerid   string
}

type Pets struct {
	pets []Pets
}

//User structure
type User struct {
	ID         bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string        `json:"name" bson:"name"`
	Email      string        `json:"email" bson:"email"`
	Phone      string        `json:"phone" bson:"phone"`
	Address    AddressType
	PetList    Pets
}

func (u *User) AddNameEmail(data forms.SignupUserCommand) error {
	// Connect to the user collection
	collection := dbConnect.Use("fudb1", "fudb1c1")
	// Assign result to error object while saving user
	err := collection.Insert(bson.M{
		"name":        data.Name,
		"email":       data.Email,
	})
	return err
}
// GetUserByEmail handles fetching user by email
func (u *User) GetUserByEmail(email string) (user User, err error) {
	// Connect to the user collection
	collection := dbConnect.Use("fudb1", "fudb1c1")
	// Assign result to error object while saving user
	err = collection.Find(bson.M{"email": email}).One(&user)
	return user, err
}
