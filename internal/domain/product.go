package domain

// Product represents an underlying URL with statistics on how it is used.
type Product struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Tipo  string `json:"type"`
	Count int    `json:"count"`
	Price int    `json:"price"`
}
