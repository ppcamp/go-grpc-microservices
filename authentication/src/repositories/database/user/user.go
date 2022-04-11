package user

func (u *UserTransaction) GetUserPassword(userId string) (string, error) {
	var password string
	var err error

	sql := `
	SELECT user_password
	FROM passwords
	WHERE user_id = ?
	ORDER BY
	LIMIT 1
	`
	err = u.Get(&password, sql)
	return password, err
}

func (u *UserTransaction) CreateUserPassword(userId string, hashedPassword string) error {
	sql := `
	INSERT INTO "password" (user_id, user_password)
	VALUES(?, ?);
	`
	_, err := u.Query(sql, userId, hashedPassword)
	return err
}
