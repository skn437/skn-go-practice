package components

type ItemType struct { //* When using `struct`, it's better to use pointers
	title string
	count uint8
	price float32
}

type BillType struct {
	name  string
	items []ItemType
	tip   float32
}
