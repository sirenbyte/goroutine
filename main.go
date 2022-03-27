package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func f(i, i2 int) {
	n := time.Now()
	for i < i2 {
		v := strconv.Itoa(i)
		var jsonData = []byte(fmt.Sprintf(`"body"{"phone":"%s"}`, v))
		s, _ := http.Post("https://api.tengeda.kz/api/json/tengeda/checkPhone", "application/json", bytes.NewBuffer(jsonData))
		var jsonMap map[string]int
		b, _ := ioutil.ReadAll(s.Body)
		json.Unmarshal(b, &jsonMap)
		if jsonMap["code"] == 0 {
			fmt.Println(i, "Not registered")
		} else {
			fmt.Println(i, "Exist")
		}
		i++

	}
	fmt.Println(n, time.Now())
}

func getAll(w http.ResponseWriter, r *http.Request) {
	var jsonMap map[string]int
	b, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(b, &jsonMap)
	t1 := fmt.Sprintf("7%d0000000", jsonMap["code"])
	t2 := fmt.Sprintf("7%d0000100", jsonMap["code"])
	flag.Parse()
	i, _ := strconv.Atoi(t1)
	i2, _ := strconv.Atoi(t2)
	go f(i, i2)
}

func getOne(w http.ResponseWriter, r *http.Request) {
	s, err := http.Post("https://api.tengeda.kz/api/json/tengeda/checkPhone", "application/json", r.Body)
	if err != nil {
		fmt.Print(err)
	}
	var jsonMap map[string]int
	b, _ := ioutil.ReadAll(s.Body)
	json.Unmarshal(b, &jsonMap)
	if jsonMap["code"] == 0 {
		w.Write([]byte("No registered"))
	} else {
		w.Write([]byte("Exists"))
	}
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", getOne).Methods("POST")
	router.HandleFunc("/all", getAll).Methods("POST")
	log.Fatal(http.ListenAndServe(":80", router))
}
