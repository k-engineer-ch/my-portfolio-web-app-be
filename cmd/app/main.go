package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	supa "github.com/nedpals/supabase-go"
)

func main() {
	fmt.Println(">>> Start API")

	// Supabaseへの接続情報「Project URL」を環境変数から取得
	supabaseUrl := os.Getenv("SUPABASE_URL")
	// Supanaseへの接続情報「Project API keys」を環境変数から取得
	supabaseKey := os.Getenv("SUPABASE_KEY")

	// Supabaseへの接続クライアント作成
	supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	// Select処理（サンプル）
	var results map[string]interface{}
	err := supabase.DB.From("expenses").Select("*").Single().Execute(&results)
	if err != nil {
		panic(err)
	}

	// 検索結果を出力
	fmt.Println(results)

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
