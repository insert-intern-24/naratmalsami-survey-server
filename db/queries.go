package db

import (
	"fmt"
	"naratmalsami-survey-server/db/model"
	"time"
)

func (db *DataDB) SearchUser(who string) (isValidUser bool, userId uint) {
	err := db.Model(&model.Users{}).Select("users_id").Where("who = ?", who).Scan(&userId).Error
	if err != nil {
		// 에러가 발생한 경우
		return false, 0
	}

	// 사용자 ID가 0인 경우 유효하지 않은 사용자로 판단
	if userId == 0 {
		return false, 0
	}

	return true, userId
}
func (db *DataDB) InsertRating(body model.VotedRequestBody, userId uint) error {
	for _, word := range body.Words {
		if err := db.Create(&model.Voted{
			WordId:    uint(word.WordId),
			UserId:    userId,
			Rating:    word.Rating,
			AtCreated: time.Now(),
		}).Error; err != nil {
			return fmt.Errorf("failed to insert rating for word ID %d: %w", word.WordId, err)
		}
	}
	return nil
}

func (db *DataDB) CreateUser() (string, error) {
	user := model.Users{}
	result := db.DB.Create(&user)
	if result.Error != nil {
		return "", fmt.Errorf("failed to create user: %w", result.Error)
	}

	var createdUser model.Users
	if err := db.DB.First(&createdUser, user.UsersId).Error; err != nil {
		return "", fmt.Errorf("failed to retrieve created user: %w", err)
	}

	return createdUser.Who, nil
}

func (db *DataDB) GetLeastVotedWords(limit int) ([]model.Words, error) {
	var words []model.Words

	// 투표 수 집계 서브쿼리
	voteCountSubQuery := db.Table("voteds").
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
		Order("COALESCE(vote_counts.vote_count, 0) ASC, words.weight DESC, LENGTH(words.meaning) ASC").
		Limit(limit).
		Find(&words).Error

	if err != nil {
		return nil, err
	}
	return words, nil
}
