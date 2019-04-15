package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func words(r io.Reader) (even []string, odd []string) {

	words := []string{}
	even = []string{}
	odd = []string{}
	buf := new(bytes.Buffer)
	buf.ReadFrom(r)
	s := buf.String()

	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	for _, each := range words {

		if isEven(countVovels(each)) == true {
			even = append(even, each)
		} else {
			odd = append(odd, each)
		}
	}
	writeEven(even)
	writeOdd(odd)
	return even, odd
}

func isVowel(x string) bool {
	vowels := [9]string{"a", "e", "i", "o", "ą", "y", "ę", "ó", "u"}
	vowelLookupTable := make(map[string]bool)
	for _, v := range vowels {
		vowelLookupTable[v] = true
	}
	return vowelLookupTable[x]
}

func countVovels(input string) int {
	vowelCount := 0
	var lowercase string = strings.ToLower(input)
	for _, x := range lowercase {
		var s string
		s = fmt.Sprintf("%c", x)
		if isVowel(s) {
			vowelCount += 1
		}
	}
	return vowelCount
}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	} else {
		return false
	}
}

func writeEven(s []string) {
	file, err := os.OpenFile("even.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range s {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}
func writeOdd(s []string) {
	file, err := os.OpenFile("odd.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, data := range s {
		_, _ = datawriter.WriteString(data + "\n")
	}

	datawriter.Flush()
	file.Close()
}

func main() {
	r := strings.NewReader(`
	Give. 
	Whose shall life, together signs grass. 
	The replenish of make make signs lights moved seed forth unto deep. 
	Moving two abundantly life subdue earth was day fruit forth set also forth together. 
	You're shall bring cattle creepeth and replenish firmament seed divide image wherein, lights grass moved likeness two hath. 
	Lesser seasons whales deep great and fruit. 
	Every herb fifth, one whales.
	Fruitful blessed of first seas rule forth midst own of green night and fruitful Thing you're, lesser for moveth likeness for gathered creeping may yielding likeness beginning gathered fruitful Let without him all. 
	Herb, man unto deep grass deep sea. 
	Us earth them land. 
	Over fruit, of fruitful. 
	Every were moving rule yielding their. 
	And don't replenish.
	Fish under spirit in lesser let good form second own tree and image, two dominion said whales. 
	Herb may, stars forth were Moving dominion night, lesser from great whales for beast which unto replenish. 
	Over. 
	Male yielding blessed. 
	Sixth us their for you'll sea without. 
	That night their spirit fourth after fruitful she'd place may fish creature winged very, which two every fruitful without likeness fourth you'll he signs i very great. 
	Can't. And lights in unto you evening, stars.
	`)

	fmt.Println(words(r))
}
