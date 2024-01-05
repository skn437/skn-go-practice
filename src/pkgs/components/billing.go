package components

import (
	"fmt"
)

func GetBill(name string) *BillType {
	var items []ItemType = []ItemType{}
	var tip float32 = 0.0

	var bill BillType = BillType{
		name:  name,
		items: items,
		tip:   tip,
	}

	return &bill
}

func (bill *BillType) AddItem(title string, count uint8, price float32) {
	var item *ItemType = &ItemType{
		title: title,
		count: count,
		price: price,
	}

	bill.items = append(bill.items, *item)
}

func (bill *BillType) SetTip(tip float32) {
	bill.tip = tip
}

func (bill *BillType) Format() string { //* Receiver Function Only Available On An Object
	var formattedString string = fmt.Sprintf("%v Bill: \n", bill.name+"'s")

	var total float32 = 0

	for _, value := range bill.items {
		var itemTotalPrice float32 = value.price * float32(value.count)

		//* %-25 means 25 steps right & %25 means 25 steps left. Any other item after % specifier will relocate i.e. 25 steps left or right.
		formattedString += fmt.Sprintf("%-25v $%v \n", value.title+":", itemTotalPrice) //* ":" in the specifier variable is must or it will go right 25 steps if mentioned separately

		total += itemTotalPrice
	}

	formattedString += fmt.Sprintf("%-25s $%v \n", "tip:", bill.tip)

	formattedString += fmt.Sprintf("%-25s $%v \n", "Total:", total+bill.tip)

	return formattedString
}
