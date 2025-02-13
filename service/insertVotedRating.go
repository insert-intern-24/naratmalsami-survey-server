package service

import (
	"encoding/json"
	"naratmalsami-survey-server/db"
	"naratmalsami-survey-server/db/model"
	"net/http"
)

func InsertVotedRating(db *db.DataDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 요청 데이터 파싱
		var req model.VotedRequestBody
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "잘못된 요청 데이터", http.StatusBadRequest)
			return
		}
		// 비지니스 로직 요청
		db.InsertRating(req)

		// 응답 데이터 센터
		response := map[string]string{
			"status": "success",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}
