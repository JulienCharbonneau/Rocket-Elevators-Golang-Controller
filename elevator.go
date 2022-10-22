package main

import (
	"fmt"
	"sort"
	"time"
)

type Elevator struct {
	ID                    int
	status                string
	amountOfFloors        int
	currentFloor          int
	door                  Door
	completedRequestsList []int
	floorRequestsList     []int
	direction             string
}

func NewElevator(_elevatorID int, _status string, _amountOfFloors int, _currentFloor int) *Elevator {
	elevator := Elevator{ID: _elevatorID, status: _status, amountOfFloors: _amountOfFloors, currentFloor: _currentFloor}
	elevator.door = NewDoor(1, "closed")
	return &elevator
}

func (e *Elevator) move() {
	for len(e.floorRequestsList) > 0 {
		fmt.Println("========Elevator is moving==========")
		e.status = "moving"
		time.Sleep(2 * time.Second)

		e.sortFloorList()
		destination := e.floorRequestsList[0]
		if e.direction == "up" {
			for e.currentFloor < destination {
				e.currentFloor++
			}

		} else if e.direction == "down" {
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "stopped"
		fmt.Println("===== Elevator stopped=======")
		fmt.Println("request: ", e.floorRequestsList)
		fmt.Println("current floor:", e.currentFloor)
		e.operateDoors()

		e.floorRequestsList = e.floorRequestsList[1:]
		e.completedRequestsList = append(e.completedRequestsList, destination)

	}
	e.status = "idle"
	e.direction = "empty"
	fmt.Println("========Elevator stopped==========")

}

func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Ints(e.floorRequestsList)

	} else {
		sort.Sort(sort.Reverse(sort.IntSlice(e.floorRequestsList)))
	}
}

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
	if contains(e.floorRequestsList, requestedFloor) {

	} else {
		e.floorRequestsList = append(e.floorRequestsList, requestedFloor)

	}
	if e.currentFloor < requestedFloor {
		e.direction = "up"
	}
	if e.currentFloor > requestedFloor {
		e.direction = "down"
	}

}
