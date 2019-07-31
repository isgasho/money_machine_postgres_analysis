package main

import (
	"fmt"
	"reflect"
	"time"
)

type fn func(params ...interface{})

var cyclePool = []Cycle{}

func test(params ...interface{}) {
	listVal := reflect.ValueOf(params[0])
	for i := 0; i < listVal.Len(); i++ {
		// some_other_fun(listVal.Index(i).Interface())
		//Whatever function that accepts the params will have to reflect the param interface filtering accordingly
		fmt.Println(listVal.Index(i).Interface())
	}
	// fmt.Println(listVal.Index(i).Interface())
}

func createCycle(intervalSpeed int, amountOfInterval int, functionToCall fn, params ...interface{}) {
	//Identification matrix, with bool operator, pointer, list with dictionary values.
	// var i = 0

	// var boolOperate = true
	// for i > 1 {

	var cycleInstance = Cycle{
		Name:             "1",
		BooleanOperate:   true,
		IntervalSpeed:    intervalSpeed,
		AmountOfInterval: amountOfInterval,
		FunctionToCall:   functionToCall,
		Params:           params,
	}

	fmt.Println(cycleInstance.Params)
	cyclePool = append(cyclePool, cycleInstance)

	// }
	// var boolPointer *bool = &cycleInstance.BooleanPointer
	// fmt.Println(*cyclePool[0].BooleanPointer)
	// boolOperate = false
	// fmt.Println(*cyclePool[0].BooleanPointer)

	//startCycle

	// Handle at interval and interval speed
	//handle cancelation protocol
	// if boolOperate {
	// 	//goroutine
	// 	go functionToCall(params)
	// }

	// power limit BWAHAH HAHAHAHAHHAAHAHHAH
	// listVal := reflect.ValueOf(params)
	// for i := 0; i < listVal.Len(); i++ {
	// 	// some_other_fun(listVal.Index(i).Interface())
	// 	fmt.Println(listVal.Index(i).Interface())
	// }
}

func startCycle(cycleInstance *Cycle) {
	var i = 0
	//time cycle
	var functionToCall = cycleInstance.FunctionToCall
	var duration int = 2

	for i < cycleInstance.AmountOfInterval {
		if cycleInstance.BooleanOperate {
			functionToCall(cycleInstance.Params)
			time.Sleep(time.Duration(duration) * time.Second)
			if i == 2 {
				fmt.Println("hit")
				cancelCycle(cycleInstance)
			}
		}
		i++
	}
}

func cancelCycle(cycleInstance *Cycle) {
	fmt.Println("canceling")
	cycleInstance.BooleanOperate = false
}

func main() {
	createCycle(3, 10, test, "dog", "frog", 2, false)
	operatingCycle := cyclePool[0]
	go startCycle(&operatingCycle)

	fmt.Scanln()
	fmt.Println("done")
}
