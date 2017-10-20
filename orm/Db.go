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
	term string
	pagination string
	unionQuery string
	out []map[string]string
	sqlString string
}

func(m *ormsql)table(database string){
	m.database = database
	m.init()
}

func(m *ormsql)init(){
	m.row = "*"
	m.conditionOrder = ""
	m.conditionQuery = ""
	m.conditionRead = ""
	m.conditionVal = ""
	m.conditionWhere = ""
	m.WhereVal = ""
	m.term = ""
	m.pagination = ""//分页
	m.unionQuery = ""//联合查询
	m.out = nil
}

func(m *ormsql)where(condition ...interface{}){
	switch t := condition[0].(type){
		case map[string]string:
			for k,v := range t{
				m.conditionWhere += k+"=?,"
				m.WhereVal += v+","
				if sign := condition[1].(string);sign != ""{
					switch sign{
						case "or":
							m.conditionQuery += k+" = "+v+" or "
							m.term = " or "
						break
						default:
							m.conditionQuery += k+" = "+v+" and "
							m.term = " and "
						break
					}
				}
				
			}
		break
		case string:
			if condition[2].(string) != ""{
				m.conditionQuery += t+" = "+condition[1].(string)+" "+condition[2].(string)+" "
				m.term = " "+condition[2].(string)+" "
			}else{
				if m.conditionRead == ""{
					m.conditionWhere = condition[0].(string)+"=?,"
				}else{
					m.WhereVal = condition[1].(string)+","
				}
				m.conditionQuery = t+" = "+condition[1].(string)+" and "
				m.term = " and "
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
	stmt,err := m.db.Prepare("update "+m.database+" set "+m.conditionRead+" where "+m.conditionWhere)
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
		Id := m.insertGetId(add)
		num = append(num,Id)
	}
	return len(num)
}

func(m *ormsql)field(row string){
	m.row = row
}

func(m *ormsql)alias(byname string){
	m.database += " as "+byname
}

func(m *ormsql)join(database string,condition string,symbol string){
	m.unionQuery = symbol+" join "+database+" on "+condition
}

func(m *ormsql)limit(page int,limit int)*ormsql{
	switch m.sqltype{
		case "mysql":
			m.mysqllimit(page,limit)
		break
		case "sqlserver":
			m.sqlserverlimit(page,limit)
		break
	}
	return m
}

func(m *ormsql)read(){

	switch m.sqltype{//组装sql语句，根据不同数据库类型调用各自的查询方法
		case "mysql":
			
		break
		case "sqlserver":
			
		break
	}
	
}

func(m *ormsql)query(sqlString string){
	
}

func sqlerr(err error){//输出错误
	if err != nil{
		fmt.Print(err)
	}
}

func(m *ormsql)destruct(){
	m.db.Close()
}

func Connect(driverName string,dataSourceName string)*ormsql{//连接数据库
	var m = new(ormsql)
	var err error
	m.db,err = sql.Open(driverName,dataSourceName)
	m.sqltype = driverName
	defer sqlerr(err)
	return m
}
