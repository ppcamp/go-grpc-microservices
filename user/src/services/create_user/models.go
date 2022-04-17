package create_user

import "time"

type Input struct {
	FirstName  string
	MiddleName string
	LastName   string
	Nick       string
	BirthDate  time.Time
	Email      string
	Password   string
}

type Output struct{}
