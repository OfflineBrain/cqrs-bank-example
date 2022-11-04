package valueobject

type ProductStatus string

const (
	OnSale      ProductStatus = "on_sale"
	SoldOut                   = "sold_out"
	Unavailable               = "unavailable"
)
