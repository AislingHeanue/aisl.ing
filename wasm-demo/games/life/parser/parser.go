package parser

import (
	"embed"
	"fmt"
	"io"
	"math/rand"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

//go:embed patterns/*
var files embed.FS

var headerLineRe *regexp.Regexp = regexp.MustCompile("x = ([0-9]+), y = ([0-9]+)(?:, rule = (.*))?")
var dataLineRe *regexp.Regexp = regexp.MustCompile("([0-9]+)?([bo\\$])")

type ParsedStuff struct {
	x                 int
	y                 int
	paddedX           int
	paddedY           int
	commentLines      []string
	data              [][]bool
	calculatedPadding int
}

func ReadRandomFile() ([][]bool, string) {
	// FIXME: bring back the oversized folder from the collection, some of them are probably cool.
	// actually oversized files can be pruned later
	folder, err := files.ReadDir("patterns")
	if err != nil {
		fmt.Println(err)
		return nil, ""
	}

	fileNames := []string{}
	for _, subfolderEntry := range folder {
		if !subfolderEntry.IsDir() {
			continue
		}
		// if subfolderEntry.Name() != "synthesis" {
		// 	continue
		// }
		subfolder, err := files.ReadDir(filepath.Join("patterns", subfolderEntry.Name()))

		if err != nil {
			fmt.Println(err)
			return nil, ""
		}

		for _, pattern := range subfolder {
			name := filepath.Join(subfolderEntry.Name(), pattern.Name())
			fileNames = append(fileNames, name)
		}
	}
	choice := rand.Intn(len(fileNames))

	fmt.Println(fileNames[choice])
	return ReadFile(fileNames[choice]), fileNames[choice]
}

func ReadFile(path string) [][]bool {
	f, err := files.Open(filepath.Join("patterns", path))
	if err != nil {
		fmt.Println(err)
		return nil
	}

	content, err := io.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	lines := strings.Split(string(content), "\n")
	parsedStuff := ParsedStuff{}

	readLines(lines, &parsedStuff)

	return parsedStuff.data
}

func readLines(lines []string, parsedStuff *ParsedStuff) {
	dataLine := ""
	for i, line := range lines {
		if len(line) == 0 {
			continue
		}

		if line[0] == '#' {
			if len(line) < 2 || line[1] != 'C' {
				fmt.Println("Unexpected metadata line found: " + line)
			}
			parsedStuff.commentLines = append(parsedStuff.commentLines, strings.TrimPrefix(line, "#C"))
			continue
		}

		matches := headerLineRe.FindStringSubmatch(line)
		if matches[0] == "" {
			println("Error reading header line: " + line)
		}
		parsedStuff.x, _ = strconv.Atoi(matches[1])
		parsedStuff.y, _ = strconv.Atoi(matches[2])
		if matches[3] != "" && matches[3] != "B3/S23" {
			println("A new ruleset! How queer: " + matches[3])
		}
		if parsedStuff.x < 1 || parsedStuff.y < 1 {
			println("invalid dimensions: " + line)
		}
		parsedStuff.calculatedPadding = getPadding(parsedStuff.x, parsedStuff.y)
		parsedStuff.paddedX = parsedStuff.x + 2*parsedStuff.calculatedPadding
		parsedStuff.paddedY = parsedStuff.y + 2*parsedStuff.calculatedPadding

		// init the data 2d array
		parsedStuff.data = make([][]bool, parsedStuff.paddedY)
		for y := range parsedStuff.paddedY {
			parsedStuff.data[y] = make([]bool, parsedStuff.paddedX)
		}

		dataLine = strings.Split(strings.Join(lines[i+1:], ""), "!")[0]
		break
	}

	encodingEntries := dataLineRe.FindAllStringSubmatch(dataLine, -1)
	xPos := 0
	yPos := 0
	for _, entry := range encodingEntries {
		length := 1
		if entry[1] != "" {
			length, _ = strconv.Atoi(entry[1])
		}
		switch entry[2] {
		case "o":
			for range length {
				parsedStuff.data[yPos+parsedStuff.calculatedPadding][xPos+parsedStuff.calculatedPadding] = true
				xPos += 1
			}
		case "$":
			yPos += length
			xPos = 0
		case "b":
			xPos += length
		}
	}

}

// padding is 2 of the normal length on each side.
// so, the appropriate zoom level to ignore the padding is
// 1/5 = a zoom factor of 5
func getPadding(x, y int) int {
	padding := max(30, 2*float32(max(x, y)))
	return int(padding)
}
