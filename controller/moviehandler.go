package controller

import (
	"TTMS/dao"
	"TTMS/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func BackstageManagement(w http.ResponseWriter,r *http.Request)  {
	t:=template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\manager.html"))
	t.Execute(w,"")
}
func GetCinema(w http.ResponseWriter,r *http.Request)  {

	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPageMovies(pageNo)
	fmt.Println(page)
	t:=template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\movie_edit.html"))
	t.Execute(w,page)
}
func ToUpdateMovieById(w http.ResponseWriter,r*http.Request)  {
	movieID := r.FormValue("movieId")
	movie, _ := dao.GetMovieById(movieID)
	if movie.MovieID>0 {
		t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\movie_edit2.html"))
		t.Execute(w, movie)
	} else {
		t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\movie_edit2.html"))
		t.Execute(w, "")
	}
}
func UpdateOrAddMovie(w http.ResponseWriter,r *http.Request)  {
	movieID :=r.PostFormValue("movieId")
	movieName:=r.PostFormValue("name")
	duration:=r.PostFormValue("duration")
	movieNumber:=r.PostFormValue("number")
	movieInformation:=r.PostFormValue("information")
	imgAddress:=r.PostFormValue("imgaddress")
	imovieID, _ := strconv.ParseInt(movieID, 10, 0)
	movie := &model.Movie{
		MovieID :  int(imovieID),
		MovieName: movieName,
		Duration :duration,
		MovieNumber :movieNumber,
		MovieInformation :movieInformation,
		ImgAddress :imgAddress,
	}
	if movie.MovieID>0 {
		dao.UpdateMovie(movie)
	} else {
		dao.AddMovie(movie)
	}
	GetCinema(w, r)
}
func DeleteMovie(w http.ResponseWriter,r *http.Request)  {
	movieID :=r.FormValue("movieId")
	//movieName:=r.PostFormValue("name")
	//cinema:=r.PostFormValue("cinema")
	//movieNumber:=r.PostFormValue("number")
	//movieInformation:=r.PostFormValue("information")
	//imgAddress:=r.PostFormValue("imgaddress")
	fmt.Println(movieID)
	dao.DeleteMovie(movieID)
	GetCinema(w, r)
}