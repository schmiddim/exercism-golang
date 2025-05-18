package airportrobot

import "fmt"

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.
type Greeter interface {
	LanguageName() string
	Greet(string) string
}

type German struct {
}
type Italian struct {
}

type Portuguese struct {
}

func (p Portuguese) LanguageName() string {
	return "german"

}

func (p Portuguese) Greet(s string) string {
	return fmt.Sprintf("I can speak Portuguese: Olá %s!", s)

}

func (i Italian) LanguageName() string {
	return "italian"
}

func (i Italian) Greet(s string) string {
	return fmt.Sprintf("I can speak Italian: Ciao %s!", s)
}

func (g German) LanguageName() string {
	return "german"
}

func (g German) Greet(s string) string {
	return fmt.Sprintf("I can speak Germen: Hallo %s!", s)
}

func SayHello(name string, greeter Greeter) string {
	return greeter.Greet(name)

}
