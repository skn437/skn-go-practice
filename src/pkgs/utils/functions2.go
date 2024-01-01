package utils

import "strings"

func GetInitials(name string) (string, string) {
	var newName string = strings.ToUpper(name)

	var names []string = strings.Split(newName, " ")

	var initials []string

	//* TODO
	//* append for int & string
	for _, value := range names {
		initials = append(initials, value[:1]) //* value[0] cannot be used, instead use value[:1]. As value is string. you can't use value[0] in string

	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}
