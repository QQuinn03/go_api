package bahavior

type personBehavior struct {
	behaviorInterface Behavior
}

type Behavior interface {
	Eat()
	Drink()
	Sleep()
}

func (p *personBehavior) Eat() {
	p.behaviorInterface.Eat()

}
func (p *personBehavior) Drink() {
	p.behaviorInterface.Drink()

}
func (p *personBehavior) Sleep() {
	p.behaviorInterface.Sleep()

}

func NewBhavior(be Behavior) *personBehavior {
	return &personBehavior{
		behaviorInterface: be,
	}
}
