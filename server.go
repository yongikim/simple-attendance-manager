package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Attendance struct {
	Name string `json:"name"`
	In   string `json:"in"`
	Out  string `json:"out"`
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	var attendances []Attendance
	for j := 0; j < 10; j++ {
		name := fmt.Sprintf("Student %d", j)
		in := time.Now()
		out := in.Add(time.Hour * 10)
		attendances = append(
			attendances,
			Attendance{
				Name: name,
				In:   fmt.Sprintf("%d", in.Unix()),
				Out:  fmt.Sprintf("%d", out.Unix()),
			},
		)
	}
	json.NewEncoder(w).Encode(attendances)
}

func main() {
	http.HandleFunc("/", indexHandler)
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}
