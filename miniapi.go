package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func hoursHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "%00dh%00d", time.Now().Hour(), time.Now().Minute())
	default:
		fmt.Fprintf(w, "Ce n'est pas le bon type de requète !\n")
		return
	}
}

func diceHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		rand.Seed(time.Now().UnixNano())
		dice := rand.Intn(1000)
		fmt.Fprintf(w, "%d\n", dice)
	default:
		fmt.Fprintf(w, "Ce n'est pas le bon type de requète !\n")
		return
	}
}

func dicesHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		if req.URL.Query().Get("type") != "" {
			data, err := strconv.Atoi(req.URL.Query().Get("type")[1:])
			if err != nil {
				fmt.Fprintf(w, "Une erreur s'est produite: %s\n", err)
				return
			}
			for i := 0; i < 15; i++ {
				rand.Seed(time.Now().UnixNano())
				fmt.Fprintf(w, "%d ", rand.Intn(data))
			}
		} else {
			typeOfDice := [8]int{2, 4, 6, 8, 10, 12, 20, 100}
			for i := 0; i < 15; i++ {
				rand.Seed(time.Now().UnixNano())
				dice := typeOfDice[rand.Intn(8)]
				rand.Seed(time.Now().UnixNano())
				//fmt.Fprintf(w, "%d // ", dice)
				fmt.Fprintf(w, "%d ", rand.Intn(dice))
			}
		}
	default:
		fmt.Fprintf(w, "Ce n'est pas le bon type de requète !\n")
		return
	}
}

func randWordsHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		if len(req.PostFormValue("words")) > 0 {
			countWord := len(strings.Split(req.PostFormValue("words"), "%20"))
			SliceWord := strings.Split(req.PostFormValue("words"), "%20")

			for i := 0; i < countWord; i++ {
				rand.Seed(time.Now().Unix())
				randomIndex := rand.Intn(len(SliceWord))
				pick := SliceWord[randomIndex]
				fmt.Fprintf(w, "%s ", pick)
			}
		} else {
			fmt.Fprintf(w, "Le mot est vide !\n")
			return
		}
	default:
		fmt.Fprintf(w, "Ce n'est pas le bon type de requète !\n")
		return
	}
}

func main() {
	http.HandleFunc("/", hoursHandler)
	http.HandleFunc("/dice", diceHandler)
	http.HandleFunc("/dices", dicesHandler)
	http.HandleFunc("/randomize-words", randWordsHandler)
	http.ListenAndServe(":4567", nil)
}
