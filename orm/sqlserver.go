package orm

import (
	"strconv"
)

func(m *ormsql)sqlserverlimit(page int,limit int)*ormsql{
	if page == 0{
		m.pagination = " top "+strconv.Itoa(limit)
	}else{
		p := page+limit
		n := limit
		begin := strconv.Itoa(page)
		num := strconv.Itoa(limit)
		m.pagination = " limit "+begin+","+num
	}
	
	return m
}

func(m *ormsql)sqlserver
