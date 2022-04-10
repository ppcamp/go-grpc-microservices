package cache

import "time"

type Auth interface {
	StoreToken(token string, exp time.Duration) error
	CheckToken(token string) error
}

type auth struct{}

func NewAuthRepository() Auth {
	return &auth{}
}

func (s *auth) StoreToken(token string, exp time.Duration) error {
	return nil
}

func (s *auth) CheckToken(token string) error {
	return nil
}
