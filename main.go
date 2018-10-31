package main

import (
	"net/http"

	"github.com/lambofgen/test_cal_weekday/controllers"
)

func main() {
	http.HandleFunc("/", controllers.IndexController)
	http.ListenAndServe(":8080", nil)
}
