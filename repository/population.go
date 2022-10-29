package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type populationRepository struct {
	mysql *sqlx.DB
}

type PopulationRepository interface {
	QueryAllDistrict() ([]ResultQueryDistrictIdAndName, error)
	QueryTotalPopulation(districtId string) (ResultQueryTotal, error)
	QueryPeopleCommitedTheVote(districtId string) (ResultQueryPeopleCommitTheVote, error)
	QueryPeopleRightToVote(districtId string) (ResultQueryPeopleWithRightToVote, error)
}

func NewPopulationRepository(mysql *sqlx.DB) PopulationRepository {
	return &populationRepository{
		mysql: mysql,
	}
}

type ResultQueryDistrictIdAndName struct {
	LocationID string `db:"DistrictID"`
	Location   string `db:"Name"`
}

type ResultQueryTotal struct {
	Total int64 `db:"Total"`
}
type ResultQueryPeopleWithRightToVote struct {
	PeopleWithRightToVote int64 `db:"HaveRight"`
}
type ResultQueryPeopleCommitTheVote struct {
	PeopleCommitTheVote int64 `db:"Commits"`
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

func (p *populationRepository) QueryAllDistrict() ([]ResultQueryDistrictIdAndName, error) {
	res := []ResultQueryDistrictIdAndName{}
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
func (p *populationRepository) QueryTotalPopulation(districtId string) (ResultQueryTotal, error) {
	res := ResultQueryTotal{}
	q := `
	SELECT COUNT(*) as Total
	FROM Population as p
	WHERE DistrictId=?`
	err := p.mysql.Get(&res, q, districtId)
	if err != nil {
		p.errorMessage(err, "QueryTotalPopulation")
		return res, err
	}
	p.queryLog(fmt.Sprintf("Get total population districtId: %s", districtId), "QueryTotalPopulation")
	return res, nil

}

func (p *populationRepository) QueryPeopleCommitedTheVote(districtId string) (ResultQueryPeopleCommitTheVote, error) {
	res := ResultQueryPeopleCommitTheVote{}
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
	p.queryLog(fmt.Sprintf("Get total people commited the vote districtId: %s", districtId), "QueryPeopleCommitedTheVote")
	return res, nil
}

func (p *populationRepository) QueryPeopleRightToVote(districtId string) (ResultQueryPeopleWithRightToVote, error) {
	res := ResultQueryPeopleWithRightToVote{}
	q := `
	SELECT COUNT(*) as HaveRight
	FROM (
			SELECT * FROM Population AS p
			WHERE p.Birthday < DATE_ADD(CURRENT_DATE(), INTERVAL -18 YEAR) 
			AND p.CitizenID NOT IN (SELECT CitizenID FROM ApplyVote)
			AND p.DistrictId=?
		) AS c
	`
	err := p.mysql.Get(&res, q, districtId)
	if err != nil {
		p.errorMessage(err, "QueryPeopleRightToVote")
		return res, err
	}
	p.queryLog(fmt.Sprintf("Get total people have right to vote districtId: %s", districtId), "QueryPeopleRightToVote")
	return res, nil

}
