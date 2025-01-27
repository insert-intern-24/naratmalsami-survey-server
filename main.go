package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func loadEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("환경 파일 로드 실패: %w", err)
	}
	return nil
}

func startServer(serverURL string) error {
	fmt.Printf("서버가 https://%s에서 실행중입니다.\n", serverURL)
	if err := http.ListenAndServe(serverURL, nil); err != nil {
		return fmt.Errorf("서버 실행 중 오류: %w", err)
	}
	return nil
}

func main() {
	if err := loadEnv(); err != nil {
		log.Fatal(err)
	}

	serverURL := os.Getenv("SERVER_URL")
	if serverURL == "" {
		log.Fatal("SERVER_URL 환경 변수가 설정되지 않았습니다.")
	}

	if err := startServer(serverURL); err != nil {
		log.Fatal(err)
	}
}
