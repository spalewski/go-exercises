package utilities

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

func Words(r io.Reader, path string) (even []string, odd []string) {

	var words []string
	even = []string{}
	odd = []string{}

	fileReaded, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Print(err)
	}
	strFromFile := string(fileReaded)

	//buf := new(bytes.Buffer)
	//buf.ReadFrom(r)
	//s := buf.String()

	scanner := bufio.NewScanner(strings.NewReader(strFromFile))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	for _, each := range words {

		if IsEven(CountVovels(each)) == true {
			even = append(even, each)
		} else {
			odd = append(odd, each)
		}
	}
	WriteEven(even)
	WriteOdd(odd)
	return even, odd
}
