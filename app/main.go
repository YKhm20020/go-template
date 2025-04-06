package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Account struct {
	ID       int
	Name     string
	Password string
}

func ConnectDB() (*gorm.DB, error) {
	// データベース接続、.env ファイルが必要
	// 詳細は https://github.com/go-sql-driver/mysql#dsn-data-source-name を参照
	dsn := fmt.Sprintf(
		"%s:%s@tcp(db)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("データベース接続エラー: %v\n", err)
		return nil, err
	}

	// マイグレーションの確認
	err = db.AutoMigrate(&Account{})
	if err != nil {
		fmt.Printf("マイグレーションエラー: %v\n", err)
		return nil, err
	}

	return db, nil
}

func main() {
	// データベース接続
	db, err := ConnectDB()

	if err != nil {
		fmt.Printf("データベースとの接続に失敗\n")
		return
	}

	// dsn := "root@tcp(mysql:3306)/first?charset=utf8mb4&parseTime=True&loc=Local"

	// dsn := "db-user:db-password@tcp(db:3306)/go-template-db?charset=utf8mb4&parseTime=True&loc=Local"

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, Gin + MySQL + Docker + Air",
		})
	})

	// 全アカウント取得
	r.GET("/accounts", func(c *gin.Context) {
		var accounts []Account
		result := db.Find(&accounts)
		if result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": result.Error.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, accounts)
	})

	// IDによるアカウント取得
	r.GET("/accounts/:id", func(c *gin.Context) {
		var account Account
		id := c.Param("id")
		result := db.First(&account, id)
		if result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "アカウントが見つかりません",
			})
			return
		}
		c.JSON(http.StatusOK, account)
	})

	r.Run(":8080")
}
