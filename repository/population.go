package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type populationRepository struct {
	mysql *sqlx.DB
}

type PopulationRepository interface {
	GetTotalNumPopulation() (ResponseTotalPop, error)
}

func NewPopulationRepository(mysql *sqlx.DB) PopulationRepository {
	return &populationRepository{
		mysql: mysql,
	}
}

type ResponseTotalPop struct {
	logrus.WithFields(logrus.Fields{
		"Database": "MYSQL",
		"Hostname": config.Hostname,
	}).Error(err)
}

func (p *populationRepository) GetTotalNumPopulation() (ResponseTotalPop, error) {
	q := "SELECT COUNT(DISTINCT CitizenID) as Total FROM Population"
	popRow := ResponseTotalPop{}
	err := p.mysql.Get(&popRow, q)

	
	if err != nil {
		logrus.Fields()
		logrus.Error(popRow)
		return popRow, err
	}

	return popRow, nil
}
