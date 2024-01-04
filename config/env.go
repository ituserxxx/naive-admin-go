package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)
func Init()  {
	err := godotenv.Load(".env") // 加载.env文件
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}
	println(os.Getenv("Mysql"))
}
