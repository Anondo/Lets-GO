package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Student struct {
	uid      int
	username string
	dept     string
	cgpa     float64
}

/*

CREATE TABLE `Student` (
	    `uid` INT(10) NOT NULL AUTO_INCREMENT,
	    `username` VARCHAR(64) NULL DEFAULT NULL,
	    `dept` VARCHAR(64) NULL DEFAULT NULL,
	    `cgpa` Double NULL DEFAULT NULL,
	    PRIMARY KEY (`uid`)
	);

*/

func main() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golang")
	checkErr(err)

	// query
	fmt.Println(query(db, "SELECT * FROM Student"))
	//Insert
	q := "INSERT Student SET name=?,dept=?,cgpa=?"
	fmt.Println(dmlQuery(db, q, Student{username: "Damien", dept: "Not Sure", cgpa: 4.00}))

	defer db.Close()

}

func query(db *sql.DB, q string) []Student {
	rows, err := db.Query(q)
	var students []Student
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var cgpa float64
		err = rows.Scan(&uid, &username, &department, &cgpa)
		checkErr(err)
		students = append(students, Student{uid, username, department, cgpa})
	}
	rows.Close()
	return students
}
func dmlQuery(db *sql.DB, q string, s Student) string {
	stmnt, err := db.Prepare(q)
	checkErr(err)
	res, err := stmnt.Exec(s.username, s.dept, s.cgpa)
	checkErr(err)
	fmt.Sprintf("%v", res)
	return "Successful!"
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
