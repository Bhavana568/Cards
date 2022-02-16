//go run- compiles and executes files in the folder
//go build- compiles the file
//go fmt: formats file

package main

//package = project
//in gotere are 2 types of package: executable and reusable
//executable genrates a file that we can run, reusable used as helper to put a reusable logic
//package main is executable package

import "fmt"

//import fmt gives access to package of the file and code
//fmt is format

func main() {
	fmt.Println("Hi there!")
}

//func tells go we are abt to declare a func
//main sets a name
//() : list of args

//organization of code in a file: package->import->fun->func body
