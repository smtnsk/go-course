package main

import ( "fmt"; "flag" )

func main() {
        var name string
        var n int
        
        flag.StringVar(&name, "name", "", "Who to greet")
        flag.IntVar(&n, "n", 0, "How many greetings")
        flag.Parse()

        if name == "" {
                fmt.Println("Kerro nimesi, lurjus! >:(")
        } else if n == 0 {
                fmt.Println("Anna myös '--n' parametri.")
        } else {
                for i := 1; i <= n; i++ {
                        fmt.Printf("Howdy, %s! (%d)\n", name, i)
                }
        }
}
