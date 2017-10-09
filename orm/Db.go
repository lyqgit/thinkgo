package orm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "code.google.com/p/odbc"
	"database/sql"
	"strings"
)

type ormsql struct{
	database string
	sqltype string
	db *sql.DB
	row string
	conditionQuery string
	conditionVal string
	conditionRead string
	conditionOrder string
	conditionWhere string
	WhereVal string
	out []map[string]string
	sqlString string
}

func(m *ormsql)table(database string){
	m.database = database
	m.row = ""
	m.conditionOrder = ""
	m.conditionQuery = ""
	m.conditionRead = ""
	m.conditionVal = ""
	m.conditionWhere = ""
	m.WhereVal = ""
	m.out = nil
}

func(m *ormsql)where(condition ...interface{}){
	switch t := condition[0].(type){
		case map[string]string:
			for k,v := range t{
				m.conditionWhere += k+"=?,"
				m.WhereVal += v+","
			}
		break
		case string:
			if m.conditionRead == ""{
				m.conditionWhere = condition[0].(string)+"=?,"
			}else{
				m.WhereVal = condition[1].(string)+","
			}
		break
	}
}

func(m *ormsql)order(row string,sort string){
	m.conditionOrder = "order by "+row+" "+sort 
}

func(m *ormsql)insert(add map[string]string)int64{
	for k,v := range add{
		m.conditionRead += k+"=?,"
		m.conditionVal += v+","
	}
	m.conditionRead = strings.TrimRight(m.conditionRead,",")
	m.conditionVal = strings.TrimRight(m.conditionVal,",")
	stmt,err := m.db.Prepare("insert "+m.database+" set "+m.conditionRead)
	sqlerr(err)
	res,err := stmt.Exec(m.conditionVal)
	sqlerr(err)
	num,err := res.RowsAffected()
	sqlerr(err)
	return num
}

func(m *ormsql)update(renew map[string]string)int64{
	for k,v := range renew{
		m.conditionRead += k+"=?,"
		m.conditionVal += v+","
	}
	m.conditionRead = strings.TrimRight(m.conditionRead,",")
	m.conditionVal = strings.TrimRight(m.conditionVal,",")
	stmt,err := m.db.Prepare("update "+m.database+" set "+m.conditionRead+" where "+m.WhereVal)
	sqlerr(err)
	res,err := stmt.Exec(m.conditionVal+m.WhereVal)
	sqlerr(err)
	num,err := res.RowsAffected()
	sqlerr(err)
	return num
}

func(m *ormsql)insertGetId(add map[string]string)int64{
	for k,v := range add{
		m.conditionRead += k+"=?,"
		m.conditionVal += v+","
	}
	m.conditionRead = strings.TrimRight(m.conditionRead,",")
	m.conditionVal = strings.TrimRight(m.conditionVal,",")
	stmt,err := m.db.Prepare("insert "+m.database+" set "+m.conditionRead)
	sqlerr(err)
	res,err := stmt.Exec(m.conditionVal)
	sqlerr(err)
	Id,err := res.LastInsertId()
	sqlerr(err)
	return Id
}

func(m *ormsql)insertAll(addAll []map[string]string)int{
	var num []int64
	for _,add := range addAll{
		for k,v := range add{
			m.conditionRead += k+"=?,"
			m.conditionVal += v+","
		}
		m.conditionRead = strings.TrimRight(m.conditionRead,",")
		m.conditionVal = strings.TrimRight(m.conditionVal,",")
		stmt,err := m.db.Prepare("insert "+m.database+" set "+m.conditionRead)
		sqlerr(err)
		res,err := stmt.Exec(m.conditionVal)
		sqlerr(err)
		Id,err := res.LastInsertId()
		sqlerr(err)
		num = append(num,Id)
	}
	return len(num)
}



func sqlerr(err error){//输出错误
	if err != nil{
		fmt.Print(err)
	}
}

func Connect(driverName string,dataSourceName string)*ormsql{//连接数据库
	var m = new(ormsql)
	var err error
	m.db,err = sql.Open(driverName,dataSourceName)
	m.sqltype = driverName
	defer sqlerr(err)
	return m
}
