package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
type car struct {
	Speed int    `json:"speed"`
	Make  string `json:"make"`
}


func main(){

	//writeJson()
	readJson()
}

func writeJson(){
	c := car{
		Speed: 10,
		Make:  "Tesla",
	}
	dat, err := json.Marshal(c)
	if err != nil {
		fmt.Printf("Error - %s\n", err.Error())
		os.Exit(2)
	}
	err = ioutil.WriteFile("./file.json", dat, 0644)
	if err != nil {
		fmt.Printf("Error - %s\n", err.Error())
		os.Exit(2)
	}
}

func readJson(){
	dat, err := ioutil.ReadFile("./file.json")
	if err != nil {
		log.Printf("Error - %s\n", err.Error())
		os.Exit(2)
	}
	log.Printf("Dat - %s\n", string(dat))

	c := car{}
	err = json.Unmarshal(dat, &c)
	if err != nil {
		log.Printf("Error - %s\n", err.Error())
		os.Exit(2)
	}
	log.Printf("Data from file - %v\n",c)

}
