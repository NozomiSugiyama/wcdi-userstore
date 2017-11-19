package model

type User struct {
	ID                string   `json:"_id"`
	DisplayName       string   `json:"display_name"`
	FirstName         string   `json:"first_name"`
	LastName          string   `json:"last_name"`
	FirstNameEnglish  string   `json:"first_name_english"`
	LastNameEnglish   string   `json:"last_name_english"`
	Birthday          string   `json:"birthday"`
	Gender            string   `json:"gender"`
	Email             string   `json:"email"`
	EmailForShared    string   `json:"email_for_shared"`
	CountryPrefecture string   `json:"country_prefecture"`
	PhoneNumber       string   `json:"phone_number"`
	SlackID           string   `json:"slack_id"`
	GithubID          string   `json:"github_id"`
	License           []string `json:"license"`
	Neec              struct {
		StudentID      string `json:"student_id"`
		CollegeName    string `json:"college_name"`
		DepartmentName string `json:"department_name"`
		AdmissionYear  int    `json:"admission_year"`
		Class          int    `json:"class"`
	} `json:"neec"`
	Wcdi struct {
		Position string `json:"position"`
	} `json:"wcdi"`
}

type Users []User

func GetUsers() (Users, error) {
	users := Users{}
	return users, nil
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
