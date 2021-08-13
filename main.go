package main

import (
	"TTMS/controller"
	"net/http"
)


func main() {
	http.Handle("/views/", http.StripPrefix("/views/", http.FileServer(http.Dir("src/TTMS/views/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("src/TTMS/views"))))
	http.Handle("/static/",http.StripPrefix("/static/",http.FileServer(http.Dir("src/TTMS/views/static/"))))
	http.HandleFunc("/login2", controller.Login2)
	http.HandleFunc("/login", controller.Login)
	http.HandleFunc("/main", controller.Homepage)
	http.HandleFunc("/regist", controller.Regist)
	http.HandleFunc("/regist2", controller.Regist2)
	http.HandleFunc("/getCinema", controller.GetCinema)
	http.HandleFunc("/getCinemaT", controller.GetCinemaT)
	http.HandleFunc("/manageCinema", controller.ManageCinema)
	http.HandleFunc("/backstageManagement", controller.BackstageManagement)
	http.HandleFunc("/logout", controller.Logout)
	http.HandleFunc("/performancePlan", controller.PerformancePlan)
	http.HandleFunc("/revisePlan", controller.RevisePlan)
	http.HandleFunc("/addPlan", controller.AddPlan)
	http.HandleFunc("/addCinema", controller.AddCinema)
	http.HandleFunc("/updateOrAddCinema", controller.UpdateOrAddCinema)
	http.HandleFunc("/updateOrAddPlan", controller.UpdateOrAddPlan)
	http.HandleFunc("/deleteMovie", controller.DeleteMovie)
	http.HandleFunc("/deletePlan", controller.DeletePlan)
	http.HandleFunc("/buytickets",controller.BuyTickets)
	http.HandleFunc("/toGetTicket",controller.ToGetTicket)
	http.HandleFunc("/getTic",controller.GetTic)
	http.HandleFunc("/updateOrAddMovie", controller.UpdateOrAddMovie)
	http.HandleFunc("/toUpdateMovieById", controller.ToUpdateMovieById)
	http.HandleFunc("/deleteCinema", controller.DeleteCinema)
	http.ListenAndServe(":8080", nil)
}