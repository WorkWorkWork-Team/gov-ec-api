package repository

import (
	"fmt"

	"github.com/WorkWorkWork-Team/gov-ec-api/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type populationRepository struct {
	mysql *sqlx.DB
}

type PopulationRepository interface {
	QueryAllDistrict() ([]model.District, error)
	QueryTotalPopulation(districtId int) (int64, error)
	QueryPeopleCommitedTheVote(districtId int) (int64, error)
	QueryPeopleRightToVote(districtId int) (int64, error)
	QueryCandidateByDistrict(districtId int) ([]model.PopulationDatabaseRow, error)
	QueryAllCandidate() ([]model.PopulationDatabaseRow, error)
}

func NewPopulationRepository(mysql *sqlx.DB) PopulationRepository {
	return &populationRepository{
		mysql: mysql,
	}
}

func (p *populationRepository) generateLogger(functionName string) *logrus.Entry {
	return logrus.WithFields(logrus.Fields{
		"Module":  "Repository",
		"Funtion": functionName,
	})
}

func (p *populationRepository) QueryAllDistrict() ([]model.District, error) {
	logger := p.generateLogger("QueryAllDistrict")
	res := []model.District{}
	q := `SELECT Name, DistrictID
	FROM District`
	err := p.mysql.Select(&res, q)
	if err != nil {
		logger.Error(err)
		return res, err
	}
	logger.Info("Select all district")
	return res, nil
}

func (p *populationRepository) QueryTotalPopulation(districtId int) (int64, error) {
	logger := p.generateLogger("QueryTotalPopulation")
	var res int64
	q := `
	SELECT COUNT(*) as Total
	FROM Population as p
	WHERE DistrictId=?`
	err := p.mysql.Get(&res, q, districtId)
	if err != nil {
		logger.Error(err)
		return res, err
	}
	logger.Info("Get total population districtId: ", districtId)
	return res, nil

}

func (p *populationRepository) QueryPeopleCommitedTheVote(districtId int) (int64, error) {
	logger := p.generateLogger("QueryPeopleCommitedTheVote")
	var res int64
	q := `
	SELECT COUNT(*) as Commits
	FROM 
	(
		SELECT p.CitizenID, p.DistrictID
		FROM Population as p 
		JOIN ApplyVote as a 
		on a.CitizenID = p.CitizenID
	) as b
	WHERE b.DistrictId = ?
	`
	err := p.mysql.Get(&res, q, districtId)
	if err != nil {
		logger.Error(err)
		return res, err
	}
	logger.Info("Get total people commited the vote districtId: ", districtId)
	return res, nil
}

func (p *populationRepository) QueryPeopleRightToVote(districtId int) (int64, error) {
	logger := p.generateLogger("QueryPeopleRightToVote")
	var res int64
	q := `
	SELECT COUNT(*) as HaveRight
	FROM (
			SELECT * FROM Population AS p
			WHERE p.Birthday < DATE_ADD(CURRENT_DATE(), INTERVAL -18 YEAR) 
			AND p.CitizenID NOT IN (SELECT CitizenID FROM ApplyVote)
			AND p.DistrictId=?
			AND p.Nationality="Thai"
		) AS c
	`
	err := p.mysql.Get(&res, q, districtId)
	if err != nil {
		logger.Error(err)
		return res, err
	}
	logger.Info("Get total people have right to vote districtId: ", districtId)
	return res, nil
}

func (p *populationRepository) QueryCandidateByDistrict(districtId int) ([]model.PopulationDatabaseRow, error) {
	logger := p.generateLogger("QueryCandidateByDistrict")
	res := []model.PopulationDatabaseRow{}
	q := `
	SELECT p.CitizenID, LazerID, Name, Lastname, Birthday, Nationality, DistrictID
	FROM Population AS p
	JOIN Candidate AS c
	ON p.CitizenID = c.CitizenID
	WHERE p.DistrictID=?
	`
	err := p.mysql.Select(&res, q, districtId)
	if err != nil {
		logger.Error(err)
		return res, err
	}
	logger.Info(fmt.Sprintf("Get all candidates from districtId: %d", districtId))
	return res, nil
}

func (p *populationRepository) QueryAllCandidate() ([]model.PopulationDatabaseRow, error) {
	logger := p.generateLogger("QueryAllCandidate")
	res := []model.PopulationDatabaseRow{}
	q := `
	SELECT p.CitizenID, LazerID, Name, Lastname, Birthday, Nationality, DistrictID
	FROM Population AS p
	JOIN Candidate AS c
	ON p.CitizenID = c.CitizenID
	`
	err := p.mysql.Select(&res, q)
	if err != nil {
		logger.Error(err)
		return res, err
	}
	logger.Info("Get all candidates")
	return res, nil
}
