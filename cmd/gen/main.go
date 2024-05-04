package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

// What this project will achieve:
// 1) create the folders corelating to the year and day. - done
// 2) create a "solve.go" & "solve_test.go". - done
// 3) grab the input from the advent of code site.
// 4) grab the problem specs for the day specified as well and covert it from html to markdown.
// 5) use template to create a basic struct for solving problems. - done somewhat (this would
//    require trial and error)
// 6) create a lib dir which implements queue, stack, indexed queue etc etc...
// 7) create a lib for parsers

type Config struct {
	Year string
	Day  string
}

func main() {
	var cfg Config
	flag.StringVar(&cfg.Year, "year", strconv.Itoa(time.Now().Year()), "The year of the event")
	flag.StringVar(&cfg.Year, "y", cfg.Year, "short cut for year command")
	flag.StringVar(&cfg.Day, "day", strconv.Itoa(time.Now().Day()), "The day of the event")
	flag.StringVar(&cfg.Day, "d", cfg.Day, "short cut for day command")
	flag.Parse()

	if len(cfg.Day) == 1 {
		cfg.Day = fmt.Sprintf("0%s", cfg.Day)
	}
	if len(cfg.Year) == 2 { //nolint: gomnd // not reused.
		cfg.Year = fmt.Sprintf("20%s", cfg.Year)
	}

	dirPath := filepath.Join(cfg.Year, fmt.Sprintf("day_%s", cfg.Day))

	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(dirPath, os.ModePerm)
		if err != nil {
			log.Fatal(err)
		}
	}
	solve := filepath.Join(dirPath, "solve.go")
	solveTest := filepath.Join(dirPath, "solve_test.go")
	solveFile, err := os.OpenFile(solve, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer solveFile.Close()
	solveTestFile, err := os.OpenFile(solveTest, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err) //nolint: gocritic // false positive.
	}
	defer solveTestFile.Close()

	ts, err := template.ParseFiles("cmd/gen/templates/solve.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = ts.ExecuteTemplate(solveFile, "solve", cfg)
	if err != nil {
		log.Fatal(err)
	}
	ts, err = template.ParseFiles("cmd/gen/templates/solve_test.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	err = ts.ExecuteTemplate(solveTestFile, "test", cfg)
	if err != nil {
		log.Fatal(err)
	}
}
