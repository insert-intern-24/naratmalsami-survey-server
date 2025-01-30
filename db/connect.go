package db

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"naratmalsami-survey-server/db/model"
)

type WordDB struct {
	*gorm.DB
}

func Connect_DB() (*WordDB, error) {
	dsn := os.Getenv("DATABASE_DSN") // 환경 변수에서 DSN 읽기
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("데이터베이스 연결 실패: %w", err)
	}

	fmt.Println("MySQL에 성공적으로 연결되었습니다!")

	// 모델(table) 마이그레이션
	err = db.AutoMigrate(&model.WordOfRating{})
	if err != nil {
		return nil, err
	}

	return &WordDB{DB: db}, nil
}
