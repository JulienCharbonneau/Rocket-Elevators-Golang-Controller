package main

import "fmt"


func main() {

	newButton := FloorRequestButton{}
	newButton.NewFloorRequestButton(1, "up")
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




