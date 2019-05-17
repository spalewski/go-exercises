package utilities

import (
	"bufio"
	"log"
	"os"
)

func WriteEven(s []string) {
	file, err := os.OpenFile("even.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	dataWriter := bufio.NewWriter(file)

	for _, data := range s {
		_, _ = dataWriter.WriteString(data + "\n")
	}

	dataWriter.Flush()
	file.Close()
}
func WriteOdd(s []string) {
	file, err := os.OpenFile("odd.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	dataWriter := bufio.NewWriter(file)

	for _, data := range s {
		_, _ = dataWriter.WriteString(data + "\n")
	}

	dataWriter.Flush()
	file.Close()
}
