package main

import (
  "github.com/gin-gonic/gin"
  "database/sql"
  "fmt"
  "os"
)

func main() {
  dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=%s host=%s password=%s",
      os.Getenv("TEXASFISHIN_API_DB_USER"),
      os.Getenv("TEXASFISHIN_API_DB_NAME"),
      os.Getenv("TEXASFISHIN_API_DB_SSLMODE"),
      os.Getenv("TEXASFISHIN_API_DB_HOST"),
      os.Getenv("TEXASFISHIN_API_DB_PASSWORD"))
  db, err := sql.Open("postgres", dbinfo)
  if err != nil {
    fmt.Println("error")
    fmt.Println(err)
  }

  env := &Env{db: db}
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

  r.GET("/lakes", env.lakesList)
  r.GET("/lakes/:id", env.lakeShow)
  r.GET("/lakes/:id/records", env.lakeRecordsList)
  r.GET("/fish_types", env.fishTypesList)

	r.Run() // listen and serve on 0.0.0.0:8080
}
