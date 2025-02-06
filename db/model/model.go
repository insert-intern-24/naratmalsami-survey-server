package model

// 단어에 대한 평가 수치를 저장하는 테이블
type WordOfRating struct {
	ID           uint   `gorm:"primaryKey"`
	OriginalWord string `gorm:"type:varchar(30);uniqueIndex;not null"`
	RefinedWord  string `gorm:"type:varchar(30);not null"`
	Meaning      string `gorm:"type:text"`
	WholeRating  int    `gorm:"default:0"`
	Member       int    `gorm:"default:0"`
}
