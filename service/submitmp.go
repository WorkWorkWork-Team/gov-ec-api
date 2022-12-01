package service

import (
	"errors"

	"github.com/WorkWorkWork-Team/gov-ec-api/repository"
)

type submitmpService struct {
	submitmpRepository   repository.SubmitMpRepository
	populationRepository repository.PopulationRepository
}

type SubmitmpService interface {
	SubmitMp(citizenID string) error
}

var ErrCitizenIDNotFound = errors.New("CitizenID is not found")

func NewSubmitmpService(submitmpRepository repository.SubmitMpRepository, populationRepository repository.PopulationRepository) SubmitmpService {
	return &submitmpService{
		submitmpRepository:   submitmpRepository,
		populationRepository: populationRepository,
	}
}

func (a *submitmpService) SubmitMp(citizenID string) error {
	if !a.populationRepository.CheckIfPeopleExists(citizenID) {
		return ErrCitizenIDNotFound
	}
	return a.submitmpRepository.SubmitMpToDB(citizenID)
}
