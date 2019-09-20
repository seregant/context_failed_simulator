package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

// People is
type People struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			time.Sleep(100 * time.Second)
			var listPeople []People
			for i := 0; i < 3; i++ {
				var tmpPeople People
				tmpPeople.Name = "people"
				tmpPeople.Address = "address"
				listPeople = append(listPeople, tmpPeople)
			}
			json.NewEncoder(w).Encode(listPeople)
		case http.MethodPost:
			time.Sleep(120 * time.Second)
			type Response struct {
				Message string `json:"message"`
			}
			decoder := json.NewDecoder(r.Body)
			var p People
			err := decoder.Decode(&p)
			if err != nil {
				log.Println(err)
			}
			log.Println(p.Name, p.Address)
			var res Response
			res.Message = "success"
			json.NewEncoder(w).Encode(res)
		default:
			log.Fatal("unknown request")
		}
	})
	http.ListenAndServe(":2222", nil)
}
