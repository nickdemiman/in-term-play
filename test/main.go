package main

import (
	"fmt"
)

type ITest interface {
	Move()
	move()
}

type Parent struct {
	x, y int
}

type Child struct {
	Parent
}

func (p Parent) move() {
	p.Move()
}

func (p *Parent) Move() {
	fmt.Println("parent moved")
}

func (c *Child) Move() {
	fmt.Println("child moved")
}

func main() {
	child := Child{}

	child.move()
}
