package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Initializes New Reader
	cin := bufio.NewReader(os.Stdin)

	// Give Rules/Intro

	// Main Menu Loop
	userResponse := 0
	for userResponse != 3 {
		// Menu
		fmt.Println("Please choose a game length.")
		fmt.Println("(choose \"1\" if you just want to try it out)")
		fmt.Println()
		fmt.Println("1. Fast Game (play to 36 cr.)")
		fmt.Println("2. Half Game (play to 72 cr.)")
		fmt.Println("3. Full Game (play to 144 cr.)")
		fmt.Println("4. Quit")

		// Get user choice
		// If it is in [0, 2] a game will start, if it is 3, the game will exit.
		userResponse = getUserInt(cin, 1, 4) - 1
		switch userResponse {
		case 0:
			playGame(36, cin)
		case 1:
			playGame(72, cin)
		case 2:
			playGame(144, cin)
		}
	}

	fmt.Println("Thank you for Playing \"Nothing to Pebbles\"")
}

func playGame(goal int, cin *bufio.Reader) {
	fmt.Println("Please enter a seed for the Random Number Generator.")
	fmt.Println("You can also just press \"Enter\" without providing any input in order to play with a random seed.")

	// Get the user-end seed
	origSeed := getUserReply(cin)
	if origSeed == "" {
		// Gen random seed
		origSeed = strconv.FormatInt(time.Now().UnixNano(), 16)
	}

	// Hash origSeed and set result to hashedSeed
	hash := fnv.New64a()
	hash.Write([]byte(origSeed))
	hashedSeedU := hash.Sum64()
	var hashedSeed int64 = int64(hashedSeedU)

	// Create New RNG
	randGen := rand.New(rand.NewSource(hashedSeed))

	credits := 0

	// Main Game Loop
	for credits >= 0 && credits < goal {
		// Create array of cups
		var cups [3]int
		cupi := 0
		for cupi < 3 {
			cups[cupi] = randGen.Intn(10)
			cupi++
		}

		// Show Pebbles
		fmt.Println(cups)

		// Proceed
		cin.ReadString('\n')
		clearScreen()

		// Start Timer

		// Show a transposition.

		cin.ReadString('\n')
		clearScreen()

		// delay := timer.stop()
	}
	// Post-game

	// Show the user their seed if they want to play the same level again.
	fmt.Printf("Your seed: \"%v\"\n", origSeed)
}

// Gets an int from the user
func getUserInt(cin *bufio.Reader, min int, max int) int {
	userIntA := ""
	userInt := 0

	validUserChoice := false
	for validUserChoice == false {
		userIntA, _ = cin.ReadString('\n')
		userIntA = strings.Replace(userIntA, "\n", "", -1)

		// Convert to an int
		userIntL, err := strconv.Atoi(userIntA)
		if err != nil {
			fmt.Println(err) // Debugging

			fmt.Println("Invalid Input")
		} else {
			userInt = userIntL
			validUserChoice = true // Exit Loop
		}

		// If it passed the above test, ensure the number is within bounds
		if validUserChoice == true {
			if userInt < min {
				fmt.Println("That number is too low.")
				validUserChoice = false // Do not Exit Loop
			} else if userInt > max {

				fmt.Println("That number is too high.") // Do not Exit Loop
				validUserChoice = false
			}

		}

	} // User has entered a valid choice

	return userInt
}

// Gets a string from the user
func getUserReply(cin *bufio.Reader) string {

	userReply, _ := cin.ReadString('\n')
	userReply = strings.Replace(userReply, "\n", "", -1)

	return userReply
}

// There are better ways to do this, so it's in it's own function for easy bebetterment later.
func clearScreen() {
	for lines := 0; lines < 0x80; lines++ {
		fmt.Println()
	}
}
