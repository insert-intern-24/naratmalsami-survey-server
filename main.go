package main

import (
	"fmt"
	"github.com/gorilla/handlers"
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
	// CORS 처리
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),                                              // 모든 출처 허용
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),                   // 허용할 HTTP 메소드
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Cache-Control"}), // 허용할 헤더
	)

	// CORS 미들웨어를 라우터에 적용
	r = corsHandler(r)

	// 서버 설정
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
