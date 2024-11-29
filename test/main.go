package main

import "fmt"

type (
	SuperType interface {
		GetName() string
		GetMessage() string
		Overloaded()
	}

	Parent struct{}
	Child  struct {
		Parent
	}
)

func (p *Parent) based(s SuperType) {
	fmt.Println(s.GetName(), s.GetMessage())
}

func (p *Parent) GetName() string {
	return "parent"
}

func (p *Parent) GetMessage() string {
	return "hello world!"
}

func (p *Parent) Overloaded() {
	p.based(p)
}

func (c *Child) GetName() string {
	return "child"
}

func (c *Child) Overloaded() {
	c.based(c)
}

func main() {
	parent := Parent{}
	child := Child{}

	parent.Overloaded()
	child.Overloaded()
}
