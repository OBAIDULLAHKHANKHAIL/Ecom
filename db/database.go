package db

import(
	"log"
	"os"
	"obaid/models"
	
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(sqlite.Open("api.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to Connect to the database! \n", err.Error())
		os.Exit(2)
	}
	log.Println("Connected to the database Successfully")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migration")

	db.AutoMigrate(&models.User{}, &models.Product{},)

	Database = DbInstance{Db: db}
}