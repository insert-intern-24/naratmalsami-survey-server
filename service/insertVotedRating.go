package service

import (
	"encoding/json"
	"log"
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
			log.Printf("Error decoding request body: %v", err)
			return
		}

		// 사용자 유효성 검사
		if req.Who == nil {
			http.Error(w, "사용자 정보가 없습니다", http.StatusBadRequest)
			log.Println("User information is missing")
			return
		}

		var userId uint
		if isValidUser, id := db.SearchUser(*req.Who); !isValidUser {
			http.Error(w, "비정상적인 사용자 접근", http.StatusBadRequest)
			log.Println("Invalid user access")
			return
		} else {
			userId = id
		}

		// 비지니스 로직 요청
		if err := db.InsertRating(req, userId); err != nil {
			http.Error(w, "투표 데이터 입력 실패", http.StatusInternalServerError)
			log.Printf("Error inserting rating: %v", err)
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
			log.Printf("Error encoding response: %v", err)
		}
	}
}
