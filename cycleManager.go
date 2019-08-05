package main

import (
	"fmt"
	"reflect"
	"time"
)

func test(params ...interface{}) {
	listVal := reflect.ValueOf(params[0])
	for i := 0; i < listVal.Len(); i++ {
		fmt.Println(listVal.Index(i).Interface())
	}
}

func createCycle(intervalSpeed int, amountOfInterval int, functionToCall fn, params ...interface{}) {
	var cycleInstance = Cycle{
		Name:             "1",
		BooleanOperate:   true,
		IntervalSpeed:    intervalSpeed,
		AmountOfInterval: amountOfInterval,
		FunctionToCall:   functionToCall,
		Params:           params,
	}

	// fmt.Println(cycleInstance.Params)
	cyclePool = append(cyclePool, cycleInstance)
}

func startCycle(cycleInstance *Cycle) {
	var i = 0
	//time cycle
	var functionToCall = cycleInstance.FunctionToCall
	var duration = cycleInstance.AmountOfInterval
	var intervalSpeed = cycleInstance.IntervalSpeed
	fmt.Println("amount", duration)
	fmt.Println("interval speed", intervalSpeed)
	for i < duration {
		fmt.Println("i iteration ", i)
		if cycleInstance.BooleanOperate {
			functionToCall(cycleInstance.Params)
			time.Sleep(time.Duration(intervalSpeed) * time.Second)
		}
		i++
	}
}

func cancelCycle(cycleInstance *Cycle) {
	fmt.Println("canceling")
	cycleInstance.BooleanOperate = false
}

// func main() {
// 	createCycle(3, 10, test, "dog", "frog", 2, false)
// 	operatingCycle := cyclePool[0]
// 	go startCycle(&operatingCycle)

// 	fmt.Scanln()
// 	fmt.Println("done")
// }
