package db

import (
	"gorm.io/gorm"
	"naratmalsami-survey-server/db/model"
)

func (db *DataDB) UpdateRating(ids []uint, ratings []int) error {
	for i := 0; i < len(ids); i++ {
		if err := db.Model(&model.Words{}).Where("word_id = ?", ids[i]).
				Updates(map[string]interface{}{
					"whole_rating": gorm.Expr("whole_rating + ?", ratings[i]),
					"member":       gorm.Expr("member + 1"),
				}).Error; err != nil {
			return err
		}
	}
	return nil
}

func (db *DataDB) CreateUser() (string, error) {
	user := model.Users{}
	result := db.DB.Create(&user)
	if result.Error != nil {
		return "", result.Error
	}
	return user.UserId, nil
}

func (db *DataDB) GetLeastVotedWords(limit int) ([]model.Words, error) {
	var words []model.Words

	// 투표 수 집계 서브쿼리
	voteCountSubQuery := db.Table("voted").
		Select("word_id, COUNT(rating) AS vote_count").
		Group("word_id")

	// 최소 의미 길이 서브쿼리
	minMeaningLengthSubQuery := db.Table("words as w2").
		Select("MIN(LENGTH(w2.meaning))").
		Where("w2.word_id = words.word_id")

	// 메인 쿼리 실행 (vote_count는 출력 제외)
	err := db.Table("words").
		Select("words.word_id, words.original_word, words.refined_word, words.meaning").
		Joins("LEFT JOIN (?) AS vote_counts ON words.word_id = vote_counts.word_id", voteCountSubQuery).
		Where("LENGTH(words.meaning) = (?)", minMeaningLengthSubQuery).
		Order("COALESCE(vote_counts.vote_count, 0) ASC, LENGTH(words.meaning) ASC").
		Limit(limit).
		Find(&words).Error

	if err != nil {
		return nil, err
	}
	return words, nil
}
