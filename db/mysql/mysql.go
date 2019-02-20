package gmysql

import (
	"database/sql"
	"log"
	"net"
	"time"

	"github.com/go-sql-driver/mysql"
)

// base on go-sql-driver, https://github.com/go-sql-driver/mysql.git

//Params intferface is logic implement.
type Params interface {
	load() // table load from mysql server.
	save() // table insert/update to mysql server.
}

//DbMysql mysql object
type DbMysql struct {
	cfg *mysql.Config
	db  *sql.DB
}

//NewMysqlClient new client and check
func NewMysqlClient(addr, user, passwd, dbName string) *DbMysql {
	db := &DbMysql{
		cfg: &mysql.Config{
			User:         user,
			Addr:         addr,
			Passwd:       passwd,
			DBName:       dbName,
			Net:          "tcp",
			Timeout:      5 * time.Second,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}

	db.cfg.Params = make(map[string]string)
	db.cfg.Params["charset"] = "utf8"

	//connect check
	c, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("connect failed. %v", err)
		c.Close()
		return nil
	}

	return db
}

//Start mysql connect
func (d *DbMysql) Start() error {
	db, err := sql.Open("mysql", d.cfg.FormatDSN())
	if err != nil {
		log.Printf("mysql start failed. %v", err)
		return err
	}

	d.db = db
	return nil
}
