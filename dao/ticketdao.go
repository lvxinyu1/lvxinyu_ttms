package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"fmt"
	"strconv"
	"time"
)

func GetPlanTimeByPlanId(PlanId string)(string,error){
	sqlStr := "select plan_id,cinema_id,movie_name,plan_time from plan where plan_id = ?"
	row := utils.Db.QueryRow(sqlStr, PlanId)
	fmt.Println(222,row)
	plan := &model.Plan{}
	row.Scan(&plan.PlanId,&plan.CinemaId,&plan.MovieName,&plan.PlanTime)
	return plan.PlanTime, nil
}
func ChangeTime(Time string)string{
	var times []byte
	for _,v:=range Time{
		if v!=84{
			times=append(times,byte(v))
		}else{
			times=append(times,32)
		}
	}
	t1:=time.Now().Second()
	t:=strconv.FormatInt(int64(t1),10)
	return string(times)+":"+t
}
func GetSitById(planid string) *model.Sit {
	sit:= &model.Sit{}
	sit.Perf,_=GetPlanById(planid)
	iplanId,_:=strconv.ParseInt(planid,10,0)
	sit.Perf.Movie,_=GetMovieByName(sit.Perf.MovieName)
	sit.SitOne=make([]model.Seat,7)
	sit.SitTwo=make([]model.Seat,7)
	sit.SitThree=make([]model.Seat,7)
	sit.SitFour=make([]model.Seat,7)
	sit.SitFive=make([]model.Seat,7)
	sit.SitSix=make([]model.Seat,7)
	sit.SitSeven=make([]model.Seat,7)
	for i:=0;i<7;i++{
		sit.SitOne[i].Row=1
		sit.SitOne[i].Line=i+1
		sit.SitOne[i].Flag=GetTicketByRowAndList(int(iplanId),1,i+1)
	}
	for i:=0;i<7;i++{
		sit.SitTwo[i].Row=2
		sit.SitTwo[i].Line=i+1
		sit.SitTwo[i].Flag=GetTicketByRowAndList(int(iplanId),2,i+1)
	}
	for i:=0;i<7;i++{
		sit.SitThree[i].Row=3
		sit.SitThree[i].Line=i+1
		sit.SitThree[i].Flag=GetTicketByRowAndList(int(iplanId),3,i+1)
	}
	for i:=0;i<7;i++{
		sit.SitFour[i].Row=4
		sit.SitFour[i].Line=i+1
		sit.SitFour[i].Flag=GetTicketByRowAndList(int(iplanId),4,i+1)
	}
	for i:=0;i<7;i++{
		sit.SitFive[i].Row=5
		sit.SitFive[i].Line=i+1
		sit.SitFive[i].Flag=GetTicketByRowAndList(int(iplanId),5,i+1)
	}
	for i:=0;i<7;i++{
		sit.SitSix[i].Row=6
		sit.SitSix[i].Line=i+1
		sit.SitSix[i].Flag=GetTicketByRowAndList(int(iplanId),6,i+1)
	}
	for i:=0;i<7;i++{
		sit.SitSeven[i].Row=7
		sit.SitSeven[i].Line=i+1
		sit.SitSeven[i].Flag=GetTicketByRowAndList(int(iplanId),7,i+1)
	}
	return sit
}
func GetTicketByRowAndList(PlanId int,Row int,list int)int  {
	sqlStr := "select ticket_id,movie_name,cinema_id,plan_id,row,list,state from ticket where plan_id=? and row = ? and list=?"
	row:=utils.Db.QueryRow(sqlStr,PlanId,Row,list)
	ticket:=&model.Ticket{}
	err:=row.Scan(&ticket.TicketId,&ticket.MovieName,&ticket.CinemaId,&ticket.PlanId,&ticket.Row,&ticket.List,&ticket.State)
	if err!=nil {
		fmt.Println(err)
	}
	fmt.Println(99999,ticket)
	return ticket.State
}
func GetTicketByRowAndList2(PlanId int,Row int,list int) *model.Ticket  {
	sqlStr := "select ticket_id,movie_name,cinema_id,plan_id,row,list,state from ticket where plan_id=? and row = ? and list=?"
	row:=utils.Db.QueryRow(sqlStr,PlanId,Row,list)
	ticket:=&model.Ticket{}
	row.Scan(&ticket.TicketId,&ticket.MovieName,&ticket.CinemaId,&ticket.PlanId,&ticket.Row,&ticket.List,&ticket.State)
	fmt.Println(8888,ticket)
	return ticket
}
func UpdateTicket(ticket *model.Ticket) error {
	sqlStr:="update ticket set state=? where ticket_id=?"
	fmt.Println("qqqqq",ticket)
	_, err := utils.Db.Exec(sqlStr,ticket.State,ticket.TicketId)

	if err != nil {
		fmt.Println(err)
	}
	return nil
}
