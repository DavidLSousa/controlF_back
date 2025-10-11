package utils

import (
	"encoding/csv"
	"log"
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

var fqdnRegex *regexp.Regexp = regexp.MustCompile(`^[a-zA-Z-0-9]+[a-zA-Z-0-9\.]*[a-zA-Z-0-9]$`)
var letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("0123456789")

func GenString(n uint8) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func GenNumbers(n uint8) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = numbers[rand.Intn(len(numbers))]
	}
	return string(b)
}

func ParseUrlFriendly(value string) string {
	re := regexp.MustCompile(`[^a-z0-9._-]`)
	clean := strings.Map(func(r rune) rune {
		if unicode.IsPrint(r) {
			return r
		}
		return -1
	}, strings.ToLower(value))
	clean = re.ReplaceAllString(clean, "-")
	return dedupBy(clean, '-')
}

func dedupBy(value string, x rune) string {
	var last rune
	var sb strings.Builder
	for _, r := range value {
		if r != last || r != x {
			sb.WriteRune(r)
			last = r
		} else {
			continue
		}
	}
	return sb.String()
}

func SplitCommand(command string) []string {
	r := csv.NewReader(strings.NewReader(command))
	r.Comma = ' '
	record, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}

	return record
}

func IsFQDN(input string) bool {
	return fqdnRegex.MatchString(input)
}
