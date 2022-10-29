package service

import (
	"github.com/WorkWorkWork-Team/gov-ec-api/model"
	"github.com/WorkWorkWork-Team/gov-ec-api/repository"
	"github.com/sirupsen/logrus"
)

type populationService struct {
	repository repository.PopulationRepository
}

type PopulationService interface {
	GetPopulationStatistics() ([]model.PopulationResponseItem, error)
}

func NewPopulationService(populationRepository repository.PopulationRepository) PopulationService {
	return &populationService{
		repository: populationRepository,
	}
}

func (p *populationService) GetPopulationStatistics() ([]model.PopulationResponseItem, error) {
	data, err := p.repository.QueryPopulationStat()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Module":  "Service",
			"Funtion": "GetPopulationStatistics",
		})
		logrus.Error(err)
		return data, err
	}
	return data, nil
}
