package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"fmt"

	"strconv"
)

func GetPageCinema(pageNo string) (*model.Page,error)  {
	//将页码转换为int64类型
	fmt.Println(pageNo)
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取数据库中图书的总记录数
	sqlStr := "select count(*) from cinemat"
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

	sqlStr2 := "select cinema_id,cinema_name,cinema_seat_row,cinema_seat_list from cinemat limit ?,?"
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var cinemas []*model.Cinema
	for rows.Next() {
		cinema := &model.Cinema{}
		rows.Scan(&cinema.CinemaId,&cinema.CinemaName,&cinema.SeatRow,&cinema.SeatList)
		cinemas=append(cinemas,cinema)
	}
	//创建page
	page := &model.Page{
		Cinema:      cinemas,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}
func GetCinemaById(CinemaId string) (*model.Cinema,error) {
	sqlStr := "select cinema_id,cinema_name,cinema_seat_row,cinema_seat_list from cinemat where cinema_id = ?"
	fmt.Println(CinemaId)
	row := utils.Db.QueryRow(sqlStr, CinemaId)
	cinema := &model.Cinema{}
	row.Scan(&cinema.CinemaId,&cinema.CinemaName,&cinema.SeatRow,&cinema.SeatList)
	fmt.Println(cinema)
	return cinema, nil
}
func AddCinema(b *model.Cinema) error {
	slqStr := "insert into cinemat(cinema_name,cinema_seat_row,cinema_seat_list) values(?,?,?)"
	_,err:=utils.Db.Exec(slqStr,b.CinemaName,b.SeatRow,b.SeatList)
	if err!=nil{
		return err
	}
	return nil
}
func UpdateCinema(b *model.Cinema) error {
	sqlStr:="update cinemat set cinema_name=?,cinema_seat_row=?,cinema_seat_list=? where cinema_id=?"
	_, err := utils.Db.Exec(sqlStr,b.CinemaName,b.SeatRow,b.SeatList,b.CinemaId)
	fmt.Println(b)
	if err != nil {
		fmt.Println(err)
	}

	return nil
}
func DeleteCinema(cinemaID string) error {
	sqlStr:="delete from cinemat where cinema_id = ? "
	_,err:=utils.Db.Exec(sqlStr,cinemaID)
	if err!=nil{
		return err
	}
	return nil
}
func GetBadSeatByCinemaID(CinemaId string)(int,[] *model.Badseat){
	sqlStr := "select bad_seat_id,list,row,cinema_id from bad_seat where cinema_id = ?"
	fmt.Println(CinemaId)
	rows,_ := utils.Db.Query(sqlStr, CinemaId)
	var badseats []*model.Badseat
	for rows.Next(){
		badseat:=&model.Badseat{}
		err2:=rows.Scan(&badseat.BadseatId,&badseat.List,&badseat.Row,&badseat.CinemaId)
		if err2!=nil{
			return 0,nil
		}
		badseats=append(badseats,badseat)
	}
	sqlStr2:="select count(*) from bad_seat where cinema_id = ?"
	row:=utils.Db.QueryRow(sqlStr2,CinemaId)
	var count int
	row.Scan(count)
	return count,badseats
}