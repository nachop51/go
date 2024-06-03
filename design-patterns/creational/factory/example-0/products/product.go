package prodcuts

import "time"

type Product struct {
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) New() *Product {
	return &Product{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
