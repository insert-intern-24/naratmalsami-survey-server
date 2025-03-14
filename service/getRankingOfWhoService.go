package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"naratmalsami-survey-server/db"
	"naratmalsami-survey-server/db/model"
	"net/http"
)

func GetRankingOfWhoService(db *db.DataDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 요청 데이터 파싱
		vars := mux.Vars(r)
		who := vars["who"]
		if who == nil {
			http.Error(w, "잘못된 요청 데이터", http.StatusBadRequest)
			log.Printf("/ranking doesn't find who")
		}

		// 사용자 유효성 검사
		if isValidUser, _ := db.SearchUser(who); !isValidUser {
			log.Println("/ranking Unauthorized user access")
			http.Error(w, "비정상적인 사용자 접근", http.StatusBadRequest)
			return
		}

		// ranking 등수 조회
		ranking, err := db.GetRankingOfWho(who)
		if err != nil {
			log.Printf("/ranking Database query failed: %v", err)
			http.Error(w, "데이터베이스 조회 실패", http.StatusInternalServerError)
			return
		}

		// 응답 데이터 생성
		data := map[string]interface{}{
			"ranking": ranking,	
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		log.Println("/ranking success")
		if err := json.NewEncoder(w).Encode(data); err != nil {
			log.Printf("/ranking JSON serialization failed: %v", err)
			http.Error(w, "응답 생성 실패", http.StatusInternalServerError)
			return
		}
	}
}
