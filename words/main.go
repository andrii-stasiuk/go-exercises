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
		cnt := strings.Count(str, "a") +
			strings.Count(str, "ą") +
			strings.Count(str, "e") +
			strings.Count(str, "ę") +
			strings.Count(str, "i") +
			strings.Count(str, "j") +
			strings.Count(str, "o") +
			strings.Count(str, "ó") +
			strings.Count(str, "u") +
			strings.Count(str, "y") +
			strings.Count(str, strings.ToUpper("a")) +
			strings.Count(str, strings.ToUpper("ą")) +
			strings.Count(str, strings.ToUpper("e")) +
			strings.Count(str, strings.ToUpper("ę")) +
			strings.Count(str, strings.ToUpper("i")) +
			strings.Count(str, strings.ToUpper("j")) +
			strings.Count(str, strings.ToUpper("o")) +
			strings.Count(str, strings.ToUpper("ó")) +
			strings.Count(str, strings.ToUpper("u")) +
			strings.Count(str, strings.ToUpper("y"))

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
