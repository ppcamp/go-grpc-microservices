package invalidate_tokens

import (
	"authentication/helpers/services"
	"authentication/repositories/cache"
	"authentication/utils/jwt"
)

type InvalidateTokensService[In, Out any] struct {
	services.BaseBusiness

	cache cache.Auth
}

// NewService creates a service that get user password, check it, and
// return a valid JWT token
func NewService(repo cache.Auth) services.IBaseBusiness[Input, Output] {
	return &InvalidateTokensService[Input, Output]{cache: repo}
}

func (s *InvalidateTokensService[In, Out]) Execute(in Input) (*Output, error) {
	if _, err := jwt.Signer.Session(in.Token); err != nil {
		return nil, err
	}
	err := s.cache.InvalidateAll(s.Context, in.User)
	return &Output{}, err
}
