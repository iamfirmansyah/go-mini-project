package model

type Product struct {
	Id        int
	Name      string `validate:"required,min=3"`
	Price     int
	Stock     int
	CreatedAt int
	UpdatedAt int
}
