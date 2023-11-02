package model

type User struct {
	Id      int64   `validate:"gt=0"`
	Fname   string  `validate:"required"`
	City    string  `validate:"required"`
	Phone   string  `validate:"required"`
	Height  float64 `validate:"required"`
	Married bool
}
