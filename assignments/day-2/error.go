package main

import ( "fmt"; "errors" )

func four_letter_strings_only(str string) (bool, error) {
	if len(str) != 4 {
		return false, errors.New("ERROR: four letter stringz plz")
	}
	return true, nil
}

func main() {
	if bool, error := four_letter_strings_only("123"); bool {
		fmt.Println("thankz for the string <3")
	} else {
		fmt.Println(error)
	}
}
