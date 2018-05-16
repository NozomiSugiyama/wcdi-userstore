package model

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID                   string   `bson:"_id"`
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
	// users := Users{}
	var session = GetSession("mongodb")
	db := session.DB("test")

	// p := make([]User, 0, 10)
	// query := db.C("users").Find(bson.M{})
	// query.All(&p)

	var results []User
	db.C("users").Find(bson.M{}).All(&results)
	fmt.Println("Results All: ", results)

	fmt.Printf("%+v\n", results)

	// users = append(users, User{ID: "1"})
	// users = append(users, User{ID: "2"})
	// users = append(users, User{ID: "3"})

	return results, nil
}

func CreateUser(user User) error {
	return nil
}

func GetUser(id string) (User, error) {
	user := User{}
	return user, nil
}

func UpdateUser(user User) error {
	return nil
}

func DeleteUser(id string) error {
	return nil
}
