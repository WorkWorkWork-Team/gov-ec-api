package service

import (
	"database/sql"

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

func (p *populationService) errorMessage(err error, functionName string) {
	logrus.WithFields(logrus.Fields{
		"Module":  "Service",
		"Funtion": functionName,
	}).Error(err)
}

func (p *populationService) queryLog(message string, functionName string) {
	logrus.WithFields(logrus.Fields{
		"Module":  "Service",
		"Funtion": functionName,
	}).Info(message)
}

func (p *populationService) GetPopulationStatistics() ([]model.PopulationResponseItem, error) {
	out := []model.PopulationResponseItem{}
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
		out = append(out, row)
	}
	p.queryLog("return population statistic", "GetPopulationStatistics")
	return out, nil
}

func (p *populationService) GetAllCandidateInfo() ([]model.PopulationDatabaseRow, error) {
	out := []model.PopulationDatabaseRow{}
	districts, err := p.repository.QueryAllDistrict()
	if err != nil {
		p.errorMessage(err, "QueryAllDistrict")
		return []model.PopulationDatabaseRow{}, err
	}
	for _, s := range districts {
		districtId := s.DistrictID
		res, err := p.repository.QueryCandidateByDistrict(districtId)
		if err != nil {
			if err == sql.ErrNoRows {
				p.errorMessage(err, "QueryCandidateByDistrict")
			} else {
				p.errorMessage(err, "QueryCandidateByDistrict")
				return []model.PopulationDatabaseRow{}, err
			}
		}
		for _, c := range res {
			out = append(out, c)
		}
	}
	p.queryLog("return all candidate informatiom", "GetPopulationStatistics")
	return out, nil
}
