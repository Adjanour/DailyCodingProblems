package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

var day = flag.String("day", "", "Day number (e.g. 03)")

func main() {
	flag.Parse()

	if *day == "" {
		fmt.Println("Usage: go run scaffold.go --day=XX")
		os.Exit(1)
	}

	dir := fmt.Sprintf("day%s", *day)
	if err := os.MkdirAll(dir, 0755); err != nil {
		fmt.Println("Error creating directory:", err)
		os.Exit(1)
	}

	files := map[string]string{
		"run.go":           runTemplate,
		"solution.go":      solutionTemplate,
		"solution_test.go": testTemplate,
		"thoughts.md":      thinkingTemplate,
	}

	for name, content := range files {
		path := filepath.Join(dir, name)
		if err := writeTemplate(path, content, *day); err != nil {
			fmt.Println("Error writing file:", err)
			os.Exit(1)
		}
	}

	if err := updateMainGo(*day); err != nil {
		fmt.Println("Error updating main.go:", err)
		os.Exit(1)
	}

	fmt.Println("Scaffold created and CLI updated for day", *day)
}

func writeTemplate(path, tmpl, day string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	t := template.Must(template.New("").Parse(tmpl))
	return t.Execute(f, struct{ Day string }{Day: day})
}

func updateMainGo(day string) error {
	mainPath := "main.go"
	content, err := os.ReadFile(mainPath)
	if err != nil {
		return err
	}

	importLine := fmt.Sprintf(` "github.com/adjanour/DailyCodingProblems/day%s"`, day)
	registryLine := fmt.Sprintf(`   "%s": day%s.Run,`, day, day)

	lines := strings.Split(string(content), "\n")
	var newLines []string
	importInserted := false
	registryInserted := false

	for _, line := range lines {
		newLines = append(newLines, line)

		if strings.HasPrefix(line, "import (") && !importInserted {
			newLines = append(newLines, importLine)
			importInserted = true
		}

		if strings.Contains(line, "var registry = map[string]func()") && !registryInserted {
			newLines = append(newLines, registryLine)
			registryInserted = true
		}
	}

	return os.WriteFile(mainPath, []byte(strings.Join(newLines, "\n")), 0644)
}

const runTemplate = `package day{{.Day}}

import "fmt"

func Run() {
    fmt.Println("Running Day {{.Day}} solution...")
    // Call your solution here
}
`

const solutionTemplate = `package day{{.Day}}

// Your solution logic goes here
`
const thinkingTemplate = ` # My thought process {{.Day}}
// thoughts go here
`

const testTemplate = `package day{{.Day}}

import "testing"

func TestSolution(t *testing.T) {
    // Write your test cases here
}
`
