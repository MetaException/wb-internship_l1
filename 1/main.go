package main

import "fmt"

type Human struct {
	Name   string
	Age    int
	Gender rune
}

func (human *Human) jump() {
	fmt.Printf("[%s][%v][%c] Jumping\n", human.Name, human.Age, human.Gender)
}

type Action struct {
	Human
}

func main() {
	a := Action{
		Human: Human{
			Name:   "13",
			Age:    12,
			Gender: 'M',
		},
	}

	a.Age = 21
	a.Gender = 'M'
	a.Name = "Eugene"

	a.jump()
}
