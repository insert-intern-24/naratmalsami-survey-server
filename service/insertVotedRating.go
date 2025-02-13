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

		// 사용자 유효성 검사
		if req.Who == nil {
			http.Error(w, "사용자 정보가 없습니다", http.StatusBadRequest)
			return
		}

		if isValidUser := db.SearchUser(*req.Who); !isValidUser {
			http.Error(w, "비정상적인 사용자 접근", http.StatusBadRequest)
			return
		}

		// 비지니스 로직 요청
		if err := db.InsertRating(req); err != nil {
			http.Error(w, "투표 데이터 입력 실패", http.StatusInternalServerError)
			return
		}

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
