package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/liggerinileo/SeminarioGo/API-Superheroes/internal/config"
	"github.com/liggerinileo/SeminarioGo/API-Superheroes/internal/database"
	"github.com/liggerinileo/SeminarioGo/API-Superheroes/internal/service"
)

func main() {
	cfg := readConfig()

	db, err := database.NewDatabase(cfg)
	defer db.Close()

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	superheroeService, _ := service.NewSuperheroeService(cfg, db)

	makeEndpoints(superheroeService)

}

func makeEndpoints(superheroeService service.SuperheroeService) {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	superheroe := v1.Group("/superheroe")

	superheroe.GET("/", func(c *gin.Context) {
		c.JSON(200, superheroeService.FindAll())
	})

	superheroe.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		parseInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.JSON(200, superheroeService.FindByID(int(parseInt)))
	})

	superheroe.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		superheroe := new(service.Superheroe)
		parseInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		updateSuperheroe, err := superheroeService.UpdateSuperheroe(*superheroe, int(parseInt))
		c.JSON(200, updateSuperheroe)
	})

	superheroe.POST("/", func(c *gin.Context) {
		superheroe := new(service.Superheroe)
		err := c.Bind(superheroe)
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
		}
		updateSuperheroe, err := superheroeService.AddSuperheroe(*superheroe)
		c.JSON(200, updateSuperheroe)
	})

	router.Run()
}

// CreateSchema ...
func CreateSchema(db *sqlx.DB) error {
	schema := `CREATE TABLE IF NOT EXISTS superheroe (
		id integer primary key autoincrement,
		name varchar,
		strength integer,
		speed integer,
		fightingSkills integer,
		intelligence integer
		);`

	// execute a query on the server
	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./config/config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
