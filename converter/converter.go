package converter

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

var imageRegex = regexp.MustCompile(`!\[(.*?)]\(.*?\)`)
var linkRegex = regexp.MustCompile(`\[(.*?)]\(.*?\)`)
var boldRegex = regexp.MustCompile(`\*\*(.*?)\*\*`)
var italicRegex = regexp.MustCompile(`\*(.*?)\*`)
var strikethroughRegex = regexp.MustCompile(`~~(.*?)~~`)
var inlineCodeRegex = regexp.MustCompile("`(.*?)`")
var horizontalRuleRegex = regexp.MustCompile(`^[-*_]{3,}\s*$`)
var orderedListRegex = regexp.MustCompile(`^\d+\.\s(.*)`)
var tableSeparatorRegex = regexp.MustCompile(`^[|\s\-:]+$`)
var tableCellRegex = regexp.MustCompile(`\|`)

func ConvertMdToTxt(file *os.File) string {
	scanner := bufio.NewScanner(file)
	var sb strings.Builder
	insideCodeBlock := false

	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "```") {
			insideCodeBlock = !insideCodeBlock
			continue // skip the fence line itself
		} else if insideCodeBlock {
			sb.WriteString(line)
			sb.WriteString("\n")
			continue
		} else if strings.HasPrefix(line, "#") || strings.HasPrefix(line, ">") ||
			strings.HasPrefix(line, "+") || strings.HasPrefix(line, "-") ||
			strings.HasPrefix(line, "* ") {
			idx := strings.Index(line, " ")
			if idx != -1 {
				processLine(line[idx+1:], &sb)
			}
		} else if horizontalRuleRegex.MatchString(line) {
			sb.WriteString("\n")
		} else if orderedListRegex.MatchString(line) {
			processLine(orderedListRegex.ReplaceAllString(line, "$1"), &sb)
		} else if strings.HasPrefix(line, "|") {
			if tableSeparatorRegex.MatchString(line) {
				// skip separator line like |---|---|
			} else {
				processed := tableCellRegex.ReplaceAllString(line, " ")
				processLine(processed, &sb)
			}
		} else {
			processLine(line, &sb)
		}

	}

	return sb.String()
}

func removeInlineMarkings(line string) string {
	output := removeImages(line)
	output = removeLinks(output)
	output = removeBold(output)
	output = removeItalic(output)
	output = removeStrikethrough(output)
	output = removeInlineCode(output)
	return output
}

func removeImages(line string) string {
	return imageRegex.ReplaceAllString(line, "$1")
}

func removeLinks(line string) string {
	return linkRegex.ReplaceAllString(line, "$1")
}

func removeBold(line string) string {
	return boldRegex.ReplaceAllString(line, "$1")
}

func removeItalic(line string) string {
	return italicRegex.ReplaceAllString(line, "$1")
}

func removeStrikethrough(line string) string {
	return strikethroughRegex.ReplaceAllString(line, "$1")
}

func removeInlineCode(line string) string {
	return inlineCodeRegex.ReplaceAllString(line, "$1")
}

func processLine(line string, sb *strings.Builder) {
	temp := strings.TrimSpace(line)
	sb.WriteString(removeInlineMarkings(temp))
	sb.WriteString("\n")
}
