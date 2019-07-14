package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"gopkg.in/ini.v1"
	"github.com/alvin0918/gin_demo/core/config"
	"github.com/alvin0918/gin_demo/core/commin/log"
)

type MysqlInfomation struct {
	connect map[string]*sql.DB
}


var (
	Section string
	MysqlContent *MysqlInfomation
)

//链接数据库
func init() {

	go func() {
		var (
			username string
			password string
			ip string
			port string
			dbname string
			charset string
			sections *ini.Section
			db *sql.DB
			err error
			info string
			num int
			connect map[string]*sql.DB
			str string
		)

		if sections, err = config.GetSection("Mysql"); err != nil {
			log.TracePrintf("Mysql", err.Error())
		}

		if num, err = sections.Key("web_num").Int(); err != nil {
			log.TracePrintf("Mysql", err.Error())
		}

		connect = make(map[string]*sql.DB)

		for i := 0; i < num; i++ {

			str = fmt.Sprintf("Mysql_web%d", i)

			if sections, err = config.GetSection(str); err != nil {
				log.TracePrintf("Mysql", err.Error())
			}

			username 	= sections.Key("UserName").String()
			password 	= sections.Key("PassWord").String()
			ip 			= sections.Key("IP").String()
			port 		= sections.Key("Port").String()
			dbname 		= sections.Key("DBName").String()
			charset 	= sections.Key("Charset").String()

			info = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
				username, password, ip, port, dbname, charset)

			if db, err = sql.Open("mysql", info); err != nil {
				log.TracePrintf("Mysql", err.Error())
			}

			connect[str] = db

		}

		MysqlContent = &MysqlInfomation{
			connect: connect,
		}

	}()
}

func (db *MysqlInfomation) Create(sqls string, section string) (int64, error) {

	defer db.connect[section].Close()

	var (
		result sql.Result
		err error
		rows int64
	)

	if result, err = db.connect[section].Exec(sqls); err != nil {
		log.TracePrintf("Mysql", err.Error())
	}

	if rows, err = result.RowsAffected(); err != nil {
		log.TracePrintf("Mysql", err.Error())
	}

	return rows, err

}

func (db *MysqlInfomation) Update(sqls string, section string) (int64, error) {
	defer db.connect[section].Close()

	var (
		result sql.Result
		stmt *sql.Stmt
		err error
		rows int64
	)

	if stmt, err = db.connect[section].Prepare(sqls); err != nil {
		log.TracePrintf("Mysql", err.Error())
	}

	if result, err = stmt.Exec(); err != nil {
		log.TracePrintf("Mysql", err.Error())
	}

	if rows, err = result.RowsAffected(); err != nil {
		log.TracePrintf("Mysql", err.Error())
	}

	return rows, err

}

func (db *MysqlInfomation) Detele(sqls string, section string) (int64, error) {
	defer db.connect[section].Close()

	var (
		result sql.Result
		stmt *sql.Stmt
		err error
		rows int64
	)

	if stmt, err = db.connect[section].Prepare(sqls); err != nil {
		log.TracePrintf("Mysql", err.Error())
	}

	if result, err = stmt.Exec(); err != nil {
		log.TracePrintf("Mysql", err.Error())
	}

	if rows, err = result.RowsAffected(); err != nil {
		log.TracePrintf("Mysql", err.Error())
	}

	return rows, err

}

func (db *MysqlInfomation) Select(sqls string, section string) (map[int]map[string]string, error) {
	defer db.connect[section].Close()

	var (
		query *sql.Rows
		err error
		cols []string
		values [][]byte
		scans []interface{}
		results map[int]map[string]string
		row map[string]string
		i int
	)

	if query, err = db.connect[section].Query(sqls); err != nil {
		log.TracePrintf("Mysql", err.Error())
	}

	//读出查询出的列字段名
	if cols, err = query.Columns(); err != nil {
		log.TracePrintf("Mysql", err.Error())
	}

	//values是每个列的值，这里获取到byte里
	values = make([][]byte, len(cols))

	//query.Scan的参数，因为每次查询出来的列是不定长的，用len(cols)定住当次查询的长度
	scans = make([]interface{}, len(cols))

	//让每一行数据都填充到[][]byte里面
	for i := range values {
		scans[i] = &values[i]
	}

	//最后得到的map
	results = make(map[int]map[string]string)

	i = 0

	//循环，让游标往下推
	for query.Next() {

		//query.Scan查询出来的不定长值放到scans[i] = &values[i],也就是每行都放在values里
		if query.Scan(scans...); err != nil {
			log.TracePrintf("Mysql", err.Error())
		}

		//每行数据
		row = make(map[string]string)

		//每行数据是放在values里面，现在把它挪到row里
		for k, v := range values {
			key := cols[k]
			row[key] = string(v)
		}

		//装入结果集中
		results[i] = row

		i++
	}

	return results, err

}



















