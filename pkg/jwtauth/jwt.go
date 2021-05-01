package jwtauth

import (
	"strings"
	"time"

	kinet "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

type JWTAuth struct {
	p *profile
}

func New(opts ...Option) *JWTAuth {
	p := defaultProfile
	for _, opt := range opts {
		opt(&p)
	}
	return &JWTAuth{
		p: &p,
	}
}

// GenerateToken return a token string, an expires (millisecond unit) and an error if has.
func (a *JWTAuth) GenerateToken(subject string) (string, int64, error) {
	now := time.Now()
	duration := time.Duration(a.p.expired) * time.Second
	expiresAt := now.Add(duration).UnixNano() / 1e6
	token := kinet.NewWithClaims(a.p.signingMethod, &kinet.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: expiresAt,
		NotBefore: now.Unix(),
		Subject:   subject,
	})
	jwtString, err := token.SignedString(a.p.signingKey)
	if err != nil {
		return "", 0, errors.Wrap(err, "jwt auth - failed to generate jwt token string")
	}
	return a.p.tokenPrefix + jwtString, expiresAt, nil
}

func (a *JWTAuth) parseToken(jwtString string) (*kinet.StandardClaims, error) {
	token, err := kinet.ParseWithClaims(jwtString, &kinet.StandardClaims{}, a.p.keyFunc)
	if err != nil {
		return nil, err
	} else if !token.Valid {
		return nil, ErrInvalidToken
	}
	return token.Claims.(*kinet.StandardClaims), nil
}

func (a *JWTAuth) ParseSubject(jwtString string) (string, error) {
	if jwtString == "" {
		return "", ErrInvalidToken
	}
	var tokenString string
	if strings.Contains(jwtString, a.p.tokenPrefix) {
		tokenString = strings.TrimPrefix(jwtString, a.p.tokenPrefix)
	}
	claims, err := a.parseToken(tokenString)
	if err != nil {
		return "", errors.Wrap(err, "jwt auth - failed to parse jwt token string")
	}
	return claims.Subject, nil
}

// Common errors ---------------------------------------------------------------

var (
	ErrInvalidToken = errors.New("invalid token")
)

// Options ---------------------------------------------------------------------

const (
	defaultSigningKey = "jwt-sprout:v0.0.1"

	tokenPrefixBearer = "Bearer "

	appointedSeatHeader = "header:Authorization"
)

var (
	defaultProfile = profile{
		signingKey:    []byte(defaultSigningKey),
		signingMethod: kinet.SigningMethodHS512,
		keyFunc: func(t *kinet.Token) (interface{}, error) {
			if _, ok := t.Method.(*kinet.SigningMethodHMAC); !ok {
				return nil, ErrInvalidToken
			}
			return []byte(defaultSigningKey), nil
		},
		expired:       7200,
		tokenPrefix:   tokenPrefixBearer,
		appointedSeat: appointedSeatHeader,
	}
)

type profile struct {
	signingKey    []byte
	signingMethod kinet.SigningMethod
	keyFunc       kinet.Keyfunc
	expired       int64
	tokenPrefix   string
	appointedSeat string
}

type Option func(*profile)

func SetSigningKey(key []byte) Option {
	return func(p *profile) {
		p.signingKey = key
	}
}

func SetSigningMethod(method kinet.SigningMethod) Option {
	return func(p *profile) {
		p.signingMethod = method
	}
}

func SetKeyFunc(keyFunc kinet.Keyfunc) Option {
	return func(p *profile) {
		p.keyFunc = keyFunc
	}
}

func SetExpired(expired int64) Option {
	return func(p *profile) {
		p.expired = expired
	}
}

func SetTokenPrefix(prefix string) Option {
	return func(p *profile) {
		p.tokenPrefix = prefix
	}
}

// SetAppointedSeat the seat parameter is a string in the form of "<source>:<name>" that is used to extract token from the request.
// Possible values:
// - "header:<name>"
// - "query:<name>"
// - "cookie:<name>"
func SetAppointedSeat(seat string) Option {
	return func(p *profile) {
		p.appointedSeat = seat
	}
}
