package main 

import(
	"fmt"
	"os"
	"log"
	"strconv"
)

func main() {
	//first argument is the k, constant of rotation
	if len(os.Args) != 2 {
		fmt.Println("")
		log.Fatal("Please specify constant of rotation")
	}

	// need to parse args[1] to int
	k, err := strconv.Atoi(os.Args[1])
	if err != nil {	panic(err) }

	//get secret
	var secret string
	fmt.Println("What's the secret?")
	_, err = fmt.Scanf("%s", &secret)
	if err != nil {panic(err)}

	//runed is the slice of runes rep. of secret
	runed := []rune(secret)
	
	//iterate over runed, zero index, rotate, reindex, print
	for i, n := 0, len(runed); i < n; i++ {
		start := runed[i]

		if runed[i] > 64 && runed[i] < 91 {
			runed[i] = runed[i] - 'A'
		}
		if runed[i] > 96 && runed[i] < 122 {
			runed[i] = runed[i] - 'a'
		}

		moves := 0
		index := runed[i]

		for moves < (k % 26) {
			if index == 25 {
				index = 0
				moves++
			} else {
				index++
				moves++
			}
		}

		if start > 64 && start < 91 {
			runed[i] = index + 'A'
		}
		if start > 96 && start < 122 {
			runed[i] = index + 'a'
		}

		fmt.Printf("%c", runed[i])
	}
	fmt.Println()
}