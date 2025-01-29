package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

// Table :: word_of_rating :: 단어에 대한 평가 수치를 저장하는 테이블
type WordOfRating struct {
	ID           uint   `gorm:"primaryKey"`
	OriginalWord string `gorm:"type:varchar(30);uniqueIndex;not null"`
	RefinedWord  string `gorm:"type:varchar(30);not null"`
	Meaning      string `gorm:"type:text"`
	WholeRating  int    `gorm:"default:0"`
}

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

	err := db.AutoMigrate(&WordOfRating{})
	if err != nil {
		return nil, err
	}

	return &WordDB{db}, nil
}

func (db *WordDB) GetWordsInRange(startId, endId uint) ([]WordOfRating, error) {
	var words []WordOfRating
	result := db.Where("id BETWEEN ? AND ?", startId, endId).Find(&words)

	if result.Error != nil {
		log.Println("err : ", result.Error)
		return nil, result.Error
	}

	return words, nil
}
func (db *WordDB) UpdateRating(ids []uint, ratings []int) error {
	for i := 0; i < len(ids); i++ {
		if err := db.Model(&WordOfRating{}).Where("id = ?", ids[i]).
			Update("whole_rating", gorm.Expr("whole_rating + ?", ratings[i])).Error; err != nil {
			return err
		}
	}
	return nil
}
