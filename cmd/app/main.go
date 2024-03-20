package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	// supa "github.com/nedpals/supabase-go"
	supabase "github.com/supabase-community/supabase-go"
)

type Expense struct {
}

func main() {
	fmt.Println(">>> Start API")

	// .envファイルから環境変数を読み込む
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Supabaseへの接続情報「Project URL」を環境変数から取得
	supabaseUrl := os.Getenv("SUPABASE_URL")
	// Supanaseへの接続情報「Project API keys」を環境変数から取得
	supabaseKey := os.Getenv("SUPABASE_KEY")
	// Supabaseへの接続クライアント作成
	// supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	fmt.Println("Project URL:", supabaseUrl)
	fmt.Println("Project API keys:", supabaseKey)

	client, err := supabase.NewClient(supabaseUrl, supabaseKey, nil)
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	// data, count, err := client.From("expenses").Select("", "exact", false).Execute()
	data, count, err := client.From("expenses").Select("*", "exact", false).Execute()

	if err != nil {
		fmt.Println("sql error", err)
	} else {
		fmt.Println(data)
		fmt.Println(count)
	}

	// // Select処理（サンプル）
	// var results map[string]interface{}
	// err := supabase.DB.From("expenses").Select("*").Single().Execute(&results)
	// // err := supabase.DB.From("expenses").Select("*").Execute(&results)
	// if err != nil {
	// 	panic(err)
	// }

	// // 検索結果を出力
	// fmt.Println(results)

	// r := gin.Default()

	// r.GET("/sample", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "サンプル API",
	// 		"int":     1,
	// 		"map": map[string]int{
	// 			"key1": 100,
	// 			"key2": 200,
	// 		},
	// 	})
	// })
	// r.Run()
}
