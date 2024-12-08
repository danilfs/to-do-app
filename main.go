package main

import (
	"fmt"
)

func main() {

	/*var taskOne string = "Watch go crush course"
	watchCourse := "Watch course"
	var rewardDesert string = "Eat cheesecake"
	// maxItemsInGroup := 10

	var taskItems = []string{taskOne, watchCourse, rewardDesert}
	// maxItems := 10

	fmt.Println("List of Todos")

	// for index, task := range taskItems {
	// 	// fmt.Println(index+1, ".", task)
	// 	fmt.Printf("%d: %s\n", index+1, task)
	// }

	/*fmt.Println(taskOne)
	fmt.Printf("%v %T \n", watchCourse, watchCourse)
	fmt.Println("Watch video")
	fmt.Println(rewardDesert)

	fmt.Println()
	fmt.Println("Tutorials")
	fmt.Println(taskOne)

	fmt.Println()
	fmt.Println("Reward")
	fmt.Println(rewardDesert)

	fmt.Println()
	fmt.Println("My Project") */
	// fmt.Println("Tasks:", taskItems)

	/*printTask(taskItems)
	fmt.Println()

	taskItems = addTask(taskItems, "Go for a run")
	taskItems = addTask(taskItems, "Practice coding in Go")

	fmt.Println()
	printTask(taskItems)*/
}

func printTask(items []string) {
	for index, task := range items {
		// fmt.Println(index+1, ".", task)
		fmt.Printf("%d: %s\n", index+1, task)
	}
}

func addTask(taskItems []string, newTask string) []string {
	var updatedTaskItems = append(taskItems, newTask)
	/*for index, tasks := range updatedTaskItems {
		fmt.Printf("%d: %s \n", index+1, tasks)
	}*/
	return updatedTaskItems
}
