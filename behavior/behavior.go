package bahavior

type personBehavior struct {
	behaviorInterface behavior
}

type behavior interface {
	eat()
	drink()
	sleep()
}

func NewPerson(be behavior) *personBehavior {
	return &personBehavior{
		behaviorInterface: be,
	}
}
