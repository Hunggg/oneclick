package entity

type Products struct {
	ProductId int64
	ProductName string
	SupplierId int64
	CategoryId int64
	QuantityPerUint string
	UintPrice int32
	UintsInStock int32
	UintsOnOrder int32
	ReorderLevel int32
	Discontinued int32
}