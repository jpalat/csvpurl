package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	cmdargs := os.Args[1:]
	file := cmdargs[0]
	line, err := strconv.Atoi(cmdargs[1])
	csvFile, _ := os.Open(file)
	reader := csv.NewReader(csvFile)
	header, err := reader.Read()

	if err != nil {
		log.Fatal(err)
	}

	count := 0

	for {

		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if count == line {
			for index, item := range header {
				fmt.Println(item, ":", record[index])
			}
		}
		count = count + 1
	}

	// fmt.Println(header)

}
