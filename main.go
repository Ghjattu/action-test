package main

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDatabase() {
	godotenv.Load("")

	mysql_username := os.Getenv("MYSQL_USERNAME")
	mysql_password := os.Getenv("MYSQL_PASSWORD")
	mysql_ip := os.Getenv("MYSQL_IP")
	mysql_port := os.Getenv("MYSQL_PORT")
	mysql_db_name := os.Getenv("MYSQL_DB_NAME")

	// log.Println("MYSQL_USERNAME: " + mysql_username)
	// log.Println("MYSQL_PASSWORD: " + mysql_password)
	// log.Println("MYSQL_IP: " + mysql_ip)
	// log.Println("MYSQL_PORT: " + mysql_port)
	// log.Println("MYSQL_DB_NAME: " + mysql_db_name)

	dsn := mysql_username + ":" + mysql_password +
		"@tcp(" + mysql_ip + ":" + mysql_port + ")" + "/" + mysql_db_name +
		"?charset=utf8mb4&parseTime=True&loc=Local"

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func main() {

}
