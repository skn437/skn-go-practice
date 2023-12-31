package user

type UserType struct {
	firstName string
	lastName  string
	bookCount uint8
}

func init() {
	var _ UserType = UserType{
		firstName: "",
		lastName:  "",
		bookCount: 0,
	}
}
