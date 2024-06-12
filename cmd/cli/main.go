package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
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
// 8) create a lib dir which implements queue, stack, indexed queue etc etc...
// 7) create a lib for parsers - done (but its a continuously evolving lib)

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

	isSubmit := flag.Bool("submit", false, "submit the answer")
	flag.BoolVar(isSubmit, "s", *isSubmit, "short cut for submit command")
	partNum := flag.Int("part", 1, "problem part")
	flag.IntVar(partNum, "p", *partNum, "short cut for part command")
	answer := flag.String("answer", "", "answer for the solution")
	flag.StringVar(answer, "a", *answer, "short cut for answer command")
	flag.Parse()

	cfg.path = os.Getenv("AOC_PATH")
	cfg.session = os.Getenv("AOC_SESSION")

	if cfg.session == "" || cfg.path == "" {
		log.Fatal("AOC_PATH or AOC_SESSION env variable cannot be empty")
	}

	if *partNum == 2 {
		*isSubmit = true
	}
	if len(cfg.Day) == 1 {
		cfg.Day = fmt.Sprintf("0%s", cfg.Day)
	}
	if len(cfg.Year) == 2 { //nolint: gomnd // number not reused.
		cfg.Year = fmt.Sprintf("20%s", cfg.Year)
	}

	if *isSubmit {
		if *answer == "" {
			*answer = flag.Args()[len(flag.Args())-1]
		}
		err := cfg.submitAnswer(*answer, *partNum)
		handleErr(err)
		fmt.Printf("Congratulations! You have answered the Part %d correctly!", *partNum) //nolint:forbidigo // needed.
		if *partNum == 2 {
			os.Exit(0) // exit if its part 2 as there is no need to update readme or base cases.
		}
	}

	dirPath := filepath.Join(cfg.path, cfg.Year, fmt.Sprintf("day_%s", cfg.Day))

	isInitial := false
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		handleErr(cfg.createTemplate(dirPath))
		isInitial = true
	}

	err := cfg.createBaseCases(dirPath)
	handleErr(err)

	inputFile := filepath.Join(dirPath, "input.txt")
	input, err := os.OpenFile(inputFile, os.O_CREATE|os.O_RDWR, 0644)
	handleErr(err)
	handleErr(cfg.getInput(input))

	readMEFile := filepath.Join(dirPath, "README.md")
	readME, err := os.OpenFile(readMEFile, os.O_CREATE|os.O_RDWR, 0644)
	handleErr(err)
	handleErr(cfg.getSpec(readME))
	msgFmt := "Successfully updated Files at:\n%s"
	if isInitial {
		msgFmt = "Successfully created:\n%s"
	}
	fmt.Printf(msgFmt, dirPath) //nolint:forbidigo // needed.
}

