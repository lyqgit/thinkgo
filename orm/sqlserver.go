package orm

import (
	"strconv"
	"database/sql"
	"time"
)

func(m *ormsql)sqlserverlimit(page int,limit int)*ormsql{
	m.pagination = strconv.Itoa(page)+","+strconv.Itoa(limit)
	return m
}

func(m *ormsql)sqlserverquery(sqlString string)[]map[string]string{
	var rows *sql.Rows
	
	m.sqlString = sqlString
	rows,_ = m.db.Query(m.sqlString)

	columns,_ := rows.Columns()
	scanArgs := make([]interface{},len(columns))
	values := make([]interface{},len(columns))

	for i := range values{
		scanArgs[i] = &values[i]
	}

	for rows.Next(){
		record := make(map[string]string)
		rows.Scan(scanArgs...)


		for i,col := range values{
			// if columns[i] == "Id"{
			// 	v := col.([]byte)
			// 	t := col.([]uint8)
			// 	s := strconv.FormatUint(uint64(t[2]),32)
			// 	fmt.Println(v)
			// 	fmt.Println(string(s))
			// 	e,_ := json.Marshal(col.([]byte))
			// 	fmt.Println(e)
			// 	fmt.Println(string(e))
			// 	fmt.Println(base64.StdEncoding.DecodeString(string(e)))
			// 	// s := string(col.([]byte))
			// 	// enc := mahonia.NewDecoder("utf-8")
			// 	// fmt.Println(enc.ConvertString(s))
			// }
			switch col.(type){
				case string:
					record[columns[i]] = col.(string)
					break
				case int64:
					record[columns[i]] = strconv.FormatInt(col.(int64),10)
					break
				case time.Time:
					record[columns[i]] = col.(time.Time).String()[:19]
					break
				case bool:
					record[columns[i]] = strconv.FormatBool(col.(bool))
					break
				case nil:
					record[columns[i]] = "null"
					break
				default:
					// enc := mahonia.NewDecoder("utf-8")
					record[columns[i]] = string(col.([]byte))
					break

			}
			
		}
		m.out = append(m.out,record)
		
	}
	m.destruct()
	return m.out

}
