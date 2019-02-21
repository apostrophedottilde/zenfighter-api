package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"bitcrunchy.com/zenfighter-api/domain"
	"bitcrunchy.com/zenfighter-api/engine"
)

type ErrResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func HandleFindAll(e engine.Engine) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {

		knights := e.ListKnights()

		if knights == nil {
			resp = buildResponse(resp, http.StatusOK)
			resp.Write([]byte("[]"))
			return
		}

		resp = buildResponse(resp, http.StatusOK)
		data, _ := json.Marshal(knights)
		resp.Write([]byte(data))
	})
}

func HandleFindOne(e engine.Engine) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		id := strings.TrimPrefix(req.URL.Path, "/knight/")
		retStatus := http.StatusOK
		k, error := e.GetKnight(id)

		if error != nil || k == nil {
			retStatus = http.StatusNotFound
		}

		resp = buildResponse(resp, retStatus)

		if retStatus != http.StatusNotFound {
			data, _ := json.Marshal(k)
			resp.Write([]byte(data))
		} else {
			errResponse(resp, "404", "Knight #"+id+" not found.")
		}
	})
}

func HandleCreate(e engine.Engine) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		r := &domain.Knight{}

		err := json.NewDecoder(req.Body).Decode(&r)

		if err != nil {
			errResponse(resp, "400", "errr")
		} else {
			// validate post data
			if r.Name == "" || r.WeaponPower == 0 || r.Strength == 0 {
				errResponse(resp, "400", "errr")
			} else {
				e.Create(r)
				resp = buildResponse(resp, http.StatusCreated)
			}
		}
	})
}

func HandleFight(e engine.Engine) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		f1 := req.URL.Query().Get("fighter1")
		f2 := req.URL.Query().Get("fighter2")

		if f1 == "" || f2 == "" {
			panic("Both fighters must be specified in the query parameters.")
		} else {

		}

		winner := e.Fight(f1, f2)

		if winner == nil {
			resp = buildResponse(resp, http.StatusConflict)
			resp.Write([]byte(""))
		} else {
			data, _ := json.Marshal(winner)
			resp = buildResponse(resp, http.StatusOK)
			resp.Write([]byte(data))
		}
	})
}

func buildResponse(resp http.ResponseWriter, status int) http.ResponseWriter {
	resp.Header().Add("Content-Type", "application/json")
	resp.WriteHeader(status)
	return resp
}

func errResponse(resp http.ResponseWriter, code string, message string) {
	err := ErrResponse{
		Code:    code,
		Message: message,
	}
	data, _ := json.Marshal(err)
	d, _ := strconv.Atoi(code)
	resp.WriteHeader(d)
	resp.Write([]byte(data))
}
