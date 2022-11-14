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
	GetAllCandidateInfo() ([]model.PopulationDatabaseRow, error)
}

func NewPopulationService(populationRepository repository.PopulationRepository) PopulationService {
	return &populationService{
		repository: populationRepository,
	}
}

func (p *populationService) generateLogger(functionName string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"Module":  "Service",
		"Funtion": functionName,
	})
}

func (p *populationService) GetPopulationStatistics() (populationStatistic []model.PopulationResponseItem, err error) {
	logger := p.generateLogger("QueryAllDistrict")
	districts, err := p.repository.QueryAllDistrict()
	if err != nil {
		logger.Error(err)
		return []model.PopulationResponseItem{}, err
	}
	for _, s := range districts {
		districtId := s.DistrictID
		districtName := s.Name
		total, err := p.repository.QueryTotalPopulation(districtId)
		if err != nil {
			logger.Error(err)
			return []model.PopulationResponseItem{}, err
		}
		haveRight, err := p.repository.QueryPeopleRightToVote(districtId)
		if err != nil {
			logger.Error(err)
			return []model.PopulationResponseItem{}, err
		}
		commit, err := p.repository.QueryPeopleCommitedTheVote(districtId)
		if err != nil {
			logger.Error(err)
			return []model.PopulationResponseItem{}, err
		}
		row := model.PopulationResponseItem{
			LocationID:            districtId,
			Location:              districtName,
			TotalPeople:           total,
			PeopleWithRightToVote: haveRight,
			PeopleCommitTheVote:   commit,
		}
		populationStatistic = append(populationStatistic, row)
	}
	logger.Info("return population statistics")
	return populationStatistic, nil
}

func (p *populationService) GetAllCandidateInfo() (candidates []model.PopulationDatabaseRow, err error) {
	logger := p.generateLogger("GetAllCandidateInfo")
	candidates, err = p.repository.QueryAllCandidate()
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info("return all candidate info")
	return candidates, nil
}
