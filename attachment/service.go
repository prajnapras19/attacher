package attachment

import (
	"os"

	"github.com/prajnapras19/attacher/config"
	"github.com/prajnapras19/attacher/lib"
)

type Service interface {
	GetAllActiveAttachmentsByUserID(userID uint) ([]*Attachment, error)
	GetActiveAttachmentByUserIDAndSerial(userID uint, serial string) (*Attachment, error)
}

type service struct {
	cfg                  *config.Config
	attachmentRepository Repository
}

func NewService(
	cfg *config.Config,
	attachmentRepository Repository,
) Service {
	return &service{
		cfg:                  cfg,
		attachmentRepository: attachmentRepository,
	}
}

func (s *service) GetAllActiveAttachmentsByUserID(userID uint) ([]*Attachment, error) {
	res, err := s.attachmentRepository.GetAllActiveAttachmentsByUserID(userID)
	if err != nil {
		return nil, lib.ErrFailedToGetAttachments
	}
	return res, nil
}

func (s *service) GetActiveAttachmentByUserIDAndSerial(userID uint, serial string) (*Attachment, error) {
	res, err := s.attachmentRepository.GetActiveAttachmentByUserIDAndSerial(userID, serial)
	if err != nil {
		return nil, lib.ErrAttachmentNotFound
	}
	if _, err := os.Stat(res.Path); err != nil {
		return nil, lib.ErrAttachmentNotFound
	}
	return res, nil
}
