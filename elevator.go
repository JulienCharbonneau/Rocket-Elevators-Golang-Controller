package main

import (
	"fmt"
	"time"
)

type Elevator struct {
	ID                    int
	status                string
	amountOfFloors        int
	currentFloor          int
	door                  Door
	completedRequestsList []int
	floorRequestList      []int
	direction             string
}

func NewElevator(_elevatorID int, _status string, _amountOfFloors int, _currentFloor int) *Elevator {
	elevator := Elevator{ID: _elevatorID, status: _status, amountOfFloors: _amountOfFloors, currentFloor: _currentFloor}
	elevator.door = NewDoor(1, "closed")
	return &elevator
}

// func (e *Elevator) move() {

// }

func (e *Elevator) operateDoors() {
	e.door.status = "opened"
	fmt.Println("====Start operate doors======")
	time.Sleep(1 * time.Second)
	fmt.Println("open door")
	time.Sleep(2 * time.Second)
	fmt.Println("close door")
	time.Sleep(2 * time.Second)
	fmt.Println("=====End operate doors=====")

}

func (e *Elevator) addNewRequest(requestedFloor int) {
	if contains(e.floorRequestList, requestedFloor) {
	} else {
		e.floorRequestList = append(e.floorRequestList, requestedFloor)

	}
	if e.currentFloor < requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > requestedFloor {
		e.direction = "down"
	}

}
