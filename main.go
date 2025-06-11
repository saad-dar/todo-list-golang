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
	fmt.Println("###### Welcome to our Todo List App! ######")
	for i, item := range taskItems {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("#########################################")

	fmt.Println()
	fmt.Println("Work")
	fmt.Println("1. " + workOnFiverrProfile)
	fmt.Println("2. " + workOnFiverrPortfolio)

	fmt.Println()

	fmt.Println("Reward")
	fmt.Println("1. " + reward)

	fmt.Println()
	fmt.Println("Watch")
	fmt.Println("1. " + shortGoCourse)
	fmt.Println("2. " + pythonCourse)

}