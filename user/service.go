package user

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/prajnapras19/attacher/config"
	"github.com/prajnapras19/attacher/lib"
)

type Service interface {
	Login(req *LoginRequest) (*LoginResponse, error)
	ValidateToken(tokenString string) (*lib.JWTClaims, error)
}

type service struct {
	cfg            *config.Config
	userRepository Repository
}

func NewService(
	cfg *config.Config,
	userRepository Repository,
) Service {
	return &service{
		cfg:            cfg,
		userRepository: userRepository,
	}
}

func (s *service) Login(req *LoginRequest) (*LoginResponse, error) {
	user, err := s.userRepository.GetUserByUsername(req.Username)
	if err != nil {
		return nil, lib.ErrUserNotFound
	}
	if req.Password != user.Password {
		return nil, lib.ErrUserNotFound
	}
	return &LoginResponse{
		Token: s.GenerateToken(user.Username, user.Serial),
	}, nil
}

func (s *service) GenerateToken(username string, serial string) string {
	claims := lib.JWTClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    s.cfg.AuthConfig.ApplicationName,
			ExpiresAt: time.Now().Add(s.cfg.AuthConfig.LoginTokenExpirationDuration).Unix(),
		},
		Username: username,
		Serial:   serial,
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	signedToken, _ := token.SignedString(s.cfg.AuthConfig.SignatureKey)
	return signedToken
}

func (s *service) VerifyToken(token *jwt.Token) (interface{}, error) {
	if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, lib.ErrSigningMethodInvalid
	} else if method != jwt.SigningMethodHS256 {
		return nil, lib.ErrSigningMethodInvalid
	}
	return s.cfg.AuthConfig.SignatureKey, nil
}

func (s *service) ValidateToken(tokenString string) (*lib.JWTClaims, error) {
	claims := &lib.JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, s.VerifyToken)
	if err != nil {
		return nil, lib.ErrUnauthorizedRequest
	}
	if !token.Valid {
		return nil, lib.ErrUnauthorizedRequest
	}
	claims, ok := token.Claims.(*lib.JWTClaims)
	if !ok {
		return nil, lib.ErrUnauthorizedRequest
	}

	// validate claims based on db
	claimedUser, err := s.userRepository.GetUserByUsername(claims.Username)
	if err != nil {
		return nil, lib.ErrUnauthorizedRequest
	}
	if claimedUser.Serial != claims.Serial {
		return nil, lib.ErrUnauthorizedRequest
	}

	return claims, nil
}
