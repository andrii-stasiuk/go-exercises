package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

func words(r io.Reader) (even []string, odd []string) {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	strData := strings.Fields(string(buf))

	for _, str := range strData {

		cnt := 0
		for _, letter := range str {
			switch letter {
			case 'a', 'e', 'i', 'o', 'u', 'y', 'A', 'E', 'I', 'O', 'U', 'Y', 'ą', 'ę', 'ó', 'Ą', 'Ę', 'Ó':
				cnt++
			}
		}

		if cnt%2 == 0 {
			even = append(even, str)
		} else {
			odd = append(odd, str)
		}
	}

	return even, odd
}

func main() {
	var filePtr = flag.String("file", "lorem.txt", "a string")
	flag.Parse()

	content, errRead := ioutil.ReadFile(*filePtr)
	if errRead != nil {
		log.Fatal(errRead)
	}

	evenSlice, oddSlice := words(strings.NewReader(string(content)))

	errEven := ioutil.WriteFile("even.txt", []byte(strings.Join(evenSlice, " ")), 0644)
	if errEven != nil {
		log.Fatal(errEven)
	}

	errOdd := ioutil.WriteFile("odd.txt", []byte(strings.Join(oddSlice, " ")), 0644)
	if errOdd != nil {
		log.Fatal(errOdd)
	}

	fmt.Println(evenSlice, oddSlice)
}
