package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Dino struct {
	Name        string
	Description string
}

func main() {
	resp, err := http.Get("https://dino-facts-api.shultzlab.com/dinosaurs/random")
	if err != nil {
		log.Fatalln(err)
	}

	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var dino Dino
	err = json.Unmarshal(body, &dino)
	if err != nil {
		fmt.Printf("Error format Dino Fact!")
		return
	}
	fmt.Printf("DINO FACT!! \nName: %s \nDescription: %s\n", dino.Name, dino.Description)
	fmt.Printf("               __\n              / _)\n     _.----._/ /\n    /         /\n __/ (  | (  |\n/__.-'|_|--|_|")

	_, err = fmt.Scanf("h")
	if err != nil {
		return
	}
}
