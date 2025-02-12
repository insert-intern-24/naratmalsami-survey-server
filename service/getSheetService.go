package service

import (
	"encoding/json"
	"log"
	"naratmalsami-survey-server/db"
	"naratmalsami-survey-server/db/model"
	"net/http"
)

func GetSheetService(db *db.DataDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 요청 데이터 파싱
		var req model.SheetRequestBody
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "잘못된 요청 데이터", http.StatusBadRequest)
			return
		}

		// request body, who 필드 체크
		var who *string = req.Who
		if who == nil {
			user, err := db.CreateUser()
			if err != nil {
				http.Error(w, "사용자 생성 실패", http.StatusInternalServerError)
				return
			}
			who = &user
		}

		// sheet 데이터 조회
		words, err := db.GetWordsInRange(1, 10)
		if err != nil {
			http.Error(w, "데이터베이스 조회 실패", http.StatusInternalServerError)
			return
		}

		// 응답 데이터 생성
		data := map[string]interface{}{
			"data": words,
			"who":  *who,
		}

		// 응답
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Println("JSON 직렬화 실패: ", err)
			http.Error(w, "응답 생성 실패", http.StatusInternalServerError)
			return
		}
	}
}
