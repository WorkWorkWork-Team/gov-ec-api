package repository

import (
	"database/sql"
	"fmt"
)

type populationRepository struct {
	mysql *sql.DB
}

type PopulationRepository interface {
}

func NewPopulationRepository(mysql *sql.DB) PopulationRepository {
	return &populationRepository{
		mysql: mysql,
	}
}

func GetTotalNumPopulation(mysql *sql.DB) {
	q := `
	SELECT COUNT(CitizenID)
	FROM Population
	GROUP BY CitizenID`
	fmt.Println(mysql.Query(q))

}
