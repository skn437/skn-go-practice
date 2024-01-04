package components

type ItemType struct {
	title string
	count uint8
	price float32
}

type BillType struct {
	name  string
	items []ItemType
}
