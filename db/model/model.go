package model

import "time"

// Words 단어 테이블
type Words struct {
	WordId       uint   `gorm:"primaryKey" json:"word_id"`
	OriginalWord string `gorm:"column:original_word;type:varchar(30);uniqueIndex;not null" json:"original_word"`
	RefinedWord  string `gorm:"column:refined_word;type:varchar(30);not null" json:"refined_word"`
	Meaning      string `gorm:"column:meaning;type:text" json:"meaning"`
	Weight       int    `gorm:"column:weight;type:int" json:"weight"`
}

// Users 유저 테이블
type Users struct {
	UserId uint   `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Who    string `gorm:"type:char(36);default:(UUID())" json:"who"`
}

// Voted 단어 평가 테이블
type Voted struct {
	WordId    uint      `gorm:"primaryKey" json:"word_id"`
	UserId    uint      `gorm:"primaryKey" json:"user_id"`
	Rating    int       `gorm:"column:rating;not null" json:"rating"`
	AtCreated time.Time `gorm:"column:at_created;type:timestamp;not null" json:"at_created"`
}

type Ranking struct {
	Who     string `gorm:"type:char(36);default:(UUID())" json:"who"`
	Ranking int    `gorm:"column:ranking;not null" json:"ranking"`
	Code  	int    `gorm:"column:code;not null" json:"code"`
}

// [POST] /sheet request body
type SheetRequestBody struct {
	Who *string `json:"who"`
}

// [POST] /voted request body
type VotedRequestBody struct {
	Who   *string      `json:"who"`
	Words []WordRating `json:"words"`
}

type WordRating struct {
	WordId uint `json:"word_id"`
	Rating int  `json:"rating"`
}

type RankingResponseBody struct {
	Ranking int `json:"ranking"`
	Code  	int `json:"code"`
}