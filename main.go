package main

import "fmt"

func main() {

	var newButton = NewFloorRequestButton(1, "up")
	// // newButton := FloorRequestButton{}
	// newButton.floor = 1
	fmt.Println(newButton.floor)

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
