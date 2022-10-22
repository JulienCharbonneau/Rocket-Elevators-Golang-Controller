# Rocket-Elevators-Golang-Controller
This project is about implementing an elevator controller for commercial purposes.  The program is based on a pseudocode file given and for this version written in Golang. Golang ist a compiled language  designed by Google and because ist a compiled language it is faster than an interpreted language but handles less abstraction.Here a [video](https://drive.google.com/file/d/13dNLaarjPkoS8UFG3nO1uvWlKZEHt5lP/view?usp=sharing) that I explain my work

### Usage 
To run the script with golang run the command in the 
`go run`


### Running the tests

To launch the tests, make sure to be at the root of the repository and run:

`go test`

![](https://github.com/JulienCharbonneau/Rocket-Elevators-Golang-Controller/blob/main/goTest.gif)


## Description
This program creates a number of columns and elevators as needed and supports the needs of elevator request button and floor access request button with a system-based efficiency management point allowing to evaluate the best choice taking into account the floor where the request was initiated versus the availability and the direction of the cage. For commercial buildings, the first column is assigned at the basement and others are split between floors to serve.



