package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Json struct {
	Status     int      `json:"status"`
	Result     string   `json:"result"`
	Parameters int      `json:"parm"`
	Nya        []string `json:"nya"`
}

func returnResponse(w http.ResponseWriter, body Json) {
	res, err := json.Marshal(body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	return

}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["value"]
	nya := make([]string, 1)

	// パラメータが1個未満の場合エラー処理
	if !ok || len(keys[0]) < 1 {
		body := Json{http.StatusOK, "Url Param 'value' is missing", -1, nya}
		returnResponse(w, body)

		return
	}

	// パラメータが2個以上の場合エラー処理
	if len(r.URL.Query()) >= 2 {
		body := Json{http.StatusOK, "Too many Url Param", -1, nya}
		returnResponse(w, body)

		return
	}

	key, _ := strconv.Atoi(keys[0])

	if key <= 0 {
		body := Json{http.StatusOK, "invalid value", key, nya}
		returnResponse(w, body)

		return
	}
	nya[0] = "にゃーん"
	for count := 1; count < key; count++ { // ex. 1の場合 にゃーんx1, 10の場合 にゃーんx10
		nya = append(nya, "にゃーん")
	}
	body := Json{http.StatusOK, "ok", key, nya}

	returnResponse(w, body)

	return
}

func main() {
	http.HandleFunc("/api/nya-n", jsonHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe", err)
	}
}
