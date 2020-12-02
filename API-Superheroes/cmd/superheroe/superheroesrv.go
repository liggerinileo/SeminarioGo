package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
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

	// database.PopulateDatabase(db, "./test/mock.sql")
	database.PopulateDatabase(db, cfg)

	superheroeService, _ := service.NewSuperheroeService(cfg, db)

	makeEndpoints(superheroeService)

}

func makeEndpoints(superheroeService service.SuperheroeService) {
	router := gin.Default()
	api := router.Group("/api")
	v1 := api.Group("/v1")
	sword := v1.Group("/superheroe")

	sword.GET("/", func(c *gin.Context) {
		c.JSON(200, superheroeService.FindAll())
	})

	sword.GET("/:id", func(c *gin.Context) {
		id := c.Param("id")
		parseInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		c.JSON(200, superheroeService.FindById(int(parseInt)))
	})

	sword.PUT("/:id", func(c *gin.Context) {
		id := c.Param("id")
		sword := new(service.Superheroe)
		err := c.Bind(superheroe)
		parseInt, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		updateSword, err := superheroeService.UpdateSuperheroe(*superheroe, int(parseInt))
		c.JSON(200, UpdateSuperheroe)
	})

	sword.POST("/", func(c *gin.Context) {
		sword := new(service.Superheroe)
		err := c.Bind(superheroe)
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
		}
		updateSword, err := swordService.AddSuperheroe(*superheroe)
		c.JSON(200, updateSuperheroe)
	})

	router.Run()
}

func readConfig() *config.Config {
	configFile := flag.String("config", "./configs/config.yaml", "this is the service config")
	flag.Parse()

	cfg, err := config.LoadConfig(*configFile)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return cfg
}