func (c *Config) submitAnswer(answer string, part int) error {
	if part < 1 || part > 2 {
		return errors.New("submitAnswer: incorrect part number")
	}
	form := url.Values{}
	form.Add("level", strconv.Itoa(part))
	form.Add("answer", answer)
	day := c.Day
	if day[0] == '0' {
		day = day[1:]
	}
	aocURL := fmt.Sprintf("https://adventofcode.com/%s/day/%s/answer", c.Year, day)
	resp, err := c.getResponse(http.MethodPost, aocURL, strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	buffer, err := praseHTML(resp.Body, "article")
	if err != nil {
		return err
	}
	// TODO: handle correct answer, wrong answer and cool down individually.
	if !strings.Contains(buffer.String(), "That's the right answer!") {
		return errors.New("the answer provided is incorrect")
	}
	return nil
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

	ts, err := template.ParseFiles("cmd/cli/templates/solve.gotmpl")
	if err != nil {
		return err
	}
	err = ts.ExecuteTemplate(solveFile, "solve", c)
	if err != nil {
		return err
	}

	ts, err = template.ParseFiles("cmd/cli/templates/solve_test.gotmpl")
	if err != nil {
		return err
	}
	err = ts.ExecuteTemplate(solveTestFile, "test", c)
	return err
}

func (c *Config) getInput(w io.Writer) error {
	day := c.Day
	if day[0] == '0' {
		day = day[1:]
	}
	aocURL := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", c.Year, day)
	resp, err := c.getResponse(http.MethodGet, aocURL, nil)
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

func (c *Config) createBaseCases(dir string) error {
	day := c.Day
	if day[0] == '0' {
		day = day[1:]
	}
	aocURL := fmt.Sprintf("https://adventofcode.com/%s/day/%s", c.Year, day)
	data, err := c.getHTML(aocURL, "pre", "code")
	if err != nil {
		return err
	}
	bases := strings.Split(data.String(), "\n\n")
	for idx, base := range bases[:len(bases)-1] { // the last item would be an empty string.
		baseFilePath := filepath.Join(dir, fmt.Sprintf("base_%d.txt", idx+1))
		baseFile, err := os.OpenFile(baseFilePath, os.O_CREATE|os.O_RDWR, 0644) //nolint:govet // false positive
		if err != nil {
			return err
		}
		_, err = baseFile.WriteString(base)
		if err != nil {
			return err
		}
		baseFile.Close()
	}
	return nil
}

func (c *Config) getSpec(w io.Writer) error {
	day := c.Day
	if day[0] == '0' {
		day = day[1:]
	}
	aocURL := fmt.Sprintf("https://adventofcode.com/%s/day/%s", c.Year, day)
	respData, err := c.getHTML(aocURL, "article")
	if err != nil {
		return err
	}
	converter := md.NewConverter(aocURL, true, nil)
	converter.Use(plugin.GitHubFlavored())
	buf, err := converter.ConvertReader(respData)
	if err != nil {
		return err
	}
	_, err = io.Copy(w, &buf)
	return err
}

func (c *Config) getResponse(method string, aocURL string, body io.Reader) (*http.Response, error) {
	ctx, done := context.WithTimeout(
		context.Background(),
		5*time.Second) //nolint: gomnd // not reused.
	defer done()
	req, err := http.NewRequestWithContext(ctx, method, aocURL, body)
	if err != nil {
		return nil, err
	}
	req.AddCookie(&http.Cookie{Name: "session", Value: c.session})
	if method == http.MethodPost {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *Config) getHTML(aocURL string, tags ...string) (*bytes.Buffer, error) {
	resp, err := c.getResponse(http.MethodGet, aocURL, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("getHTML: got status code %d", resp.StatusCode)
	}
	cType := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(cType, "text/html") {
		return nil, errors.New("getHTML: did not get text/html")
	}

	respData, err := praseHTML(resp.Body, tags...)
	return respData, err
}

func praseHTML(reader io.Reader, tag ...string) (*bytes.Buffer, error) {
	tokenizer := html.NewTokenizer(reader)
	output := new(bytes.Buffer)
	content := false
	startTagIdx, endTagIdx := 0, 0
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			if err := tokenizer.Err(); !errors.Is(err, io.EOF) {
				return nil, err
			}
			break
		}
		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == tag[startTagIdx] {
			if startTagIdx == len(tag)-1 {
				content = true
			} else {
				startTagIdx++
			}

			if startTagIdx != 0 && endTagIdx != len(tag)-1 {
				endTagIdx++
			}
			continue
		}
		if tokenType == html.EndTagToken && token.Data == tag[endTagIdx] {
			if startTagIdx == len(tag)-1 {
				output.WriteString("\n\n")
			}
			content = false
			if startTagIdx > 0 {
				startTagIdx--
			}
			if endTagIdx > 0 {
				endTagIdx--
			}
		}
		if !content {
			continue
		}
		output.WriteString(html.UnescapeString(token.String()))
	}
	return output, nil
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
