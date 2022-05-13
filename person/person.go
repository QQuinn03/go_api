package person

type person struct {
	todo behavior
}

type behavior interface {
	eat()
	drink()
	sleep()
}

func NewPerson(be behavior) *person {
	return &person{
		todo: be,
	}
}
