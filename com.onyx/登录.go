package main

import (
	"net/http"
	"fmt"
	"log"
)

/**
获取参数,进行登录的判断
 */
func login(w http.ResponseWriter,r *http.Request){
	fmt.Println(r.Method)
	if r.Method=="POST"{
		/**
		默认情况下，Handler里面是不会自动解析form的，
		必须显式的调用 r.ParseForm() 后，你才能对这个表单数据进行操作
		 */
		r.ParseForm()
		if len(r.Form["name"][0])==0{
			fmt.Println("用户名不能为空")
			return
		}
		fmt.Println("username,",r.Form["name"])
		fmt.Println("password,",r.Form["password"])
	}
}


func main() {
	http.HandleFunc("/login",login)
	err :=http.ListenAndServe("127.0.0.1:9000",nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

/**
<form action="http://127.0.0.1:9000/login" method="POST">
    账号:<input name="name">
    密码:<input name="password">
    <input type="submit" value="提交" />
</form>
 */
