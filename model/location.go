package model

type District struct {
	DistrictID string `db:"DistrictID"`
	ProvinceID string `db:"ProvinceID"`
	Name       string `db:"Name"`
}
