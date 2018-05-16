package model

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID                   int      `bson:"_id"`
	DisplayName          string   `bson:"display_name"`
	FirstName            string   `bson:"first_name"`
	LastName             string   `bson:"last_name"`
	FirstNameEnglish     string   `bson:"first_name_english"`
	LastNameEnglish      string   `bson:"last_name_english"`
	Birthday             string   `bson:"birthday"`
	Gender               string   `bson:"gender"`
	Email                string   `bson:"email"`
	EmailForShared       string   `bson:"email_for_shared"`
	CountryPrefecture    string   `bson:"country_prefecture"`
	PhoneNumber          string   `bson:"phone_number"`
	SlackID              string   `bson:"slack_id"`
	GithubID             string   `bson:"github_id"`
	License              []string `bson:"license"`
	SubscriptionReligion struct {
		ID string `bson:"_id"`
	} `bson:"subscription_religion"`
}

type Users []User

func GetUsers() (Users, error) {
	var session = GetSession("mongodb")
	db := session.DB("test")

	var users []User
	db.C("users").Find(bson.M{}).All(&users)
	fmt.Println("Results All: ", users)

	return users, nil
}

func CreateUser(user User) error {
	return nil
}

func GetUser(id int) (User, error) {
	var session = GetSession("mongodb")
	db := session.DB("test")

	var users []User
	db.C("users").Find(bson.M{}).All(&users)
	fmt.Println("Results All: ", users)
	
	user := User{}
	return user, nil
}

func UpdateUser(user User) error {
	return nil
}

func DeleteUser(id string) error {
	return nil
}
