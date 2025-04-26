package object

import "time"

type User struct {
	UserID       int       `json:"user_id"`
	UserName     string    `json:"username"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	Subscription string    `json:"subscription"`
	Registration time.Time `json:"registration"`
	Age          int       `json:"age"`
	ProfilePic   string    `json:"profile_pic"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
