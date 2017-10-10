package orm

import (
	"strconv"
)

func(m *ormsql)mysqllimit(page int,limit int)*ormsql{
	begin := strconv.Itoa(page)
	num := strconv.Itoa(limit)
	m.pagination = " limit "+begin+","+num
	return m
}



