package data

import (
	"encoding/json"
	"io"
	"time"
)

type User struct {
	Id        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"-"`
	CreatedAt string `json:"-"`
	UpdatedAt string `json:"-"`
}
type Users []*User

func GetUsers() *Users {
	return &users
}

func (users *Users) ToJson(w io.Writer) error {

	encoder := json.NewEncoder(w)
	return encoder.Encode(users)

}

var users = Users{
	{
		Id:        0,
		Username:  "Neo",
		Email:     "Neo@matrix.com",
		Password:  hash("theonethatbreakthematrix"),
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
	{
		Id:        1,
		Username:  "Morpheus",
		Email:     "morpheus@matrix.com",
		Password:  hash("theonewhofindtheonethatbreakthematrix"),
		CreatedAt: time.Now().UTC().String(),
		UpdatedAt: time.Now().UTC().String(),
	},
}

func hash(input string) string {
	//TODO impliment hash
	return input
}
