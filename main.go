package main

import "fmt"

func main() {

	var shortGoCourse = "Watch Go crash course"
	var pythonCourse = "Watch Python crash course"
	var workOnFiverrProfile = "Work on Fiverr profile"
	var workOnFiverrPortfolio = "Work on Fiverr portfolio"
	var workOnFiverrGig = "Work on Fiverr gig"
	var reward = "Reward myself with a fruit juice"

	var taskItems = []string{ shortGoCourse, pythonCourse, workOnFiverrProfile, workOnFiverrPortfolio, workOnFiverrGig, reward }
	// Print the welcome message and todo list
	printTasks(taskItems)
	// Add a new task
	fmt.Println()
	addTask(taskItems, "Go for a run")
}

func printTasks (taskItems []string) {
	fmt.Println("###### Welcome to our Todo List App! ######")
	for i, item := range taskItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("#########################################")

}

func addTask(taskItems []string, newTask string) {
	taskItems = append(taskItems, newTask)
	printTasks(taskItems)
}