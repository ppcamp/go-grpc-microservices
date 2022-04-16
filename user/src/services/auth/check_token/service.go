package check_token

import (
	"streamer/helpers/services"
	"streamer/repositories/cache"
)

type CheckTokenService[In, Out any] struct {
	services.BaseBusiness

	cache cache.Auth
}

// NewService creates a service that get user password, check it, and
// return a valid JWT token
func NewService(repo cache.Auth) services.IBaseBusiness[Input, Output] {
	return &CheckTokenService[Input, Output]{cache: repo}
}

func (s *CheckTokenService[In, Out]) Execute(in Input) (*Output, error) {
	err := s.cache.Valid(s.Context, in.Token, in.Token)
	if err != nil {
		return nil, err
	}
	return &Output{}, nil
}
