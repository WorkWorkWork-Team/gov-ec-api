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
	out := []model.PopulationResponseItem{}
	districts, err := p.repository.QueryAllDistrict()
	if err != nil {
		logrus.Error(err)
	}
	for _, s := range districts {
		districtId := s.LocationID
		districtName := s.Location
		total, err := p.repository.QueryTotalPopulation(districtId)
		if err != nil {
			logrus.Error(err)
		}
		haveRight, err := p.repository.QueryPeopleRightToVote(districtId)
		if err != nil {
			logrus.Error(err)
		}
		commit, err := p.repository.QueryPeopleCommitedTheVote(districtId)
		if err != nil {
			logrus.Error(err)
		}
		row := model.PopulationResponseItem{
			LocationID:            districtId,
			Location:              districtName,
			TotalPeople:           total,
			PeopleWithRightToVote: haveRight,
			PeopleCommitTheVote:   commit,
		}
		out = append(out, row)
	}
	return out, nil
}
