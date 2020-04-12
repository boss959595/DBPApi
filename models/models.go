package models

import (
	"database/sql"
	_ "database/sql"
	"fmt"

	"github.com/boss959595/dbp/db"
)

type Employee struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Salary int    `json: "salary"`
	Age    int    `json : "age"`
}
type Employees struct {
	Employees []Employee `json:"employee"`
}

var con *sql.DB

func GetEmployee() Employees {
	con := db.CreateCon()

	rows, err := con.Query("SELECT id,name,age,salary FROM employee order by id")

	fmt.Println("row::", rows)
	fmt.Println("err::", err)
	if err != nil {
		fmt.Println(err)
		//fmt.Fatal("err")
		//return c.JSON(http.StatusCreated, u)
	}
	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}

		err := rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		// Exit if we get an error
		if err != nil {
			fmt.Print(err)
		}
		result.Employees = append(result.Employees, employee)
	}
	return result

}
