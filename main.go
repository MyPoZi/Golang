package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}
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

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// TODO パラメータ2項目以上のエラー処理

	// TODO 数値以外のエラー処理
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
