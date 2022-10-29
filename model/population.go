package model

import "time"

type PopulationResponseItem struct {
	LocationID            string `db:"DistrictID"`
	Location              string `db:"Name"`
	PeopleWithRightToVote int64  `db:"HaveRight"`
	PeopleCommitTheVote   int64  `db:"Commits"`
	TotalPeople           int64  `db:"Total"`
}

type PopulationDatabaseRow struct {
	CitizenID   int       `db:"CitizenID"`
	LazerID     string    `db:"LazerID"`
	Name        string    `db:"Name"`
	Lastname    string    `db:"Lastname"`
	Birthday    time.Time `db:"Birthday"`
	Nationality string    `db:"Nationality"`
	DistrictID  string    `db:"DistrictID"`
}
