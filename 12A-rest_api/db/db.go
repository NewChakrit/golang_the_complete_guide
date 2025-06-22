package db

import (
	"database/sql"
	_ "github.com/gin-gonic/gin"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("Could not connected database.")
	}

	DB.SetMaxOpenConns(10) // กำหนดค่่าการ pulling เชื่อมต่อ
	DB.SetMaxIdleConns(5)  // กำหนดค่าการเชื่อมต่อสูงสุดต่อการใช้งาน

}
