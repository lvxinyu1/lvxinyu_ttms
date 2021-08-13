package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"fmt"
	"strconv"
)

func GetPagePlan(pageNo string) (*model.Page,error)  {
	//将页码转换为int64类型
	fmt.Println(pageNo)
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取数据库中图书的总记录数
	sqlStr := "select count(*) from plan"
	//设置一个变量接收总记录数
	var totalRecord int64
	//执行
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	//设置每页只显示4条记录
	var pageSize int64 = 4
	//设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%pageSize == 0 {
		totalPageNo = totalRecord / pageSize
	} else {
		totalPageNo = totalRecord/pageSize + 1
	}

	sqlStr2 := "select plan_id,cinema_id,movie_name,plan_time from plan limit ?,?"
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var plans []*model.Plan
	for rows.Next() {
		plan := &model.Plan{}
		rows.Scan(&plan.PlanId,&plan.CinemaId,&plan.MovieName,&plan.PlanTime)
		plans=append(plans,plan)
	}
	//创建page
	page := &model.Page{
		Plan: 		 plans,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

func UpdatePlan(plan *model.Plan) error {
	sqlStr:="update plan set cinema_id=?,movie_name=?,plan_time=? where plan_id=?"
	_, err := utils.Db.Exec(sqlStr,plan.CinemaId,plan.MovieName,plan.PlanTime, plan.PlanId)
	fmt.Println(plan)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func GetPlanById(PlanId string) (*model.Plan,error) {
	sqlStr := "select plan_id,cinema_id,movie_name,plan_time from plan where plan_id = ?"
	row := utils.Db.QueryRow(sqlStr, PlanId)
	fmt.Println(222,row)
	plan := &model.Plan{}
	row.Scan(&plan.PlanId,&plan.CinemaId,&plan.MovieName,&plan.PlanTime)
	fmt.Println(333,plan)
	return plan, nil
}
func DelatePlanById(PlanId string) error {
	sqlStr :="delete from plan where plan_id = ?"
	_,err:=utils.Db.Exec(sqlStr,PlanId)
	if err!=nil {
		fmt.Println(err)
	}
	return nil
}
func PreserveTicket(tickets [] *model.Ticket) error {
	sqlStr:="insert into ticket(movie_name,cinema_id,plan_id,row,list,state) values(?,?,?,?,?,?)"
	for _,v:=range tickets{
		fmt.Println("票",v.MovieName,v.PlanId)
		_,err:=utils.Db.Exec(sqlStr,v.MovieName,v.CinemaId,v.PlanId,v.Row,v.List,v.State)
		if err!=nil{
			fmt.Println(err)
		}
	}
	return nil
}
func AddPlan(plan *model.Plan) error {
	slqStr := "insert into plan(cinema_id,movie_name,plan_time) values(?,?,?)"
	//执行
	_, err := utils.Db.Exec(slqStr,plan.CinemaId,plan.MovieName,plan.PlanTime)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func GetPlanByName(movieName string) ([]*model.Plan,error) {
	sqlStr := "select plan_id,cinema_id,movie_name,plan_time from plan where movie_name = ?"
	rows,err := utils.Db.Query(sqlStr, movieName)
	if err!=nil{
		fmt.Println(err)
		return nil,err
	}
	var plans []*model.Plan
	for rows.Next() {
		plan:=&model.Plan{}
		rows.Scan(&plan.PlanId,&plan.CinemaId,&plan.MovieName,&plan.PlanTime)
		plans=append(plans,plan)
	}
	return plans, nil
}
func GetPlanById2() (*model.Plan,error) {
	sqlStr1:="select max(plan_id) from plan"
	row1:=utils.Db.QueryRow(sqlStr1)
	var PlanId int
	row1.Scan(&PlanId)
	fmt.Println("定义出来就是这",PlanId)
	sqlStr := "select plan_id,cinema_id,movie_name,plan_time from plan where plan_id = ?"
	row := utils.Db.QueryRow(sqlStr, PlanId)
	fmt.Println(222,row)
	plan := &model.Plan{}
	row.Scan(&plan.PlanId,&plan.CinemaId,&plan.MovieName,&plan.PlanTime)
	fmt.Println(333,plan)
	return plan, nil
}

