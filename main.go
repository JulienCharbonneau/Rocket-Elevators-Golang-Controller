package main

import "fmt"

func main() {

	floorServed := [5]int{1, 2, 3, 4, 5}
	var newColumn = NewColumn(1, 2, floorServed[:], false)
	newColumn.requestElevator(2, "up")
	newColumn.requestElevator(10, "down")

	fmt.Println("best elevator", newColumn.best_elevator_informations.bestElevator.ID)
	// var newElevator = NewElevator(1, "idle", 5, 2)
	// fmt.Println("elevator", newElevator.ID)
	// newElevator.addNewRequest(5)
	// newElevator.addNewRequest(1)
	// newElevator.move()

	// var newCallButton = NewCallButton(2, "up")
	// fmt.Println("call button: ", newCallButton.direction)
	// fmt.Println("call button: ", newCallButton.floor)

	// var newDoor = NewDoor(1, "closed")
	// fmt.Println("door id: ", newDoor.ID)
	// fmt.Println("door status: ", newDoor.status)

	// var newButton = NewFloorRequestButton(1, "up")
	// newButton.floor = 1
	// fmt.Println(newButton.floor)

	// var newDoor = Door{ID: 1,status: "closed"}
	// fmt.Println(newDoor)
	// newDoor.status = "opend"
	// fmt.Println(newDoor)

	// scenarioNumber, err := strconv.Atoi(os.Args[1])
	// if err == nil {
	// 	runScenario(scenarioNumber)
	// } else {
	// 	fmt.Println(err)
	// }
}
