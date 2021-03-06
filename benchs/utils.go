package benchs

import (
	"database/sql"
	"fmt"
	"os"
)

type Model struct {
	Id      int `orm:"auto" gorm:"primary_key" db:"id"`
	Name    string
	Title   string
	Fax     string
	Web     string
	Age     int
	Right   bool
	Counter int64
}

func NewModel() *Model {
	m := new(Model)
	m.Name = "Orm Benchmark"
	m.Title = "Just a Benchmark for fun"
	m.Fax = "99909990"
	m.Web = "http://blog.milkpod29.me"
	m.Age = 100
	m.Right = true
	m.Counter = 1000

	return m
}

var (
	OrmMulti   int
	OrmMaxIdle int
	OrmMaxConn int
	OrmSource  string
)

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}
}

func wrapExecute(b *B, cbk func()) {
	b.StopTimer()
	defer b.StartTimer()
	cbk()
	b.ResetTimer()
}

func initDB() {
	sqls := []string{
		`DROP TABLE IF EXISTS models;`,
		`CREATE TABLE models (
			id SERIAL NOT NULL,
			name text NOT NULL,
			title text NOT NULL,
			fax text NOT NULL,
			web text NOT NULL,
			age integer NOT NULL,
			"right" boolean NOT NULL,
			counter bigint NOT NULL,
			CONSTRAINT models_pkey PRIMARY KEY (id)
			) WITH (OIDS=FALSE);`,
	}

	DB, err := sql.Open("postgres", OrmSource)
	checkErr(err)
	defer func() {
		err := DB.Close()
		checkErr(err)
	}()

	err = DB.Ping()
	checkErr(err)

	for _, line := range sqls {
		_, err = DB.Exec(line)
		checkErr(err)
	}
}
