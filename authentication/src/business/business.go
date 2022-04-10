package business

import (
	"streamer/business/base"
	"streamer/business/user_login"
	"streamer/repositories/cache"
	"streamer/utils/jwt"
)

type handler struct {
	Cache  cache.Cache
	Signer jwt.Jwt
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Handle(ctrl base.BaseBusiness[any, any]) {
	switch ctrl := ctrl.(type) {
	case base.BaseTransactionBusiness[any, any]:
	default:
		result, err := h.Handle()
	}
}

func (h *handler) HandleUserLogin() base.BaseBusiness[user_login.UserLoginInput, user_login.UserLoginOutput] {
	return user_login.NewUserLoginHandler(h.Cache, h.Signer)
}

func a() {
	b := &handler{}

	b.HandleUserLogin().Execute()
}
