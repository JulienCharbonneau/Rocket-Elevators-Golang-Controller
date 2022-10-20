package main

//FloorRequestButton is a button on the pannel at the lobby to request any floor
type FloorRequestButton struct {
	ID int
	status string
	floor int
	direction string
}

func(requestButton FloorRequestButton) NewFloorRequestButton(_floor int, _direction string) *FloorRequestButton {
	requestButton.floor = _floor
	requestButton.direction =_direction
	buttton := FloorRequestButton{ID: 1,status: "online",floor: requestButton.floor,direction: requestButton.direction}

	return &buttton
}

