package model

type District struct {
	DistrictID int    `db:"DistrictID"`
	ProvinceID int    `db:"ProvinceID"`
	Name       string `db:"Name"`
}
