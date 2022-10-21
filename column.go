package main

import "fmt"

type Column struct {
	ID                         int
	status                     string
	amountOfFloors             int
	amountOfElevators          int
	callButtonsList            []CallButton
	servedFloorsList           []int
	elevatorList               []Elevator
	isBasement                 bool
	best_elevator_informations BestElevatorInformations
}

func NewColumn(_id, _amountOfElevators int, _servedFloors []int, _isBasement bool) *Column {
	column := Column{ID: _id, amountOfElevators: _amountOfElevators, servedFloorsList: _servedFloors, isBasement: _isBasement}
	column.amountOfFloors = len(_servedFloors)
	column.best_elevator_informations = best_elevator_Info()
	column.createCallButtons(column.amountOfFloors, column.isBasement)
	column.createElevators(column.amountOfFloors, column.amountOfElevators)
	return &column
}

func (column *Column) createCallButtons(amountOfFloors int, isBasement bool) {
	if column.isBasement {
		buttonFloor := -1
		callButtonID := 1
		for i := 0; i < column.amountOfFloors; i++ {
			callButton := NewCallButton(callButtonID, "off", buttonFloor, "up")
			column.callButtonsList = append(column.callButtonsList, callButton)
			buttonFloor--
			callButtonID++
		}
	} else {
		buttonFloor := 1
		callButtonID := 1
		for i := 0; i < column.amountOfFloors; i++ {
			callButton := NewCallButton(callButtonID, "off", buttonFloor, "down")
			column.callButtonsList = append(column.callButtonsList, callButton)
			buttonFloor++
			callButtonID++
		}
	}

}

func (column *Column) createElevators(_amountOfFloors int, _amountOfElevators int) {
	elevatorID := 1
	for i := 0; i < column.amountOfElevators; i++ {
		elevator := NewElevator(elevatorID, "idle", _amountOfElevators, 1)
		column.elevatorList = append(column.elevatorList, *elevator)
	}

}

// //Simulate when a user press a button on a floor to go back to the first floor
func (c *Column) requestElevator(_requestedFloor int, _direction string) *Elevator {
	elevator := c.findElevator(_requestedFloor, _direction)
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	elevator.addNewRequest(1)
	elevator.move()
	return &elevator
}

func (column *Column) findElevator(requestedFloor int, requestedDirection string) Elevator {
	if requestedFloor == 1 {
		for _, elevator := range column.elevatorList {
			if elevator.currentFloor == 1 && elevator.status == "stopped" {
				column.best_elevator_informations = *column.checkIfElevatorIsBetter(1, elevator, column.best_elevator_informations, requestedFloor)
			} else if elevator.currentFloor == 1 && elevator.status == "idle" {
				column.best_elevator_informations = *column.checkIfElevatorIsBetter(2, elevator, column.best_elevator_informations, requestedFloor)

			} else if elevator.currentFloor < 1 && elevator.direction == "up" {
				column.best_elevator_informations = *column.checkIfElevatorIsBetter(3, elevator, column.best_elevator_informations, requestedFloor)

			} else if elevator.currentFloor > 1 && elevator.direction == "down" {
				column.best_elevator_informations = *column.checkIfElevatorIsBetter(3, elevator, column.best_elevator_informations, requestedFloor)

			} else if elevator.status == "idle" {
				column.best_elevator_informations = *column.checkIfElevatorIsBetter(4, elevator, column.best_elevator_informations, requestedFloor)

			} else {
				column.best_elevator_informations = *column.checkIfElevatorIsBetter(5, elevator, column.best_elevator_informations, requestedFloor)

			}
			fmt.Println(elevator)
		}
	} else {
		for _, elevator := range column.elevatorList {
			if requestedFloor == elevator.currentFloor && elevator.status == "stopped" && requestedDirection == elevator.direction {
				column.best_elevator_informations = *column.checkIfElevatorIsBetter(1, elevator, column.best_elevator_informations, requestedFloor)

			} else if requestedFloor > elevator.currentFloor && elevator.direction == "up" && requestedDirection == "up" {
				column.best_elevator_informations = *column.checkIfElevatorIsBetter(2, elevator, column.best_elevator_informations, requestedFloor)

			} else if requestedFloor < elevator.currentFloor && elevator.direction == "down" && requestedDirection == "down" {
				column.best_elevator_informations = *column.checkIfElevatorIsBetter(2, elevator, column.best_elevator_informations, requestedFloor)

			} else if elevator.status == "idle" {
				column.best_elevator_informations = *column.checkIfElevatorIsBetter(4, elevator, column.best_elevator_informations, requestedFloor)

			} else {
				column.best_elevator_informations = *column.checkIfElevatorIsBetter(1, elevator, column.best_elevator_informations, requestedFloor)

			}
		}
	}
	return column.best_elevator_informations.bestElevator
}

func (column *Column) checkIfElevatorIsBetter(scoreToCheck int, new_Elevator Elevator, info BestElevatorInformations, floor int) *BestElevatorInformations {
	if scoreToCheck < info.bestScore {
		info.bestScore = scoreToCheck
		info.bestElevator = new_Elevator
		info.referenceGap = Abs(new_Elevator.currentFloor - floor)
	} else if scoreToCheck == info.bestScore {
		gap := Abs(new_Elevator.currentFloor - floor)
		if info.referenceGap > gap {
			info.bestElevator = new_Elevator
			info.referenceGap = gap
		}
	}
	return &info
}

type BestElevatorInformations struct {
	bestScore    int
	referenceGap int
	bestElevator Elevator
}

func best_elevator_Info() BestElevatorInformations {
	info := BestElevatorInformations{referenceGap: 10000000}
	return info
}
