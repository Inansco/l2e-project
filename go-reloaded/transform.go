package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

// Main transformer that calls all individual transformations
func Transform(text string) string {
	text = FixHex(text)
	text = FixBin(text)
	text = FixCase(text)
	text = FixPunctuation(text)
	text = FixQuotes(text)
	text = FixAtoAn(text)
	return text
}

// 1. Converting hex to decimal
func FixHex(text string) string {
	re := regexp.MustCompile(`(\S+)\s+\(hex\)`)
	return re.ReplaceAllStringFunc(text, func(m string) string {
		hex := strings.Fields(m)[0]
		num, _ := strconv.ParseInt(hex, 16, 64)
		return strconv.FormatInt(num, 10)
	})
}

// 2. Convert binary to decimal
func FixBin(text string) string {
	re := regexp.MustCompile(`(\S+)\s+\(bin\)`)
	return re.ReplaceAllStringFunc(text, func(m string) string {
		bin := strings.Fields(m)[0]
		num, _ := strconv.ParseInt(bin, 2, 64)
		return strconv.FormatInt(num, 10)
	})
}

// 3. Handle case conversions (up, low, cap)
func FixCase(text string) string {
	// Handle (up)
	re := regexp.MustCompile(`(\S+(?:\s+\S+)*?)\s+\(up(?:,\s*(\d+))?\)`)
	text = re.ReplaceAllStringFunc(text, func(m string) string {
		return CaseHelper(m, "up")
	})

	// Handle (low)
	re = regexp.MustCompile(`(\S+(?:\s+\S+)*?)\s+\(low(?:,\s*(\d+))?\)`)
	text = re.ReplaceAllStringFunc(text, func(m string) string {
		return CaseHelper(m, "low")
	})

	// Handle (cap)
	re = regexp.MustCompile(`(\S+(?:\s+\S+)*?)\s+\(cap(?:,\s*(\d+))?\)`)
	text = re.ReplaceAllStringFunc(text, func(m string) string {
		return CaseHelper(m, "cap")
	})

	return text
}

// Helper for case conversions
func CaseHelper(match string, caseType string) string {
	numRe := regexp.MustCompile(`\(` + caseType + `(?:,\s*(\d+))?\)`)
	numMatch := numRe.FindStringSubmatch(match)
	count := 1
	if len(numMatch) > 1 && numMatch[1] != "" {
		count, _ = strconv.Atoi(numMatch[1])
	}

	idx := strings.LastIndex(match, "(")
	words := strings.Fields(match[:idx])

	start := len(words) - count
	if start < 0 {
		start = 0
	}

	for i := start; i < len(words); i++ {
		switch caseType {
		case "up":
			words[i] = strings.ToUpper(words[i])
		case "low":
			words[i] = strings.ToLower(words[i])
		case "cap":
			words[i] = Capitalize(words[i])
		}
	}

	return strings.Join(words, " ")
}
func Capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

// Fix punctuation spacing
func FixPunctuation(text string) string {
	text = regexp.MustCompile(`\s+([.!?:;,])`).ReplaceAllString(text, "$1")
	text = regexp.MustCompile(`([.!?:;,])([^\s])`).ReplaceAllString(text, "$1 $2")
	return text
}

// Fix quotes
func FixQuotes(text string) string {
	re := regexp.MustCompile(`'\s+([^']+?)\s+'`)
	return re.ReplaceAllString(text, "'$1'")
}

// Change "a" to "an"
func FixAtoAn(text string) string {
	re := regexp.MustCompile(`\ba\s+([aeiouAEIOU]|[hH])`)
	return re.ReplaceAllString(text, "an $1")
}
