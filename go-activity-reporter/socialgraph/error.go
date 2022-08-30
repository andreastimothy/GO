package socialgraph

type UserDidNotExist struct {}

func (err UserDidNotExist) Error() string {
	return "User did not exist"
}