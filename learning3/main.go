package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tobgu/qframe"
	qsql "github.com/tobgu/qframe/config/sql"
	groupby "github.com/tobgu/qframe/config/groupby"

)

import (
	"log"
)

func main() {
	// Create a new in-memory SQLite database.
	db, _ := sql.Open("sqlite3", ":memory:")
	// Add a new table.
	db.Exec(`
	CREATE TABLE test (
		COL1 INT,
		COL2 REAL,
		COL3 TEXT,
		COL4 BOOL
	);`)
	// Create a new QFrame to populate our table with.
	qf := qframe.New(map[string]interface{}{
		"COL1": []int{1, 2, 3},
		"COL2": []float64{1.1, 2.2, 3.3},
		"COL3": []string{"one", "two", "three"},
		"COL4": []bool{true, true, true},
	})
	fmt.Println("qf:\n",qf)
	
	// Start a new SQL Transaction.
	tx, _ := db.Begin()
	
	// Write the QFrame to the database.
	qf.ToSQL(tx,
		// Write only to the test table
		qsql.Table("test"),
		// Explicitly set SQLite compatibility.
		qsql.SQLite(),
	)

	// Create a new QFrame from SQL.
	newQf := qframe.ReadSQL(tx,
		// A query must return at least one column. In this 
		// case it will return all of the columns we created above.
		qsql.Query("SELECT * FROM test"),
		// SQLite stores boolean values as integers, so we
		// can coerce them back to bools with the CoercePair option.
		qsql.Coerce(qsql.CoercePair{Column: "COL4", Type: qsql.Int64ToBool}),
		qsql.SQLite(),
	)

	fmt.Println("newQf:\n",newQf)
	fmt.Println(newQf.Equals(qf))
	log.Println("main() of learning3/main.go")

	f := qframe.New(map[string]interface{}{"COL1": []int{1, 2, 3},
		"COL2": []string{"a", "b", "c"}})
	newF := f.Filter(qframe.Or(
    	qframe.Filter{Column: "COL1", Comparator: ">", Arg: 2},
    	qframe.Filter{Column: "COL2", Comparator: "=", Arg: "a"}))
	fmt.Println("newF\n",newF)

	intSum := func(xx []int) int {
		result := 0
		for _, x := range xx {
			result += x
		}
		return result
	}
	
	f1 := qframe.New(map[string]interface{}{"COL1": []int{1, 2, 2, 3, 3}, "COL2": []string{"a", "b", "c", "a", "b"}})
	fmt.Println("f1 before grouping:\n",f1.Sort(qframe.Order{Column: "COL2"}))
	f1 = f1.GroupBy(groupby.Columns("COL2")).Aggregate(qframe.Aggregation{Fn: intSum, Column: "COL1"})
	fmt.Println("f1:\n",f1.Sort(qframe.Order{Column: "COL2"}))

}

// func main() {
// }
