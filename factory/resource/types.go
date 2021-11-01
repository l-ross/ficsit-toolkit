package resource

type Amount struct {
	ClassName string
	Amount    int
}

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

type SchematicType int

const (
	Custom SchematicType = iota
	Cheat
	Tutorial
	Milestone
	Alternate
	Story
	MAM
	ResourceSink
	HardDrive
	Prototype
)

type EquipmentSlot int

const (
	None EquipmentSlot = iota
	Arms
	Back
)

type DroneDockingState int

const (
	Undocked DroneDockingState = iota
	Docking
	Docked
	Takeoff
)

type PowerPoleType int

const (
	Pole PowerPoleType = iota
	Wall
	WallDouble
)
