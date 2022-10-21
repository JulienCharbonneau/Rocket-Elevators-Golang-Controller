package main

type Door struct {
	ID     int
	status string
}

func NewDoor(_ID int, _status string) Door {

	door := Door{ID: +_ID, status: _status}

	return door
}
