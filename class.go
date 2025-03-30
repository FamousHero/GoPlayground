package main

type Person struct {
	name string
	age int
	alive bool
}

func (p *Person) getName() string{
	return p.name
}