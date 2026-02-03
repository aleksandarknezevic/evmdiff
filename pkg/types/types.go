package types

type Address string
type Slot string
type Word string
type Hash string

type Delta struct {
	Before string `json:"before"`
	After  string `json:"after"`
}
