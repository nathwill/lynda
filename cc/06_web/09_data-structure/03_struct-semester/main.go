package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

type course struct {
	Number, Name, Units string
}

type semester struct {
	Term		string
	Courses 	[]course
}

func main() {

	records := getRecords("data/first_semester.txt")
	xc := make([]course, 0, len(records))
	s := semester{}
	// #3 display output
	for i, row := range records {
		if i == 0 {
			// get header info
			xs := strings.SplitN(row[0], ",", 2)
			t := xs[0]
			s.Term = t
		} else {
			// get data
			c := course{}
			xs := strings.SplitN(row[0], " ", 2)
			c.Number = xs[0]
			c.Name = xs[1]
			c.Units = row[1]
			xc = append(xc, c)
		}
	}
	s.Courses = xc
	fmt.Println(s)
}

func getRecords(path string) [][]string {
	// #1 open a file
	f, err := os.Open(path)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// #2 parse a csv file
	rdr := csv.NewReader(f)
	rows, err := rdr.ReadAll()
	if err != nil {
		panic(err)
	}
	return rows
}