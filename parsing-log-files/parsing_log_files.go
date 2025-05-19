package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	expr := `(^[TRC]|^\[DBG\]|^\[INF\]|^\[WRN\]|^\[ERR\]|^\[FTL\])`

	re := regexp.MustCompile(expr)

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	expr := "<([~*=\\-]||)+>"
	re := regexp.MustCompile(expr)
	return re.Split(text, -1)

}

func CountQuotedPasswords(lines []string) int {
	expr := `(?i)".*password.*"`
	re := regexp.MustCompile(expr)
	ctr := 0
	for _, line := range lines {
		val := re.FindAllString(line, -1)
		ctr += len(val)
	}
	return ctr
}

func RemoveEndOfLineText(text string) string {
	expr := `end-of-line\d*`
	re := regexp.MustCompile(expr)
	return re.ReplaceAllString(text, "")

}

func TagWithUserName(lines []string) []string {
	expr := `User\s+(\w+)`
	re := regexp.MustCompile(expr)
	for i, line := range lines {
		founds := re.FindStringSubmatch(line)
		if founds != nil {
			line = fmt.Sprintf("[USR] %s %s", founds[1], line)
		}
		lines[i] = line
	}
	return lines
}
