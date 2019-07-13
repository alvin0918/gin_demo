package mysql

import (
	"database/sql"
	"gopkg.in/ini.v1"
	"github.com/alvin0918/gin_demo/core/config"
	"github.com/alvin0918/gin_demo/core/commin/log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

//链接数据库
func connect(username string, password string, ip string, port string, dbname string) (db *sql.DB, err error) {

	var (
		info string
	)

	info = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", username, password, ip, port, dbname)
	fmt.Println(info)

	if db, err = sql.Open("mysql", info); err != nil {
		log.ErrorPrintf(err.Error())
		panic(err)
	}

	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)

	if err := db.Ping(); err != nil{
		log.ErrorPrintf(err.Error())
		panic(err.Error())
	}

	log.ErrorPrintf("ACB")

	return db, err
}

// 分布式链接数据库
func getConnect(section string) (*sql.DB) {

	var (
		username string
		password string
		ip string
		port string
		dbanme string
		sections *ini.Section
		err error
		db *sql.DB
	)

	if sections, err = config.GetSection(section); err != nil {
		panic(err)
	}

	username 	= sections.Key("UserName").String()
	password 	= sections.Key("PassWord").String()
	ip 			= sections.Key("IP").String()
	port 		= sections.Key("Port").String()
	dbanme 		= sections.Key("DBName").String()

	if db, err = connect(username, password, ip, port, dbanme); err != nil {
		panic(err)
	}

	return db

}

func Create(sqls string, section string)  {
	
}

func Update(sqls string, section string)  {

}

func Detele(sqls string, section string)  {

}

func Select(sqls string, section string) map[string]string {

	var (
		db *sql.DB
		err error
		rows *sql.Rows
		str []string
		scanArgs []interface{}
		values []interface{}
		record map[string]string
	)
	db = getConnect(section)
	defer db.Close()

	if rows, err = db.Query(sqls); err != nil {
		log.ErrorPrintf("sql '" + sqls + "' exec filed")
	}

	if str, err = rows.Columns(); err != nil {
		log.ErrorPrintf(err.Error())
	}

	scanArgs = make([]interface{}, len(str))
	values = make([]interface{}, len(str))

	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		record = make(map[string]string)
		for i, col := range values {
			if col != nil {
				record[str[i]] = string(col.([]byte))
			}
		}
	}

	return record
}

func Insert(sqls string, section string) (int64, error) {
	var (
		db *sql.DB
		err error
		res sql.Result
	)

	db = getConnect(section)

	defer db.Close()

	if res, err = db.Exec(sqls); err != nil {
		log.TracePrintf("sql '" + sqls + "' exec filed")
	}

	return res.LastInsertId()
}



















