package main

import (
	"fmt"
	"log"
	"time"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/logging/log15"
	"github.com/goadesign/goa/middleware"
	"github.com/goadesign/gorma-cellar/app"
	"github.com/goadesign/gorma-cellar/models"
	"github.com/inconshreveable/log15"
	"github.com/jinzhu/gorm"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
)

var db *gorm.DB
var logger log15.Logger
var adb *models.AccountDB
var bdb *models.BottleDB

// settings holds the gorma cellar configuration.
var settings Settings

// Settings is the struct that holds config information retrieved
// from the environment
type Settings struct {
	DatabaseHost     string `envconfig:"gorma_cellar_db_host" default:"localhost"`
	DatabaseUsername string `envconfig:"gorma_cellar_db_username" default:"gorma"`
	DatabasePassword string `envconfig:"gorma_cellar_db_password" default:"gorma"`
	DatabaseName     string `envconfig:"gorma_cellar_db_name" default:"gorma"`
	DatabasePort     int    `envconfig:"gorma_cellar_db_port" default:"5432"`
	MaxOpenConns     int    `envconfig:"gorma_cellar_db_max_open" default:"100"`
	MaxIdleConns     int    `envconfig:"gorma_cellar_db_max_idle" default:"10"`
	Debug            bool   `envconfig:"gorma_cellar_debug" default:"false"`
}

func main() {

	err := envconfig.Process("gorma", &settings)
	if err != nil {
		log.Fatal(err)
	}

	// Create service
	url := fmt.Sprintf("dbname=%s user=%s password=%s sslmode=disable port=%d host=%s",
		settings.DatabaseName,
		settings.DatabaseUsername,
		settings.DatabasePassword,
		settings.DatabasePort,
		settings.DatabaseHost)

	fmt.Println(url)
	var dbStartupWait time.Duration = 7 * time.Second
	time.Sleep(dbStartupWait)

	db, err = gorm.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.LogMode(true)

	db.DropTable(&models.Account{}, &models.Bottle{})
	db.AutoMigrate(&models.Account{}, &models.Bottle{})

	adb = models.NewAccountDB(db)
	bdb = models.NewBottleDB(db)
	db.DB().SetMaxOpenConns(settings.MaxOpenConns)
	db.DB().SetMaxIdleConns(settings.MaxIdleConns)
	// Create service
	service := goa.New("API")
	logger := log15.New()
	service.WithLogger(goalog15.New(logger))

	// Setup middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.Recover())

	// Mount swagger controller onto service
	sc := NewSwagger(service)
	app.MountSwaggerController(service, sc)

	// Mount "account" controller
	c := NewAccountController(service)
	app.MountAccountController(service, c)
	// Mount "bottle" controller
	c2 := NewBottleController(service)
	app.MountBottleController(service, c2)
	// Start service, listen on port 8081
	service.ListenAndServe(":8081")
}
