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
}

func NewPopulationRepository(mysql *sqlx.DB) PopulationRepository {
	return &populationRepository{
		mysql: mysql,
	}
}

func (p *populationRepository) queryLog(message string, functionName string) {
	logrus.WithFields(logrus.Fields{
		"Module":  "Repository",
		"Funtion": functionName,
	}).Info(message)
}

func (p *populationRepository) errorMessage(err error, functionName string) {
	logrus.WithFields(logrus.Fields{
		"Module":  "Repository",
		"Funtion": functionName,
	}).Error(err)
}

func (p *populationRepository) QueryAllDistrict() ([]model.District, error) {
	res := []model.District{}
	q := `SELECT Name, DistrictID
	FROM District`
	err := p.mysql.Select(&res, q)
	if err != nil {
		p.errorMessage(err, "QueryAllDistrict")
		return res, err
	}
	p.queryLog("Select all district", "QueryAllDistrict")
	return res, nil
}
func (p *populationRepository) QueryTotalPopulation(districtId int) (int64, error) {
	var res int64
	q := `
	SELECT COUNT(*) as Total
	FROM Population as p
	WHERE DistrictId=?`
	err := p.mysql.Get(&res, q, districtId)
	if err != nil {
		p.errorMessage(err, "QueryTotalPopulation")
		return res, err
	}
	p.queryLog(fmt.Sprintf("Get total population districtId: %d", districtId), "QueryTotalPopulation")
	return res, nil

}

func (p *populationRepository) QueryPeopleCommitedTheVote(districtId int) (int64, error) {
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
		p.errorMessage(err, "QueryPeopleCommitedTheVote")
		return res, err
	}
	p.queryLog(fmt.Sprintf("Get total people commited the vote districtId: %d", districtId), "QueryPeopleCommitedTheVote")
	return res, nil
}

func (p *populationRepository) QueryPeopleRightToVote(districtId int) (int64, error) {
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
		p.errorMessage(err, "QueryPeopleRightToVote")
		return res, err
	}
	p.queryLog(fmt.Sprintf("Get total people have right to vote districtId: %d", districtId), "QueryPeopleRightToVote")
	return res, nil

}
