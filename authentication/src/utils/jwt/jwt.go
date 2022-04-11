package jwt

import (
	"crypto/ecdsa"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

const UnexpectedMethod string = "unexpected signing method: %v"

type Session struct {
	UserId *string `json:"userId,omitempty"`
}

type Claims struct {
	*Session
	jwt.StandardClaims
}

type Jwt interface {
	Generate(session *Session, expSecs time.Duration) (string, error)
	Parse(signedToken string) (*Claims, error)
}

type impl struct {
	secret *ecdsa.PrivateKey
	issuer string
}

// NewJwt creates a new jwt encoder/decoder. To this implementation, I'm using
// an ES512 algorithm
//
// Note
//
// - The HMAC signing method (HS256,HS384,HS512) expect []byte values for
// signing and validation
//
// - The RSA signing method (RS256,RS384,RS512) expect *rsa.PrivateKey for
// signing and *rsa.PublicKey for validation
//
// - The ECDSA signing method (ES256,ES384,ES512) expect *ecdsa.PrivateKey for
// signing and *ecdsa.PublicKey for validation
//
// - The EDSA signing method (Ed25519) expect ed25519.PrivateKey for signing
// and ed25519.PublicKey for validation
func NewJwt(secret *ecdsa.PrivateKey) Jwt {
	return &impl{
		secret: secret,
		issuer: "authentication",
	}
}

// Generate some token for a given session
func (j *impl) Generate(session *Session, expSecs time.Duration) (string, error) {
	claims := &Claims{
		Session: session,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: 15000,
			Issuer:    j.issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES512, claims)
	return token.SignedString(j.secret)
}

// Parse does the decrypt of some token. It also checks for the correct alg
// using Go's reflection
func (j *impl) Parse(signedToken string) (*Claims, error) {
	keyFunc := func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodECDSA); !ok {
			return nil, fmt.Errorf(UnexpectedMethod, token.Header["alg"])
		}
		return j.secret, nil
	}

	var claims *Claims

	_, err := jwt.ParseWithClaims(signedToken, claims, keyFunc)
	if err != nil {
		return nil, err
	}
	return claims, nil
}
