package person

type MyPerson struct{
	Name string
	Age int
	Grade int
}

func (m MyPerson) SayHello()string{
	return "hello " + m.Name
}