package benchs

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/rocketlaunchr/dbq/v2"
)

func (m *Model) ScanFast() []interface{} {
	return []interface{}{&m.Id, &m.Name, &m.Title, &m.Fax, &m.Web, &m.Age, &m.Right, &m.Counter}
}

var dbqdb *sql.DB

func init() {
	st := NewSuite("dbq")
	st.InitF = func() {
		st.AddBenchmark("Insert", 200*OrmMulti, DbqInsert)
		//st.AddBenchmark("MultiInsert 100 row", 200*ORM_MULTI, DbqInsertMulti)
		//st.AddBenchmark("Update", 200*ORM_MULTI, DbqUpdate)
		//st.AddBenchmark("Read", 200*ORM_MULTI, DbqRead)
		//st.AddBenchmark("MultiRead limit 100", 200*ORM_MULTI, DbqReadSlice)

		var err error
		dbqdb, err = sql.Open("pgx", OrmSource)
		checkErr(err)
		err = dbqdb.Ping()
		checkErr(err)
	}
}

func DbqInsert(b *B) {
	var m *Model
	wrapExecute(b, func() {
		initDB()
		m = NewModel()
	})

	for i := 0; i < b.N; i++ {
		m.Id = 0

		stmt := dbq.INSERTStmt("models", []string{"name", "title", "fax", "web", "age", "\"right\"", "counter"}, 1, dbq.PostgreSQL)
		//println(stmt)
		//stmt := "INSERT INTO models ( name,title,fax,web,age,\"right\",counter ) VALUES ($1,$2,$3,$4,$5,$6,$7)"

		_, err := dbq.E(context.Background(), dbqdb, stmt, nil, dbq.Struct(m)[1:])
		//_, err := dbq.E(context.Background(), dbqdb, stmt, nil, "Orm Benchmark", "Just a Benchmark for fun", "99909990", "http://blog.milkpod29.me", 100, true, 1000)

		if err != nil {
			fmt.Println(err)
			b.FailNow()
		}
	}
}

//func DbqInsertMulti(b *B) {
//	var ms []*Model
//	wrapExecute(b, func() {
//		initDB()
//		ms = make([]*Model, 0, 100)
//		for i := 0; i < 100; i++ {
//			ms = append(ms, NewModel())
//		}
//	})
//
//	for i := 0; i < b.N; i++ {
//		for _, m := range ms {
//			m.Id = 0
//		}
//		if err := pgdb.Insert(&ms); err != nil {
//			fmt.Println(err)
//			b.FailNow()
//		}
//	}
//}

//func DbqUpdate(b *B) {
//	var m *Model
//	wrapExecute(b, func() {
//		initDB()
//		m = NewModel()
//		stmt := dbq.INSERTStmt("models", []string{"name", "title", "fax", "web", "age", "\"right\"", "counter"}, 1, dbq.PostgreSQL)
//		_, err := dbq.E(context.Background(), dbqdb, stmt, nil, dbq.Struct(m)[1:])
//		if err != nil {
//			fmt.Println(err)
//			b.FailNow()
//		}
//	})
//
//	for i := 0; i < b.N; i++ {
//
//		//... how to to update?
//
//		if err != nil {
//			fmt.Println(err)
//			b.FailNow()
//		}
//	}
//}

//func DbqRead(b *B) {
//	var m *Model
//	wrapExecute(b, func() {
//		initDB()
//		m = NewModel()
//		if err := pgdb.Insert(m); err != nil {
//			fmt.Println(err)
//			b.FailNow()
//		}
//	})
//
//	for i := 0; i < b.N; i++ {
//		if err := pgdb.Select(m); err != nil {
//			fmt.Println(err)
//			b.FailNow()
//		}
//	}
//}
//
//func DbqReadSlice(b *B) {
//	var m *Model
//	wrapExecute(b, func() {
//		initDB()
//		m = NewModel()
//		for i := 0; i < 100; i++ {
//			m.Id = 0
//			if err := pgdb.Insert(m); err != nil {
//				fmt.Println(err)
//				b.FailNow()
//			}
//		}
//	})
//
//	for i := 0; i < b.N; i++ {
//		var models []*Model
//		if err := pgdb.Model(&models).Where("id > ?", 0).Limit(100).Select(); err != nil {
//			fmt.Println(err)
//			b.FailNow()
//		}
//	}
//}
