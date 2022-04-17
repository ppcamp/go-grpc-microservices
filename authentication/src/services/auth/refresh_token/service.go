package refresh_token

import (
	"time"

	"authentication/helpers/services"
	"authentication/repositories/cache"
	"authentication/utils/jwt"
)

type RefreshTokenService[In, Out any] struct {
	services.BaseBusiness

	cache     cache.Auth
	signerExp time.Duration
}

// NewService creates a service that get user password, check it, and
// return a valid JWT token
func NewService(
	repo cache.Auth,
	exp time.Duration,
) services.IBaseBusiness[Input, Output] {
	return &RefreshTokenService[Input, Output]{cache: repo, signerExp: exp}
}

func (s *RefreshTokenService[In, Out]) Execute(in Input) (*Output, error) {
	exp := time.Now().Add(s.signerExp)

	session, err := jwt.Signer.Session(in.Token)
	if err != nil {
		return nil, err
	}

	err = s.cache.Invalidate(s.Context, *session.UserId, in.Token)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Signer.Generate(session, s.signerExp)
	if err != nil {
		return nil, err
	}

	err = s.cache.Authorize(s.Context, *session.UserId, token, exp)
	return new(Output), err
}
