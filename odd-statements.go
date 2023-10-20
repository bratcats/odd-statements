package odd_statements

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type catFact struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func CatRandomStatement() string {
	resp, err := http.Get("https://catfact.ninja/fact")
	if err != nil {
		panic(err)
	}

	byteResp, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	catFact := catFact{}
	err = json.Unmarshal(byteResp, &catFact)
	if err != nil {
		panic(err)
	}
	log.Printf("The fact: %q\n", catFact.Fact)

	return catFact.Fact
}
