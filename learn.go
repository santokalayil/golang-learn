package main

import (
	"fmt"

	functions "github.com/santokalayil/golang-learn/functions"
	maps "github.com/santokalayil/golang-learn/maps"
	module "github.com/santokalayil/golang-learn/module"
	structs "github.com/santokalayil/golang-learn/structs"
)

func main() {
	message := module.ModuleFunction("Santo")
	fmt.Println(message)

	fmt.Println("----")
	fmt.Println("Functions >>>>>")
	function_output := functions.ParmeterizedFunction(20, 30)
	fmt.Println("Function output", function_output)

	fmt.Println("Structs >>>>>")
	structs.RunStructExamples()

	fmt.Println("Maps OR dictionaries>>>>")
	maps.RunMapsRelated()
}
