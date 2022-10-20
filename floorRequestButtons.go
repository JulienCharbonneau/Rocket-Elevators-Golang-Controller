package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

func NewFloorRequestButton(_floor int, _direction string) FloorRequestButton {

	buttton := FloorRequestButton{ID: 1, status: "online", floor: _floor, direction: _direction}

	return buttton
}
