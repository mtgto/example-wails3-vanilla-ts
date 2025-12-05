package main

import "time"

type GreetService struct{}

func (g *GreetService) Greet(name string) string {
	return "Hello " + name + "!"
}

func (g *GreetService) GetTime() time.Time {
	return time.Now()
}
