package model

import "time"

type PopulationResponseItem struct {
	LocationID            int
	Location              string
	PeopleWithRightToVote int64
	PeopleCommitTheVote   int64
	TotalPeople           int64
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
