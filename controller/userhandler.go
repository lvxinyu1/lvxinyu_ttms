package controller

import (
	"TTMS/dao"
	"TTMS/model"
	"TTMS/utils"
	"html/template"
	"net/http"
)
func Homepage(w http.ResponseWriter,r *http.Request)  {
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
	t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\index.html"))
	t.Execute(w,page)
}
func Login(w http.ResponseWriter,r *http.Request) {
	//判断是否已经登录
	flag, _ := dao.IsLogin(r)
	if flag {
		//已经登录
		//去首页
		Homepage(w, r)
	} else {
		//获取用户名和密码
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		//调用userdao中验证用户名和密码的方法
		user, _ := dao.CheckUserNameAndPassword(username, password)
		if user.ID > 0 {
			//用户名和密码正确
			//生成UUID作为Session的id
			uuid := utils.CreateUUID()
			//创建一个Session
			sess := &model.Session{
				SessionID: uuid,
				UserName:  user.Username,
				UserID:    user.ID,
			}
			//将Session保存到数据库中
			dao.AddSession(sess)
			//创建一个Cookie，让它与Session相关联
			cookie := http.Cookie{
				Name:     "user",
				Value:    uuid,
				HttpOnly: true,
			}
			//将cookie发送给浏览器
			http.SetCookie(w, &cookie)
			t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\users\\login_success.html"))
			t.Execute(w, user)
		} else {
			//用户名或密码不正确
			t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\users\\login.html"))
			t.Execute(w, "用户名或密码不正确！")
		}
	}
}
func Login2(w http.ResponseWriter,r *http.Request)  {
	t:=template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\users\\login.html"))
	t.Execute(w,"")
}
func Regist(w http.ResponseWriter,r *http.Request)  {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	email := r.PostFormValue("email")
	number := r.PostFormValue("number")
	//调用userdao中验证用户名和密码的方法
	user, _ := dao.CheckUserName(username)
	if user.ID > 0 {
		//用户名已存在
		t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\users\\regist.html"))
		t.Execute(w, "用户名已存在！")
	} else {
		//用户名可用，将用户信息保存到数据库中
		dao.SaveUser(username, password, email,number)
		//用户名和密码正确
		t := template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\users\\regist_success.html"))
		t.Execute(w, "")
	}
}
func Regist2(w http.ResponseWriter,r *http.Request)  {
	t:=template.Must(template.ParseFiles("E:\\WorkSpace\\awesomeProject5\\src\\TTMS\\views\\pages\\users\\regist.html"))
	t.Execute(w,"")
}
func Logout(w http.ResponseWriter, r *http.Request) {
	//获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取cookie的value值
		cookieValue := cookie.Value
		//删除数据库中与之对应的Session
		dao.DeleteSession(cookieValue)
		//设置cookie失效
		cookie.MaxAge = -1
		//将修改之后的cookie发送给浏览器
		http.SetCookie(w, cookie)
	}
	//去首页
	Homepage(w, r)
}