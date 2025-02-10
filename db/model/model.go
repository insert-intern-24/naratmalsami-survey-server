package model

// 단어에 대한 평가 수치를 저장하는 테이블
type WordOfRating struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	OriginalWord string `gorm:"column:original_word;type:varchar(30);uniqueIndex;not null" json:"original_word"`
	RefinedWord  string `gorm:"column:refined_word;type:varchar(30);not null" json:"refined_word"`
	Meaning      string `gorm:"column:meaning;type:text" json:"meaning"`
	WholeRating  int    `gorm:"column:whole_rating;default:0" json:"whole_rating"`
	Member       int    `gorm:"column:member;default:0" json:"member"`
}
