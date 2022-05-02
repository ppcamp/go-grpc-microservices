package refresh_token

import "time"

type Input struct {
	Token string
}

type Output struct {
	Token string
	Exp   time.Time
}
