package repository

import (
	"github.com/WorkWorkWork-Team/gov-ec-api/model"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type populationRepository struct {
	mysql *sqlx.DB
}

type PopulationRepository interface {
	QueryPopulationStat() ([]model.PopulationResponseItem, error)
}

func NewPopulationRepository(mysql *sqlx.DB) PopulationRepository {
	return &populationRepository{
		mysql: mysql,
	}
}

func (p *populationRepository) QueryPopulationStat() ([]model.PopulationResponseItem, error) {
	res := []model.PopulationResponseItem{}
	q := `
	SELECT r.DistrictID, d.Name, COALESCE(r.HaveRight, 0) as HaveRight, COALESCE(Commits.Commits, 0 ) as Commits, COALESCE(t.Total, 0) as Total
	FROM
	(    
		SELECT d.DistrictID, COUNT(*) as HaveRight
		FROM (
				SELECT * FROM Population AS p
				WHERE p.Birthday < DATE_ADD(CURRENT_DATE(), INTERVAL -18 YEAR) 
				AND p.CitizenID NOT IN (SELECT CitizenID FROM ApplyVote)
			) AS c
		JOIN
		District AS d
		on d.DistrictID = c.DistrictID
		GROUP By c.DistrictID
	) AS r
	LEFT JOIN
	(
		SELECT COUNT(*) as Commits, b.DistrictId
		FROM 
		(
			SELECT p.CitizenID, p.DistrictID
			FROM Population as p 
			JOIN ApplyVote as a 
			on a.CitizenID = p.CitizenID
		) as b
		GROUP BY b.DistrictID
	) as Commits
	ON Commits.DistrictId = r.DistrictId
	LEFT JOIN
	(
		SELECT COUNT(*) as Total, p.DistrictID
		FROM Population as p
		GROUP BY p.DistrictID
	) as t
	ON t.DistrictId = r.DistrictId
	LEFT JOIN
	(
		SELECT Name, DistrictID
		FROM District
	) as d
	ON d.DistrictID = r.DistrictId
`
	err := p.mysql.Select(&res, q)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"Module":  "Repository",
			"Funtion": "GetTotalNumPopulation",
		})
		logrus.Error(err)
		return res, err
	}
	logrus.Info(res)
	return res, nil
}
