package controller

import (
	"TTMS/dao"
	"TTMS/model"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func PerformancePlan(w http.ResponseWriter,r *http.Request)  {
	pageNo := r.FormValue("pageNo")
	//获取价格范围
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	if pageNo == "" {
		pageNo = "1"
	}
	var page *model.Page
	if minPrice == "" && maxPrice == "" {
		//调用bookdao中获取带分页的图书的函数
		page, _ = dao.GetPageMovies(pageNo)
	}
	//} else {
	//	//调用bookdao中获取带分页和价格范围的图书的函数
	//	page, _ = dao.GetPageBooksByPrice(pageNo, minPrice, maxPrice)
	//	//将价格范围设置到page中
	//	page.MinPrice = minPrice
	//	page.MaxPrice = maxPrice
	//}
	flag, session := dao.IsLogin(r)

	if flag {
		//已经登录，设置page中的IsLogin字段和Username的字段值
		page.IsLogin = true
		page.UserName = session.UserName
	}
	fmt.Println(page)
	t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\plan.html"))
	t.Execute(w,page)
}
//修改演出计划
func RevisePlan(w http.ResponseWriter,r *http.Request){
	pageNo := r.FormValue("pageNo")
	if pageNo == "" {
		pageNo = "1"
	}
	page, _ := dao.GetPagePlan(pageNo)
	fmt.Println(page)
	t:=template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\reviseplan.html"))
	t.Execute(w,page)
}
//增加排片
func AddPlan(w http.ResponseWriter,r *http.Request)  {

	planid:=r.FormValue("PlanId")
	//moviename := r.FormValue("MovieName")
	plan,_:=dao.GetPlanById(planid)
	fmt.Println("排片序号是：",planid)
		t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\manager\\addplan.html"))
		t.Execute(w, plan)

}
func UpdateOrAddPlan(w http.ResponseWriter,r *http.Request)  {
	moviename :=r.PostFormValue("name")
	cinemaId:=r.PostFormValue("cinemaid")
	plantime:=r.PostFormValue("plantime")
	plan1,_:=dao.GetPlanById2()
	planid:=plan1.PlanId
	fmt.Println("计划为",planid)
	iplanId,_:=strconv.ParseInt(string(planid),10,0)
	icinemaId,_:=strconv.ParseInt(cinemaId,10,0)
	TDate:=dao.ChangeTime(plantime)
	cinemaid:=r.FormValue("cinemaid")
	iScreen,_:=strconv.ParseInt(cinemaid,10,0)
	screen,_:=dao.GetCinemaById(cinemaid)
	badSeatCount,badSeat:=dao.GetBadSeatByCinemaID(cinemaid)
	//把坏座位加入哈希表里
	MapLine:=make(map[int]bool,badSeatCount)
	MapRow:=make(map[int]bool,badSeatCount)
	fmt.Println("planid",planid)
	for _,v:=range badSeat{

		MapLine[v.List]=true
		MapRow[v.Row]=true
	}

	var Tickets []*model.Ticket
	for i:=1;i<=screen.SeatList;i++{
		for j:=1;j<=screen.SeatRow;j++{
			if !(MapLine[i]==true&&MapRow[j]==true){
				//如果影厅座位的行数和列数都满足哈希表条件，就生成该座位的票
				ticket:=&model.Ticket{
					MovieName: moviename,
					PlanId: planid,
					CinemaId:  int(iScreen),
					List:      i,
					Row:       j,
					State:     0,
					PlanTime:  TDate,
				}
				Tickets=append(Tickets,ticket)
			}
		}
	}
	//将票切片保存
	dao.PreserveTicket(Tickets)
	plan := &model.Plan{
		PlanId: int(iplanId),
		MovieName: moviename,
		CinemaId : int(icinemaId),
		PlanTime: plantime,
	}
	if plan.PlanId>0 {
		dao.UpdatePlan(plan)
	} else {
		dao.AddPlan(plan)
	}
	RevisePlan(w,r)
}
func DeletePlan(w http.ResponseWriter,r *http.Request)  {
	planid :=r.FormValue("PlanId")
	dao.DelatePlanById(planid)
	RevisePlan(w,r)
}
