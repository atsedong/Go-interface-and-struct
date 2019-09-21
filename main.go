package main

import "fmt"

type car interface {
	Name() string
}

type bmw struct {
	ID           int
	Colour       string
	Year         string
	Manufacturer string
	Country      string
}

type nissan struct {
	ID           int
	Colour       string
	Year         string
	Manufacturer string
	Country      string
	Owner        string
}

func (b bmw) Name() string {
	return "BMW"
}

func (n nissan) Name() string {
	return "Nissan"
}

func do(i car) {
	switch v := i.(type) {
	case bmw:
		fmt.Printf("This is BMW, the colour is %v\n", v.Colour)
	case nissan:
		fmt.Printf("This is Nissan, the colour is %v and the owner of this Nissan is %v\n", v.Colour, v.Owner)
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
