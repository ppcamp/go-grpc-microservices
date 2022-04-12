package user_password

type Input struct {
	User     string
	Password string
}

type Output struct {
	ActivateToken string
}
