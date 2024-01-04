package utils

import "fmt"

func GetData() DataType {
	var user DataType = DataType{
		id:    "1",
		name:  "SKN",
		count: 7,
	}

	return user
}

func (u DataType) Format() string {
	var fs string = fmt.Sprintf("(%s) %s ....%v", u.id, u.name, u.count)

	return fs
}
