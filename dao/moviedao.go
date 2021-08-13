package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"fmt"
	"strconv"
)

//获取所有影院信息
func GetCinema()([]*model.Movie,error) {
	sqlStr := "select movie_id,movie_name,duration,movie_number,movie_information,img_address from cinema"
	rows, err := utils.Db.Query(sqlStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var cinemas []*model.Movie
	for rows.Next() {
		cinema := &model.Movie{}
		rows.Scan(&cinema.MovieID, &cinema.MovieName, &cinema.Duration, &cinema.MovieNumber, &cinema.MovieInformation, &cinema.ImgAddress)
		cinemas = append(cinemas, cinema)
	}
	return cinemas,nil
}
func GetMovieById(MovieID string) (*model.Movie, error)  {
	sqlStr := "select movie_id,movie_name,duration,movie_number,movie_information,img_address from cinema where movie_id = ?"
	fmt.Println(MovieID)
	row := utils.Db.QueryRow(sqlStr, MovieID)
	movie := &model.Movie{}
	row.Scan(&movie.MovieID, &movie.MovieName, &movie.Duration, &movie.MovieNumber, &movie.MovieInformation,&movie.ImgAddress)
	fmt.Println(movie)
	return movie, nil
}
func UpdateMovie(b *model.Movie) error {
	sqlStr:="update cinema set movie_name=?,duration=?,movie_number=?,movie_information=?,img_address=? where movie_id=?"
	_, err := utils.Db.Exec(sqlStr, b.MovieName, b.Duration, b.MovieNumber, b.MovieInformation, b.ImgAddress, b.MovieID)
	fmt.Println(b)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func AddMovie(b *model.Movie) error {
	slqStr := "insert into cinema(movie_name,duration,movie_number,movie_information,img_address) values(?,?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(slqStr, b.MovieName, b.Duration, b.MovieNumber, b.MovieInformation, b.ImgAddress)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
func DeleteMovie(movieID string) error {
	sqlStr :="delete from cinema where movie_id = ?"
	_,err:=utils.Db.Exec(sqlStr,movieID)
	if err!=nil {
		fmt.Println(err)
	}
	return nil
}
func GetPageMovies(pageNo string) (*model.Page, error) {
	//将页码转换为int64类型
	iPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	//获取数据库中图书的总记录数
	sqlStr := "select count(*) from cinema"
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

	sqlStr2 := "select movie_id,movie_name,duration,movie_number,movie_information,img_address from cinema limit ?,?"
	//执行
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, err
	}
	var movies []*model.Movie
	for rows.Next() {
		movie := &model.Movie{}
		rows.Scan(&movie.MovieID,&movie.MovieName,&movie.Duration,&movie.MovieNumber,&movie.MovieInformation,&movie.ImgAddress)
		movies=append(movies,movie)
	}
	//创建page
	page := &model.Page{
		Movie:       movies,
		PageNo:      iPageNo,
		PageSize:    pageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}
func GetMovieByName(MovieName string) (*model.Movie, error)  {
	sqlStr := "select movie_id,movie_name,duration,movie_number,movie_information,img_address from cinema where movie_name = ?"
	fmt.Println(MovieName)
	row := utils.Db.QueryRow(sqlStr, MovieName)
	movie := &model.Movie{}
	row.Scan(&movie.MovieID, &movie.MovieName, &movie.Duration, &movie.MovieNumber, &movie.MovieInformation,&movie.ImgAddress)
	fmt.Println(movie)
	return movie, nil
}