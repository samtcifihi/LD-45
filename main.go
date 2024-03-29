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
	fmt.Println("Welcome to \"Nothing to Pebbles\"!")
	fmt.Println("You will be shown the numbers under each of three cups numbered 1, 2, and 3.")
	fmt.Println("Then, you will be shown pairs of number which indicate which two cups have been swapped with eachother. You will have to hit enter once you've accounted for the first swap, after that each swap will be shown only for the amount of time you took on the first one, so give yourself enough time.")
	fmt.Println("After that, you will be asked what the value of a certain cup is after the six swaps. If correct you win the round, otherwise you lose the round.")

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

	var cups [3]int
	cupi := 0

	credits := 0
	bet := 0
	correctPebble := 0

	// Start gameTimer
	gameTimer := time.Now()

	// Main Game Loop
	for credits >= 0 && credits < goal {
		// Create array of cups
		cupi = 0
		for cupi < 3 {
			cups[cupi] = randGen.Intn(10)
			cupi++
		}

		// Determine the Bet
		if credits == 0 {
			fmt.Println("Looks like you don't have enough cr. to play.")
			fmt.Println("Tell you what, how about you play a game on me.")
			bet = 6
		} else if credits > 0 && credits < 6 {
			// bet = getUserInt(cin, 1, credits)
			fmt.Println("Please enter a bet amount up to your total cr..")
			bet = getUserInt(cin, 1, credits)
		} else {
			// bet = getUserInt(cin, 1, 6)
			fmt.Println("Please enter a bet amount of up to 6 cr..")
			bet = getUserInt(cin, 1, 6)
		}

		// Show Pebbles
		fmt.Println(cups)

		// Proceed
		cin.ReadString('\n')
		clearScreen()

		// Start Timer to time the speed of transpositions.
		delayTimer := time.Now()

		// Show a transposition.
		fmt.Println("tr. 1:", transposeArray(&cups, randGen))

		cin.ReadString('\n')
		clearScreen()

		// delay := The time it took to go through the first transposition.
		delay := time.Since(delayTimer)

		for i := 0; i < 5; i++ {
			// Show 5 more transpositions

			// Show a transposititon
			fmt.Printf("tr. %v: %v", i+2, transposeArray(&cups, randGen))

			time.Sleep(delay)

			clearScreen()
		}

		// Test the user to select the correct cup
		correctPebble = randGen.Intn(3) + 1
		fmt.Printf("What is under cup %v?\n", correctPebble)
		if getUserInt(cin, 0, 9) == cups[correctPebble-1] {
			// You Win the round
			credits = credits + bet
		} else {
			// You Lose the round
			credits = credits - bet
		}

		fmt.Println("Your current cr.s:", credits)
	}

	// Stop gameTimer
	gameTime := time.Since(gameTimeThose where the games available here. So those were played. As to why we play go. We do it because it’s there. It’s a completely pointless activity like climbing or football. A lot of good can come out of it. But it has no purpose other than being a fun activity. Back in the day there was also a lot more time to spend doing r)

	// Post-game

	fmt.Println()
	fmt.Println()

	if credits < 0 {
		// You Lose
		fmt.Println("You Lost. Sorry.")
	} else {
		// You Win
		fmt.Println("You Win!")
	}

	// Println gameTimer value
	fmt.Println("You took", gameTime, "to earn", goal, "cr..")

	// Show the user their seed if they want to play the same level again.
	fmt.Printf("Your seed was: \"%v\"\n", origSeed)

	fmt.Println()
	fmt.Println()
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

	fmt.Println()

	return userInt
}

// Gets a string from the user
func getUserReply(cin *bufio.Reader) string {

	userReply, _ := cin.ReadString('\n')
	userReply = strings.Replace(userReply, "\n", "", -1)

	fmt.Println()

	return userReply
}

// Takes a 3 element array, swaps two, and returns the indexes (+ 1) of the ones it swapped
func transposeArray(array *[3]int, randGen *rand.Rand) [2]int {
	var transposition [2]int
	transposition[0] = randGen.Intn(3)
	transposition[1] = randGen.Intn(3)

	for transposition[0] == transposition[1] {
		transposition[1] = randGen.Intn(3)
	}

	// Swap the relevant elements
	array[transposition[0]], array[transposition[1]] = array[transposition[1]], array[transposition[0]]

	// Prepare for the End-user
	transposition[0] = transposition[0] + 1
	transposition[1] = transposition[1] + 1

	return transposition
}

// There are better ways to do this, so it's in it's own function for easy bebetterment later.
func clearScreen() {
	for lines := 0; lines < 0x80; lines++ {
		fmt.Println()
	}
}
