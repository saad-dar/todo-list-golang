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

}