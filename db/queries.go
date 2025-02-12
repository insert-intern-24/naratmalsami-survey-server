package db

import (
	"gorm.io/gorm"
	"log"
	"naratmalsami-survey-server/db/model"
)

func (db *DataDB) GetWordsInRange(startId, endId uint) ([]model.Words, error) {
	var words []model.Words
	result := db.Where("word_id BETWEEN ? AND ?", startId, endId).Table("words").Find(&words)
	if result.Error != nil {
		log.Println("조회 에러: ", result.Error)
		return nil, result.Error
	}
	return words, nil
}

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
