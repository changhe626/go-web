package main

import (
	_ "github.com/Go-SQL-Driver/MySQL"
	"database/sql"
	"fmt"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
		/**
		panic一般会导致程序挂掉（除非recover）  然后Go运行时会打印出调用栈
        但是，关键的一点是，即使函数执行的时候panic了，函数不往下走了，运行时并不是立刻向上传递panic，而是到defer那，
		等defer的东西都跑完了，panic再向上传递。所以这时候 defer 有点类似 try-catch-finally 中的 finally。panic就是这么简单。抛出个真正意义上的异常。
		 */
	}
}

/**
这玩意是写一行代码就是一个异常错误的检查,简直是脑壳进水了的设计..........
除了插入检查了异常,其他的都没有检查异常....好麻烦,以后自己用的时候需要自己再好好封装一遍
 */
func main() {

	// 插入数据
	db,err:=sql.Open("mysql","root:123@/vspmanager?charset=utf8")
	checkErr(err)
	stmt, err :=db.Prepare("INSERT INTO `diag_info` (`diag_id`, `diag_name`, `diag_code`, `diag_alias_name`) VALUES (?, ?, ?, ?)")
	checkErr(err)
	res,err:=stmt.Exec(2,"zhaojun","110","无名")
	checkErr(err)
	id, err := res.LastInsertId()
	fmt.Println(id)



	// 更新数据
	stmt, err = db.Prepare("update diag_info set diag_name=? where diag_id=?")
	res, err = stmt.Exec("zhaojun222",1)
	affect, err :=res.RowsAffected()
	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("SELECT * FROM diag_info limit 10")
	for rows.Next() {
		var id int
		var diag_name string
		var diag_code string
		var diag_alias_name string
		err = rows.Scan(&id, &diag_name, &diag_code, &diag_alias_name)
		fmt.Println(id)
		fmt.Println(diag_name)
		fmt.Println(diag_code)
		fmt.Println(diag_alias_name)
	}

	// 删除数据
	stmt, err = db.Prepare("delete from diag_info where diag_id=?")
	res, err = stmt.Exec(1)
	affect, err = res.RowsAffected()
	fmt.Println(affect)
	db.Close()


	
}
