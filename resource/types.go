package resource

type StackSize int

const (
	One StackSize = iota
	Small
	Medium
	Big
	Huge
	Fluid
)

type Form int

const (
	Invalid Form = iota
	Solid
	Liquid
	Gas
	Heat
)
