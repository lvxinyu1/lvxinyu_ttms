package controller

import (
	"TTMS/dao"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
)

func ToGetTicket(w http.ResponseWriter, r *http.Request) {
	movieID := r.FormValue("movieId")
	Movie, _ := dao.GetMovieById(movieID)
	Plans, _ := dao.GetPlanByName(Movie.MovieName)
	t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\buyticket.html"))
	t.Execute(w, Plans)
}
func BuyTickets(w http.ResponseWriter, r *http.Request) {
	planId := r.FormValue("planId")
	fmt.Println("太傻逼了",planId)
	sit := dao.GetSitById(planId)
	t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\users\\Buytickets.html"))
	t.Execute(w, sit)
}
func GetTic(w http.ResponseWriter, r *http.Request) {
	//model.DeleteTicByFlag()
	//tickets := make([]model.Ticket, 0)
	//监听是否为POST方法
		b, err := ioutil.ReadAll(r.Body)
	//	_, sess := dao.IsLogin(r)
		PlanId := r.FormValue("planId")
	//

		if err != nil {
			fmt.Println(err)
		}
	//	//fmt.Println(b)

		for i := 10; i < len(b); i += 11 {

			r := int64(b[i] - '0')
			l := int64(b[i+1] - '0')
	//orderid := utils.CreateUUID()
			iplanId,_:=strconv.ParseInt(PlanId,10,0)
			fmt.Println("傻逼项目",iplanId,r,l)
			ticket:=dao.GetTicketByRowAndList2(int(iplanId),int(r),int(l))
			ticket.State=1
			fmt.Println("打印",ticket)
			dao.UpdateTicket(ticket)
		}

	BuyTickets(w,r)
	//err:=
	//if err != nil {
	//	BuyTickets(w,r)
	//}
	//if r.Method == "POST" {
	//	b, _ := ioutil.ReadAll(r.Body)
	//	fmt.Println(string(b))
	//}
	//body = bytes.TrimPrefix(r.Body, []byte("\xef\xbb\xbf")) // Or []byte{239, 187, 191}
	//var num []struct {
	//	Id int `json:"id,omitempty"`
	//}
	//err := json.NewDecoder(r.Body).Decode(&num)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(num)
}
git  
