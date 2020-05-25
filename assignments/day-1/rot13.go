package main

import ( "fmt"; "os"; "strings" )

const alphabet_lo string = "abcdefghijklmnopqrstuvwxzy"
const alphabet_up string = "ABCDEFGHILKJMNOPQRSTUVWXZY"
const alphabet_len int = len(alphabet_lo)

func encode(str string) string {
	ret_str := ""

	var alphabet_i int

	for i := 0; i < len(str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') {
			alphabet_i = strings.IndexRune(alphabet_lo, rune(str[i]))
			alphabet_i = (alphabet_i + 13) % alphabet_len
			ret_str += string(alphabet_lo[alphabet_i])
		} else if (str[i] >= 'A' && str[i] <= 'Z') {
			alphabet_i = strings.IndexRune(alphabet_up, rune(str[i]))
			alphabet_i = (alphabet_i + 13) % alphabet_len
			ret_str += string(alphabet_up[alphabet_i])
		} else { // Invalid character
			fmt.Println("ERROR: Invalid character in input string")
			fmt.Println("Only [a-zA-Z] are allowed.")
			os.Exit(1)
		}
	}

	return ret_str
}

func decode (str string) string {
	ret_str := ""

	var alphabet_i int

	for i := 0; i < len(str); i++ {
		if (str[i] >= 'a' && str[i] <= 'z') {
			alphabet_i = strings.IndexRune(alphabet_lo, rune(str[i]))
			alphabet_i = (alphabet_i - 13) % alphabet_len
			if alphabet_i < 0 {
				alphabet_i += alphabet_len
			}
			ret_str += string(alphabet_lo[alphabet_i])
		} else if (str[i] >= 'A' && str[i] <= 'Z') {
			alphabet_i = strings.IndexRune(alphabet_up, rune(str[i]))
			alphabet_i = (alphabet_i - 13) % alphabet_len
			if alphabet_i < 0 {
				alphabet_i += alphabet_len
			}
			ret_str += string(alphabet_up[alphabet_i])
		} else { // Invalid character
			fmt.Println("ERROR: Invalid character in input string")
			fmt.Println("Only [a-zA-Z] are allowed.")
			os.Exit(1)
		}
	}

	return ret_str
}

func main() {
	// Check that the user provided the correct number of args
	if len(os.Args) != 3 {
		fmt.Println("ERROR: Invalid number of arguments.")
		fmt.Println("Usage: rot13 (encode|decode) STRING")
		os.Exit(1)
	}

	// Check that the user provided the correct args; encode or decode
	switch os.Args[1] {
	case "encode":
		fmt.Println(encode(os.Args[2]))
	case "decode":
		fmt.Println(decode(os.Args[2]))
	default:
		fmt.Println("ERROR: Invalid argument")
		fmt.Println("Usage: rot13 (encode|decode) STRING")
		os.Exit(1)
	}
}
