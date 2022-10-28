package repository

import (
	"database/sql"
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
