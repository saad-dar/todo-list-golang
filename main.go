package main

import (
	"fmt"
	"net/http"
)
var shortGoCourse = "Watch Go crash course"
var pythonCourse = "Watch Python crash course"
var workOnFiverrProfile = "Work on Fiverr profile"
var workOnFiverrPortfolio = "Work on Fiverr portfolio"
var workOnFiverrGig = "Work on Fiverr gig"
var reward = "Reward myself with a fruit juice"

var taskItems = []string{ shortGoCourse, pythonCourse, workOnFiverrProfile, workOnFiverrPortfolio, workOnFiverrGig, reward }

func main() {

	fmt.Println("###### My Todo List ! ######")

	http.HandleFunc("/", helloUser)
	http.HandleFunc("/show-tasks", showTasks)
	http.ListenAndServe(":8080", nil)
}

func helloUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, User! Welcome to your Todo List!")
}

func showTasks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "###### My Todo List ! ######\n")
	fmt.Fprintf(w, "#########################################\n")
	for i, item := range taskItems {
		fmt.Fprintf(w, "%d. %s\n", i+1, item)
	}
	fmt.Fprintf(w, "#########################################\n")
}
