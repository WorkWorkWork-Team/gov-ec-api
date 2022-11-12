package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type submitMpRepository struct {
	mysql *sqlx.DB
}

type SubmitMpRepository interface {
	SubmitMpToDB(citizenID string) error
}

func NewSubmitMpRepository(mysql *sqlx.DB) SubmitMpRepository {
	return &submitMpRepository{
		mysql: mysql,
	}
}

func (a *submitMpRepository) SubmitMpToDB(citizenID string) error {
	rows, err := a.mysql.Query("INSERT INTO Mp (CitizenID) VALUES (?)", citizenID)
	if err != nil {
		logrus.Info("Have error when query that is: ", err)
		return err
	}
	rows.Close()
	return nil
}
