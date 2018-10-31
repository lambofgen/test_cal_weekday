package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/lambofgen/test_cal_weekday/services"
)

func IndexController(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// get ค่าจาก Form
		r.ParseForm()
		v := r.Form
		// แปลงค่า เป็น float64
		y, _ := strconv.ParseFloat(v.Get("year"), 64)
		m, _ := strconv.ParseFloat(v.Get("month"), 64)
		d, _ := strconv.ParseFloat(v.Get("day"), 64)
		// คำนวณ weekDay
		weekDay, err := services.CalWeekDay(y, m, d)
		if err != nil {
			renderTemplate(w, "index", err.Error())
			return
		}
		// render view พร้อม result
		renderTemplate(w, "index", "Weekday is "+weekDay)
	} else {
		// render view
		renderTemplate(w, "index", nil)
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, p interface{}) {
	t, _ := template.ParseFiles("views/" + tmpl + ".html")
	t.Execute(w, p)
}
