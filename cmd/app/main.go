package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println(">>> Start API")

	supabaseUrl := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_KEY")

	fmt.Println(supabaseUrl)
	fmt.Println(supabaseKey)

	// supabaseUrl := "https://glfgqyvbbqkpxrxizccv.supabase.co"
	// supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImdsZmdxeXZiYnFrcHhyeGl6Y2N2Iiwicm9sZSI6ImFub24iLCJpYXQiOjE3MTA5MDQyNDgsImV4cCI6MjAyNjQ4MDI0OH0.PFX1AZP2DJwgQWrqI3zNJDi9WUs8UpIJS2fRZzol8Uo"
	// supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	// var results map[string]interface{}
	// err := supabase.DB.From("countries").Select("*").Single().Execute(&results)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(results) // Selected rows

	r := gin.Default()

	r.GET("/sample", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "サンプル API",
			"int":     1,
			"map": map[string]int{
				"key1": 100,
				"key2": 200,
			},
		})
	})
	r.Run()
}
