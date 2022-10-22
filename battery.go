package main

import (
	"math"
)

type Battery struct {
	Id                       int
	status                   string
	columnsList              []Column
	floorRequestsButtonsList []FloorRequestButton
	amountOfColumns,
	amountOfFloors,
	amountOfBasements,
	amountOfElevatorPerColumn,
	columnID,
	floorRequestButtonID int
}

func NewBattery(_id, _amountOfColumns, _amountOfFloors, _amountOfBasements, _amountOfElevatorPerColumn int) *Battery {
	battery := Battery{Id: _id, status: "online", amountOfColumns: _amountOfColumns, amountOfFloors: _amountOfFloors, amountOfBasements: _amountOfBasements, amountOfElevatorPerColumn: _amountOfElevatorPerColumn}

	if _amountOfBasements > 0 {
		battery.createFloorRequestButtons(_amountOfBasements)
		battery.createBasementColumn(_amountOfBasements, _amountOfElevatorPerColumn)
		_amountOfColumns--
	}
	battery.createColumns(float64(_amountOfColumns), float64(_amountOfFloors), _amountOfElevatorPerColumn)
	return &battery
}

func (b *Battery) createBasementColumn(amountOfBasements int, amountOfElevatorPerColumn int) {
	servedFloorsList := []int{}
	floor := -1
	for i := 0; i < amountOfBasements; i++ {
		servedFloorsList = append(servedFloorsList, floor)
		floor--
	}
	column := NewColumn(b.Id, "online", amountOfBasements, amountOfElevatorPerColumn, servedFloorsList, true)
	b.columnsList = append(b.columnsList, *column)
	b.columnID++
}

func (b *Battery) createColumns(amountOfColumns float64, amountOfFloors float64, amountOfElevatorPerColumn int) {
	amountOfFloorsPerColumn := math.Round(amountOfFloors / amountOfColumns)
	floor := 1
	for i := 0; i < int(amountOfColumns); i++ {
		servedFloorsList := []int{}
		for j := 0; j < int(amountOfFloorsPerColumn); j++ {
			if floor <= int(amountOfFloors) {
				servedFloorsList = append(servedFloorsList, floor)
				floor++
			}
		}
		column := NewColumn(b.columnID, "online", int(amountOfFloors), amountOfElevatorPerColumn, servedFloorsList, false)
		b.columnsList = append(b.columnsList, *column)
		b.columnID++

	}
}

func (b *Battery) createFloorRequestButtons(amountOfFloors int) {
	buttonFloor := 1
	for i := 0; i < amountOfFloors; i++ {
		floorRequestButton := NewFloorRequestButton(b.floorRequestButtonID, "off", buttonFloor, "up")
		b.floorRequestsButtonsList = append(b.floorRequestsButtonsList, floorRequestButton)
		b.floorRequestButtonID++
		buttonFloor++
	}
}

func (b *Battery) createBasementFloorRequestButtons(amountOfBasements int) {
	buttonFloor := -1
	for i := 0; i < amountOfBasements; i++ {
		floorRequestButton := NewFloorRequestButton(b.floorRequestButtonID, "off", buttonFloor, "down")
		b.floorRequestsButtonsList = append(b.floorRequestsButtonsList, floorRequestButton)
		b.floorRequestButtonID++
		buttonFloor--

	}
}

func (b *Battery) findBestColumn(_requestedFloor int) Column {
	for _, column := range b.columnsList {
		if contains(column.servedFloorsList, _requestedFloor) {
			return column
		}
	}
	return b.columnsList[0]
}

//Simulate when a user press a button at the lobby
func (b *Battery) assignElevator(_requestedFloor int, _direction string) (*Column, *Elevator) {
	column := b.findBestColumn(_requestedFloor)
	elevator := column.findElevator(1, _direction)
	elevator.addNewRequest(1)
	elevator.move()
	elevator.addNewRequest(_requestedFloor)
	elevator.move()
	return &column, elevator
}
