package main

import "fmt"

func main() {

	var newElevator = NewElevator(1, "idle", 5, 2)
	fmt.Println("elevator", newElevator.ID)
	fmt.Println("elevator door", newElevator.door.status)
	newElevator.door.status = "opend"
	fmt.Println("elevator door", newElevator.door.status)
	newElevator.floorRequestList = append(newElevator.floorRequestList, 1, 2, 3)
	fmt.Println("request list:", newElevator.floorRequestList)

	newElevator.addNewRequest(5)

	newElevator.operateDoors()
	fmt.Println("request list:", newElevator.floorRequestList)

	fmt.Println("elevator direction :", newElevator.direction)

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
