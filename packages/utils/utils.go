// Package utils
package utils

import (
	"log"
	"encoding/json"
)

type element struct {
	Key   string
	Value int
}

func Unmarshale(json string) {
	d := json.NewDecoder(r)
	t, err := d.Token()
	if err != nil || t != json.Delim('{') {
		log.Fatal("expected object")
	}
	var result []*element
	for d.More() {
		k, err := d.Token()
		if err != nil {
			log.Fatal(err)
		}
		var v element
		if err := d.Decode(&v); err != nil {
			log.Fatal(err)
		}
		result = append(result, &v)
	}
}
