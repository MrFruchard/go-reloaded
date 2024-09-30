package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	text, err := readFileToString(inputFile)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Applique les transforamtions
	text = applyTransformations(text)

	// phrase final dans le result
	err = writeStringToFile(outputFile, text)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	}

	fmt.Println(text)
	fmt.Println("Transformations applied and output saved to:", outputFile)
}

// Permet d'extraire la data
func readFileToString(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// Ecrit dans le result.txt
func writeStringToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

// Applique les transformations
func applyTransformations(text string) string {
	text = replaceHexBin(text)

	text = applyCaseTransformations(text)

	text = replaceAToAn(text)

	text = correctPunctuation(text)

	return text
}

// (hex), (bin))
func replaceHexBin(text string) string {
	reHex := regexp.MustCompile(`\b([0-9A-Fa-f]+) \(hex\)`)
	reBin := regexp.MustCompile(`\b([01]+) \(bin\)`)

	text = reHex.ReplaceAllStringFunc(text, func(m string) string {
		hexNumber := reHex.FindStringSubmatch(m)[1]
		decimal, _ := strconv.ParseInt(hexNumber, 16, 64)
		return strconv.FormatInt(decimal, 10)
	})

	text = reBin.ReplaceAllStringFunc(text, func(m string) string {
		binNumber := reBin.FindStringSubmatch(m)[1]
		decimal, _ := strconv.ParseInt(binNumber, 2, 64)
		return strconv.FormatInt(decimal, 10)
	})

	return text
}

// (up, 2), (low, 3), (cap)
func applyCaseTransformations(text string) string {
	// patern regex (a bosser)
	reUp := regexp.MustCompile(`\b(\w+(?: \w+){0,2}) \(up(?:, (\d+))?\)`)
	reLow := regexp.MustCompile(`\b(\w+(?: \w+){0,2}) \(low(?:, (\d+))?\)`)
	reCap := regexp.MustCompile(`\b(\w+(?: \w+){0,2}) \(cap(?:, (\d+))?\)`)

	// Up transformation
	text = reUp.ReplaceAllStringFunc(text, func(m string) string {
		match := reUp.FindStringSubmatch(m)
		count := 1
		if match[2] != "" {
			count, _ = strconv.Atoi(match[2])
		}
		words := strings.Fields(match[1])
		for i := len(words) - 1; i >= 0 && count > 0; i-- {
			words[i] = strings.ToUpper(words[i])
			count--
		}
		return strings.Join(words, " ")
	})

	// Low transformation
	text = reLow.ReplaceAllStringFunc(text, func(m string) string {
		match := reLow.FindStringSubmatch(m)
		count := 1
		if match[2] != "" {
			count, _ = strconv.Atoi(match[2])
		}
		words := strings.Fields(match[1])
		for i := len(words) - 1; i >= 0 && count > 0; i-- {
			words[i] = strings.ToLower(words[i])
			count--
		}
		return strings.Join(words, " ")
	})

	// Cap transformation
	text = reCap.ReplaceAllStringFunc(text, func(m string) string {
		match := reCap.FindStringSubmatch(m)
		count := 1
		if match[2] != "" {
			count, _ = strconv.Atoi(match[2])
		}
		words := strings.Fields(match[1])
		for i := len(words) - 1; i >= 0 && count > 0; i-- {
			words[i] = strings.Title(words[i])
			count--
		}
		return strings.Join(words, " ")
	})

	// Enleve les balises
	text = regexp.MustCompile(`\(\w+(?:, \d+)?\)`).ReplaceAllString(text, "")

	return text
}

// Remplace a par an
func replaceAToAn(text string) string {
	re := regexp.MustCompile(`\ba ([aeiouhAEIOUH])`)
	return re.ReplaceAllString(text, "an $1")
}

// Correction de ponctuation / espaces
func correctPunctuation(text string) string {
	// Permet de supprimez les espaces apres les ponctuations
	rePunct := regexp.MustCompile(`\s*([.,!?;:])\s*`)
	text = rePunct.ReplaceAllString(text, "$1 ")

	// Permet de supprimez les espaces autur des apostrophe
	reSingleQuote := regexp.MustCompile(`\s*'([^']+?)'\s*`)
	text = reSingleQuote.ReplaceAllString(text, " '$1' ")

	// permet de supprimez les espace inutile
	text = strings.ReplaceAll(text, " '", "'")
	text = strings.ReplaceAll(text, "' ", "'")

	// permet un espace apres :
	reColon := regexp.MustCompile(`:\s*`)
	text = reColon.ReplaceAllString(text, ": ")

	return strings.TrimSpace(text)
}
