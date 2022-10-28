package model

type PopulationResponseItem struct {
	LocationID            string
	Location              string
	PeopleWithRightToVote int64
	PeopleCommitTheVote   int64
	TotalPeople           int64
}
