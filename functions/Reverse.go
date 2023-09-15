package functions

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

var ReversedDict map[string]byte
var Space string

func Reverse(filename string) error {
	var err error
	if filename != "" && !strings.HasSuffix(filename, ".txt") {
		return errors.New("wrong format of the file")
	}

	ReversedDict, err = getReversedASCII("standard")
	if err != nil {
		return err
	}

	fileSample, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	fmt.Print(convertFile(fileSample))
	return nil
}

func convertFile(fileSample []byte) string {
	arrSample := strings.Split(string(fileSample), "\n")
	var arrSentence [][]string
	var result []string

	for len(arrSample) > 0 {
		if len(arrSample)-8 > 0 && hasSameLength(arrSample[:8]) {

			arrSentence = append(arrSentence, arrSample[:8])
			arrSample = arrSample[8:]
		} else {
			arrSentence = append(arrSentence, arrSample[:1])
			arrSample = arrSample[1:]
		}
	}

	for _, sentence := range arrSentence {
		var letters []string
		if len(sentence) == 8 && len(sentence[0]) == 0 {
			letters = append(letters, "\n\n\n\n\n\n\n")
		}
		for line := 0; line < len(sentence); line++ {
			if len(sentence) == 1 {
				letters = append(letters, "\n")
				continue
			}

			for i, ch := range sentence[line] {

				if ch == 32 {
					if isDivider(sentence, i) {
						var letter []string
						SkipSpace := false
						if i < len(sentence[line])-5 && isSpace(sentence, i) {
							SkipSpace = true
						}
						sentence, letter = takeLetter(sentence, i)
						letters = append(letters, strings.Join(letter, "\n"))
						line = 0
						if SkipSpace {
							sentence, _ = takeLetter(sentence, 5)
							letters = append(letters, Space)
							line = 0
							break
						}
						break
					}
				}

			}
			if line == 7 {
				letters = append(letters, "\n")
			}
		}

		for i := range letters {
			switch letters[i] {
			case "\n":
				result = append(result, "\n")
			case "\n\n":
				result = append(result, "\n\n")
			case "\n\n\n\n\n\n\n":
				result = append(result, "\n\n\n\n\n\n\n")
			default:
				result = append(result, string(ReversedDict[letters[i]]))
			}
		}
	}

	var res string
	for i := range result {
		res += result[i]
	}
	return res
}

func hasSameLength(s []string) bool {
	k := len(s[0])
	for i := range s[1:] {
		if len(s[1:][i]) != k {
			return false
		}
	}
	return true
}

func isDivider(sentence []string, n int) bool {
	for i := 0; i < 8; i++ {
		if sentence[i][n] != 32 {
			return false
		}
	}
	return true
}

func isSpace(sentence []string, n int) bool {
	var space []string
	for i := 0; i < 8; i++ {
		space = append(space, sentence[i][n:n+6])
	}
	return Space == strings.Join(space, "\n")
}

func takeLetter(sentence []string, n int) ([]string, []string) {
	var letter []string
	for i := 0; i < 8; i++ {
		letter = append(letter, sentence[i][:n+1])
		sentence[i] = sentence[i][n+1:]

	}
	return sentence, letter
}

func getReversedASCII(font string) (map[string]byte, error) {
	dict := make(map[string]byte)

	f, err := os.ReadFile("fonts/" + font + ".txt")
	if err != nil {
		return nil, err
	}

	content := strings.Split(strings.ReplaceAll(string(f[1:]), "\r", ""), "\n\n")
	Space = content[0]
	letter := byte(32)
	for i := range content {
		dict[content[i]] = letter
		letter++
	}
	return dict, nil
}
