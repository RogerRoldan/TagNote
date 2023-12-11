package infraestructure

import (
	"encoding/json"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Credentials struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
	Port     string `json:"port"`
	SSLMode  string `json:"sslmode"`
	TimeZone string `json:"timezone"`
}

// function return connection database
func GetConnection() *gorm.DB {
	// Open JSON file
	file, err := os.Open("C:\\Users\\PC\\Desktop\\Santiago\\GO\\Workhub\\Backend\\credentials.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode JSON data
	var creds Credentials
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&creds)
	if err != nil {
		log.Fatal(err)
	}

	// Create connection string
	stringConnection := "host=" + creds.Host + " user=" + creds.User + " password=" + creds.Password + " dbname=" + creds.DBName + " port=" + creds.Port + " sslmode=" + creds.SSLMode + " TimeZone=" + creds.TimeZone

	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	log.Println("Conectado exitosamente a la DB...")

	return db
}
