package service

import (
	"encoding/json"
	"log"
	"naratmalsami-survey-server/db"
	"naratmalsami-survey-server/db/model"
	"net/http"
)

func GetSheetService(db *db.WordDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		words, err := db.GetWordsInRange(1, 10)
		if err != nil {
			http.Error(w, "데이터베이스 조회 실패", http.StatusInternalServerError)
			return
		}

		data := map[string][]model.WordOfRating{
			"data": words,
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Println("JSON 직렬화 실패: ", err)
			http.Error(w, "응답 생성 실패", http.StatusInternalServerError)
			return
		}
	}
}
