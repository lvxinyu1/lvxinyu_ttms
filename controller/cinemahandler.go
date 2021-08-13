package controller

import (
	"TTMS/dao"
	"TTMS/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func ManageCinema(w http.ResponseWriter,r *http.Request)  {
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageCinema(pageNo)
	fmt.Println(page)
	t:=template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\cinema_edit.html"))
	t.Execute(w,page)
}
func AddCinema(w http.ResponseWriter,r *http.Request)  {
	CinemaId := r.FormValue("CinemaId")
	cinema, _ := dao.GetCinemaById(CinemaId)
	if cinema.CinemaId>0 {
		t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\cinema_edit2.html"))
		t.Execute(w, cinema)
	} else {
		t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\cinema_edit2.html"))
		t.Execute(w, "")
	}
}
func UpdateOrAddCinema(w http.ResponseWriter,r *http.Request) {
	CinemaId := r.PostFormValue("CinemaId")
	CinemaName := r.PostFormValue("name")
	CinemaSeatRow := r.PostFormValue("seatRow")
	CinemaSeatList := r.PostFormValue("seatList")
	icinemaID, _ := strconv.ParseInt(CinemaId, 10, 0)
	icinemaSeatRow, _ := strconv.ParseInt(CinemaSeatRow, 10, 0)
	icinemaSeatList, _ := strconv.ParseInt(CinemaSeatList, 10, 0)
	cinema := &model.Cinema{
		CinemaId:   int(icinemaID),
		CinemaName: CinemaName,
		SeatRow:    int(icinemaSeatRow),
		SeatList:   int(icinemaSeatList),
	}
	if cinema.CinemaId>0 {
		dao.UpdateCinema(cinema)
	} else {
		dao.AddCinema(cinema)
	}
	GetCinemaT(w,r)
}
func GetCinemaT(w http.ResponseWriter,r *http.Request)  {

	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageCinema(pageNo)
	fmt.Println(page)
	t:=template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\cinema_edit.html"))
	t.Execute(w,page)
}
func DeleteCinema(w http.ResponseWriter,r *http.Request)  {
	cinemaId :=r.FormValue("CinemaId")
	dao.DeleteCinema(cinemaId)
	GetCinemaT(w,r)

}