package db

import (
	"fmt"
	"naratmalsami-survey-server/db/model"
	"time"
)

func (db *DataDB) SearchUser(who string) (isValidUser bool, userId uint) {
	err := db.Model(&model.Users{}).Select("user_id").Where("who = ?", who).Scan(&userId).Error
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
	if err := db.DB.First(&createdUser, user.UserId).Error; err != nil {
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

	// 메인 쿼리 실행
	err := db.Table("words").
		Select("words.word_id, words.original_word, words.refined_word, words.meaning, words.weigh").
		Joins("LEFT JOIN (?) AS vote_counts ON words.word_id = vote_counts.word_id", voteCountSubQuery).
		Order("COALESCE(vote_counts.vote_count, 0) ASC, words.weigh DESC").
		Limit(limit).
		Find(&words).Error

	if err != nil {
		return nil, err
	}
	return words, nil
}

func (db *DataDB) GetRankingOfWho(who string) (model.RankingResponseBody, error) {
	var count int64
	if err := db.Table("Rankings").Where("who = ?", who).Count(&count).Error; err != nil {
		return model.RankingResponseBody{Ranking: -1, Code: -1}, err
	}

	if count == 0 {
		return model.RankingResponseBody{Ranking: -1, Code: -1}, nil
	}

	var result model.RankingResponseBody
	err := db.Table("Rankings").Select("ranking, code").Where("who = ?", who).Scan(&result).Error
	if err != nil {
		return model.RankingResponseBody{Ranking: -1, Code: -1}, err
	}
	return result, nil
}
