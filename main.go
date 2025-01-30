package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"naratmalsami-survey-server/db"
	"naratmalsami-survey-server/routes"
	"net/http"
	"os"
	"time"
)

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("환경 파일 로드 실패: %w", err)
	}
	return nil
}

func startServer(serverURL string, r http.Handler) error {
	srv := &http.Server{
		Handler:      r,
		Addr:         serverURL,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Printf("서버가 http://%s에서 실행 중입니다.\n", serverURL)
	return srv.ListenAndServe()
}

func main() {
	if err := loadEnv(); err != nil {
		log.Fatal(err)
	}

	wordDB, err := db.Connect_DB()
	if err != nil {
		log.Fatal("데이터베이스 초기화 실패:", err)
	}

	serverURL := os.Getenv("SERVER_URL")

	r := routes.SetupRouter(wordDB)

	if err := startServer(serverURL, r); err != nil {
		log.Fatal(err)
	}
}
