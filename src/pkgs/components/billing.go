package components

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func getPrompt(prompt string, reader *bufio.Reader) (string, error) {
	fmt.Print(prompt)

	input, error := reader.ReadString('\n') //* You must use '' in ReadString. You can't use ""

	input = strings.TrimSpace(input)

	return input, error
}

func getBillOptions(bill *BillType) {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)

	const optionString string = `
	Choose an option:
	(a) Press 'a' to add an item
	(b) Press 't' to add a tip
	(c) Press 's' to save the bill
	Your Choice: `

	option, _ := getPrompt(optionString, reader)

	switch option {
	case "a":
		var item ItemType = ItemType{}

		title, _ := getPrompt("Enter Item Name: ", reader)
		item.title = title

		count, _ := getPrompt("Enter Item Count: ", reader)
		count64, err64 := strconv.ParseUint(count, 10, 8)

		if err64 != nil {
			log.Fatal(err64)
		}

		item.count = uint8(count64)

		price, _ := getPrompt("Enter Unit Price: ", reader)
		price64, errF64 := strconv.ParseFloat(price, 32)

		if errF64 != nil {
			log.Fatal(errF64)
		}

		item.price = float32(price64)

		bill.items = append(bill.items, item)

		getBillOptions(bill)
	case "t":
		tip, _ := getPrompt("Enter Your Tip: $", reader)

		tip64, errT64 := strconv.ParseFloat(tip, 32)

		if errT64 != nil {
			log.Fatal(errT64)
		}

		bill.tip = float32(tip64)

		getBillOptions(bill)
	case "s":
		saveBill(bill)
	default:
		fmt.Println("Invalid Selection!!! Try Again")
		getBillOptions(bill)
	}
}

func CreateBill() *BillType {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin) //* Very Important!!!

	name, _ := getPrompt("Create a name for the bill: ", reader)

	var bill *BillType = GetBill(name)

	getBillOptions(bill)

	fmt.Printf("New Bill Created: %v \n", *bill)

	return bill
}

func saveBill(bill *BillType) {
	var data []byte = []byte(bill.Format())

	var err error = os.WriteFile("files/"+bill.name+".txt", data, 0644) //* The location of created folder will be where `go.mod` is situated

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill Is Saved!")
}
