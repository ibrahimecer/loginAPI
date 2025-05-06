package config

import (
	"fmt"
	"io/ioutil"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "postgres"
	dbname := "login_api"

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port,
	)

	// GORM bağlantısı
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Veritabanına bağlanırken hata oluştu:", err)
	}
	DB = db

	// sql.DB kullanarak migration dosyasını çalıştır
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("sql.DB bağlantısı alınamadı:", err)
	}

	initSQL, err := ioutil.ReadFile("sql/create_table.sql")
	if err != nil {
		log.Fatal("SQL dosyası okunamadı:", err)
	}

	_, err = sqlDB.Exec(string(initSQL))
	if err != nil {
		log.Fatal("Migration SQL çalıştırılamadı:", err)
	}

	fmt.Println("✅ Veritabanı bağlantısı başarılı, migration dosyası çalıştırıldı.")
}
