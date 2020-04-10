package rtcmlib

import (
	"encoding/json"
	"fmt"
	"log"
)

func Print_json(dict interface{}) {
	b, err := json.MarshalIndent(dict, "", "  ")
	Chk(err)
	fmt.Println(string(b))
}

// Chk - Function for check errors
func Chk(err error) {
	if err != nil {
		log.Printf("%T\n", err)
		log.Fatal(err)
	}
}