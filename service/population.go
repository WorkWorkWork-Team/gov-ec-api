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
	districts, err := p.repository.QueryAllDistrict()
	if err != nil {
		p.errorMessage(err, "QueryAllDistrict")
		return []model.PopulationResponseItem{}, err
	}
	for _, s := range districts {
		districtId := s.DistrictID
		districtName := s.Name
		total, err := p.repository.QueryTotalPopulation(districtId)
		if err != nil {
			p.errorMessage(err, "QueryTotalPopulation")
			return []model.PopulationResponseItem{}, err
		}
		haveRight, err := p.repository.QueryPeopleRightToVote(districtId)
		if err != nil {
			p.errorMessage(err, "QueryPeopleRightToVote")
			return []model.PopulationResponseItem{}, err
		}
		commit, err := p.repository.QueryPeopleCommitedTheVote(districtId)
		if err != nil {
			p.errorMessage(err, "QueryPeopleCommitedTheVote")
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
	p.queryLog("return population statistic", "GetPopulationStatistics")
	return populationStatistic, nil
}

func (p *populationService) GetAllCandidateInfo() (candidates []model.PopulationDatabaseRow, err error) {
	candidates, err = p.GetAllCandidateInfo()
	if err != nil {
		p.errorMessage(err, "GetAllCandidateInfo")
		return
	}
	p.queryLog("return all candidate informatiom", "GetPopulationStatistics")
	return candidates, nil
}
