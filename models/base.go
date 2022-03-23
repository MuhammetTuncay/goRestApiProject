package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB //database

func init() {

	e := godotenv.Load() //Load .env file
	if e != nil {
		fmt.Print(e)
	}

	//username := os.Getenv("db_user")
	//password := os.Getenv("db_pass")
	//dbName := os.Getenv("db_name")
	//dbHost := os.Getenv("db_host")

	// Connection stringi yaratılır
	//dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)

	// Eğer Heroku üzerinde bir PostgreSQL'e sahipseniz, bu ayarlamaları yapmak yerine doğrudan
	// heroku tarafından verilen database url'i kullanabilirsiniz
	dbUri := fmt.Sprintf("postgres://gequsarxgnbbdp:6c2330a697c14e04e2fe2519ea06b0aefc5794a0703b4fcab7bbd1717be50aca@ec2-52-18-116-67.eu-west-1.compute.amazonaws.com:5432/dfup5tilrb601r") // Database url

	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Account{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
