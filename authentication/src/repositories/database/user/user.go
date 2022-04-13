package user

func (u *UserTransaction) GetUserPassword(userId string) (string, error) {
	var password string
	var err error

	sql := `
	SELECT user_password
	FROM passwords
	WHERE user_id = $1
		AND active = TRUE
	ORDER BY created_at DESC
	LIMIT 1
	`
	err = u.Get(&password, sql, userId)
	return password, err
}

func (u *UserTransaction) CreateUserPassword(userId string, hashedPassword string) error {
	sql := `
	INSERT INTO passwords(user_id, user_password)
	VALUES($1, $2);
	`
	_, err := u.Query(sql, userId, hashedPassword)
	return err
}
