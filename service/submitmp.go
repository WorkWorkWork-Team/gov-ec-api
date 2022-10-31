package service

import (
	"github.com/WorkWorkWork-Team/gov-ec-api/repository"
)

type submitmpService struct {
	submitmpRepository repository.SubmitMpRepository
}

type SubmitmpService interface {
	SubmitMp(citizenID string) error
}

func NewSubmitmpService(submitmpRepository repository.SubmitMpRepository) SubmitmpService {
	return &submitmpService{
		submitmpRepository: submitmpRepository,
	}
}

func (a *submitmpService) SubmitMp(citizenID string) error {
	return a.submitmpRepository.SubmitMpToDB(citizenID)
}
