package utils

import (
	// "database/sql"
	// "log"
	"strings"
	"time"
)

func ForeverSleep(d time.Duration, f func(int) error) {
	for i := 0; ; i++ {
		err := f(i)
		if err == nil {
			return
		}
		time.Sleep(d)
	}
}

func FormatUUID(s string) string {
	return strings.ReplaceAll(s, "-", "")
}

/*func LoopRow(row *sql.Rows, array []struct{}) []map[string]interface{}{
	var allMaps []map[string]interface{}
	columns, _ := row.Columns()
	for row.Next(){
		values := make([]interface{}, len(columns))
		pointers := make([]interface{}, len(columns))
		for i,_ := range values{
			log.Println(pointers[i])
			pointers[i] = &values[i]
		}
		err := row.Scan(pointers...)
		result:= make(map[string]interface{})
		for i, val := range values {
			log.Printf("Adding key=%s val=%v\n", columns[i], val)
        	result[columns[i]] = val 
		}
		allMaps = append(allMaps, result)
		return allMaps

	}
}*/

