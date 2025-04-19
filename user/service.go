package user

import (
	"github.com/prajnapras19/attacher/config"
)

type Service interface {
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
