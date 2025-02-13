package model

// Words 단어 테이블
type Words struct {
	WordId       uint   `gorm:"primaryKey" json:"word_id"`
	OriginalWord string `gorm:"column:original_word;type:varchar(30);uniqueIndex;not null" json:"original_word"`
	RefinedWord  string `gorm:"column:refined_word;type:varchar(30);not null" json:"refined_word"`
	Meaning      string `gorm:"column:meaning;type:text" json:"meaning"`
}

// Users 유저 테이블
type Users struct {
	UserId string `gorm:"type:char(36);primaryKey;default:(UUID())" json:"user_id"`
}

// Voted 단어 평가 테이블
type Voted struct {
	WordId    uint   `gorm:"primaryKey" json:"word_id"`
	UserId    string `gorm:"primaryKey" json:"user_id"`
	Rating    int    `gorm:"column:rating;not null" json:"rating"`
	AtCreated string `gorm:"column:at_created;type:timestamp;not null" json:"at_created"`
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
