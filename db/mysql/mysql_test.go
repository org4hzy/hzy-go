package gmysql

import (
	"log"
	"testing"
)

type GoTable struct {
	ID    int32
	Txt   string
	Point int32
}


func TestMySQL(t *testing.T) {
	dbCli := NewMysqlClient("127.0.0.1:3306", "root", "", "test")
	dbCli.Start()

	rows, err := dbCli.db.Query("SELECT * FROM gotable")
	for rows.Next() {
		t := new(GoTable)
		err = rows.Scan(&t.ID, &t.Txt, &t.Point)
		if err != nil {
			continue
		}
		log.Printf("%v", t)
	}

	
}
