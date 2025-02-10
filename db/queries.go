package db

import (
	"gorm.io/gorm"
	"log"
	"naratmalsami-survey-server/db/model"
)

func (db *WordDB) GetWordsInRange(startId, endId uint) ([]model.WordOfRating, error) {
	var words []model.WordOfRating
	result := db.Where("id BETWEEN ? AND ?", startId, endId).Table("word_of_rating").Find(&words)
	if result.Error != nil {
		log.Println("조회 에러: ", result.Error)
		return nil, result.Error
	}
	return words, nil
}

func (db *WordDB) UpdateRating(ids []uint, ratings []int) error {
	for i := 0; i < len(ids); i++ {
		if err := db.Model(&model.WordOfRating{}).Where("id = ?", ids[i]).
			Updates(map[string]interface{}{
				"whole_rating": gorm.Expr("whole_rating + ?", ratings[i]),
				"member":       gorm.Expr("member + 1"),
			}).Error; err != nil {
			return err
		}
	}
	return nil
}
