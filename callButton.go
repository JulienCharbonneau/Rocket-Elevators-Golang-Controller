package main

//Button on a floor or basement to go back to lobby
type CallButton struct {
	Id        int
	status    string
	floor     int
	direction string
}

func NewCallButton(_id int, _status string, _floor int, _direction string) *CallButton {
	callButton := CallButton{Id: _id, status: "on", floor: _floor, direction: _direction}
	return &callButton
}
