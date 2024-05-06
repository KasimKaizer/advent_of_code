package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
	"time"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/JohannesKaufmann/html-to-markdown/plugin"
	"golang.org/x/net/html"
)

// What this project will achieve:
// 1) create the folders corelating to the year and day. - done
// 2) create a "solve.go" & "solve_test.go". - done (add benchmark funcs)
// 3) grab the input from the advent of code site. - done
// 4) grab the problem specs for the day specified as well as covert it from html to markdown. -
//    done
// 5) use template to create a basic struct for solving problems. - done somewhat (this would
//    require trial and error)
// 6) create a lib dir which implements queue, stack, indexed queue etc etc...
// 7) create a lib for parsers

type Config struct {
	Year    string
	Day     string
	path    string
	session string
}

func main() {
	cfg := new(Config)
	flag.StringVar(&cfg.Year, "year", strconv.Itoa(time.Now().Year()), "The year of the event")
	flag.StringVar(&cfg.Year, "y", cfg.Year, "short cut for year command")
	flag.StringVar(&cfg.Day, "day", strconv.Itoa(time.Now().Day()), "The day of the event")
	flag.StringVar(&cfg.Day, "d", cfg.Day, "short cut for day command")
	flag.Parse()
	cfg.path = os.Getenv("AOC_PATH")
	cfg.session = os.Getenv("AOC_SESSION")

	if len(cfg.Day) == 1 {
		cfg.Day = fmt.Sprintf("0%s", cfg.Day)
	}
	if len(cfg.Year) == 2 { //nolint: gomnd // not reused.
		cfg.Year = fmt.Sprintf("20%s", cfg.Year)
	}

	dirPath := filepath.Join(cfg.path, cfg.Year, fmt.Sprintf("day_%s", cfg.Day))
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		handleErr(cfg.createTemplate(dirPath))
	}

	inputFile := filepath.Join(dirPath, "input.txt")
	input, err := os.OpenFile(inputFile, os.O_CREATE|os.O_RDWR, 0644)
	handleErr(err)
	handleErr(cfg.getInput(input))

	readMEFile := filepath.Join(dirPath, "README.md")
	readME, err := os.OpenFile(readMEFile, os.O_CREATE|os.O_RDWR, 0644)
	handleErr(err)
	handleErr(cfg.getSpec(readME))
}

func (c *Config) createTemplate(dirPath string) error {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return err
	}
	solve := filepath.Join(dirPath, "solve.go")
	solveFile, err := os.OpenFile(solve, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer solveFile.Close()

	solveTest := filepath.Join(dirPath, "solve_test.go")
	solveTestFile, err := os.OpenFile(solveTest, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer solveTestFile.Close()

	ts, err := template.ParseFiles("/Users/kaizersuterwala/go/src/advent_of_code/cmd/gen/templates/solve.tmpl")
	if err != nil {
		return err
	}
	err = ts.ExecuteTemplate(solveFile, "solve", c)
	if err != nil {
		return err
	}

	ts, err = template.ParseFiles("/Users/kaizersuterwala/go/src/advent_of_code/cmd/gen/templates/solve_test.tmpl")
	if err != nil {
		return err
	}
	err = ts.ExecuteTemplate(solveTestFile, "test", c)
	if err != nil {
		return err
	}

	return nil
}

func (c *Config) getInput(w io.Writer) error {
	day := c.Day
	if day[0] == '0' {
		day = day[1:]
	}
	resp, err := c.getResponse(fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", c.Year, day))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getInput: got status code %d", resp.StatusCode)
	}
	_, err = io.Copy(w, resp.Body)
	return err
}

func (c *Config) getSpec(w io.Writer) error {
	day := c.Day
	if day[0] == '0' {
		day = day[1:]
	}
	resp, err := c.getResponse(fmt.Sprintf("https://adventofcode.com/%s/day/%s", c.Year, day))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("getSpec: got status code %d", resp.StatusCode)
	}
	cType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(cType, "text/html") {
		return errors.New("getSpec: did not get text/html")
	}
	tokenizer := html.NewTokenizer(resp.Body)
	output := new(bytes.Buffer)
	isArticle := false
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			if err = tokenizer.Err(); !errors.Is(err, io.EOF) {
				return err
			}
			break
		}
		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == "article" {
			isArticle = true
			continue
		}
		if tokenType == html.EndTagToken && token.Data == "article" {
			isArticle = false
		}
		if !isArticle {
			continue
		}
		output.WriteString(html.UnescapeString(token.String()))
	}
	converter := md.NewConverter("", true, &md.Options{EscapeMode: "basic"}) //nolint:exhaustruct // not needed
	converter.Use(plugin.GitHubFlavored())
	buf, err := converter.ConvertReader(output)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, &buf)
	return err
}

func (c *Config) getResponse(aocURL string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, aocURL, nil)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: c.session})
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
