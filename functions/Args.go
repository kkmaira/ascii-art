package functions

import (
	"errors"
	"strings"
)

// GetArgs function receives os.Args and returns:
//
//  1. banner - the name of the font
//  2. input - the input string to convert
//  3. output - the name of the file for the output result of the ASCII ART PROGRAM
//  4. align - the align option flag
//  5. color - slice of colors and their corresponding values of target letters
func GetArgs(args []string) (banner string, input string, output string, align string, reverse string, err error) {
	args = args[1:]
	k := 0
	for i := 0; i < len(args); {
		k++
		switch {
		case strings.HasPrefix(args[i], "--output="):

			output = strings.Split(args[i], "=")[1]
			args = append(args[:i], args[i+1:]...)

		case strings.HasPrefix(args[i], "--align="):

			align = strings.Split(args[i], "=")[1]
			args = append(args[:i], args[i+1:]...)
			align = strings.ToLower(align)
		case strings.HasPrefix(args[i], "--reverse="):
			reverse = strings.Split(args[i], "=")[1]
			args = append(args[:i], args[i+1:]...)
		default:
			i++

		}
	}

	if reverse != "" {
		return "", "", "", "", reverse, nil
	}

	if len(args) == 1 {
		return "standard", args[0], output, align, reverse, nil
	}

	if len(args) == 2 && reverse == "" {
		var count int
		for i := range args {
			if strings.ToLower(args[i]) != "shadow" && strings.ToLower(args[i]) != "standard" && strings.ToLower(args[i]) != "thinkertoy" {
				input = args[i]
				count++
			} else {
				if banner != "" && i == len(args)-1 {

					input = banner
					banner = args[i]

					return banner, input, output, align, reverse, nil

				}

				banner = args[i]

			}
		}
		if count == len(args) {
			return banner, input, output, align, reverse, errors.New("ERROR")
		}

		return banner, input, output, align, reverse, nil
	}

	return "", "", "", "", "", errors.New("ERROR: something's wrong with the arguments...TRY THIS:\n\t\t    go run . <option> <input> <banner>")
}
