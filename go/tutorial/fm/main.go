package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

const BaseURL = "https://swapi.dev/api/"

type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet
}

type AllPeople struct {
	People []Person `json:"results"`
}

func (p *Person) getHomeworld() {
	res, err := http.Get(p.HomeworldURL)
	if err != nil {
		log.Fatal(err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Fatal(err)
	}

	json.Unmarshal(bytes, &p.Homeworld)
}

func (p *AllPeople) getHomeworlds() {
	for i, person := range p.People {
		person.getHomeworld()
		p.People[i].Homeworld = person.Homeworld
	}
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	res, err := http.Get(BaseURL + "people")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := ioutil.ReadAll(res.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// no schema
	// a := map[string]interface{}{}
	// json.Unmarshal(data, &a)
	// n := a["results"].([]interface{})[0].(map[string]interface{})["name"]
	// fmt.Println(n)
	// w.Write([]byte(n.(string)))

	var people AllPeople
	if err := json.Unmarshal(bytes, &people); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	people.getHomeworlds()

	// 	for _, person := range people.People[:1] {
	// 		person.getHomeworld()
	// 		fmt.Printf("person: %v\n", person)
	// 	}
	var s []byte
	if s, err = json.Marshal(people); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(s)

}

func main() {
	// testGoRoutine()

	// wg.Add(1)
	// go printStuff()
	// wg.Wait()

	http.Handle("/people", http.HandlerFunc(GetPeople))
	http.HandleFunc("/people", GetPeople)
	fmt.Println("Serving on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

}

var wg sync.WaitGroup

func printStuff() {
	defer wg.Done()
	defer handlePanic()
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}
}

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("PANIC")
	}
}

func testGoRoutine() {
	go say("HELLO")
	say("There")
}

func say(s string) {
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(300 * time.Millisecond)
	}
}
