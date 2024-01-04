package components

import "fmt"

func GetBill(name string) BillType {
	var items []ItemType

	var item1 ItemType = ItemType{
		title: "pizza",
		count: 23,
		price: 24.5,
	}

	var item2 ItemType = ItemType{
		title: "coke",
		count: 12,
		price: 12.8,
	}

	items = append(items, item1)
	items = append(items, item2)

	var bill BillType = BillType{
		name:  name,
		items: items,
	}

	return bill
}

func (bill BillType) Format() string { //* Receiver Function Only Available On An Object
	var formattedString string = fmt.Sprintf("%v Bill: \n", bill.name+"'s")

	var total float32 = 0

	for _, value := range bill.items {
		var itemTotalPrice float32 = value.price * float32(value.count)

		//* %-25 means 25 steps right & %25 means 25 steps left. Any other item after % specifier will relocate i.e. 25 steps left or right.
		formattedString += fmt.Sprintf("%-25v %v \n", value.title+":", itemTotalPrice) //* ":" in the specifier variable is must or it will go right 25 steps if mentioned separately

		total += itemTotalPrice
	}

	formattedString += fmt.Sprintf("%-25s %v \n", "Total:", total)

	return formattedString
}
