package main


import(
	"fmt"
	"runtime"
	"time"
	"math/rand"
)

var gameboard = make([]int, 9)


//being made fine
type player struct {
	name string
	order int
	computer bool
}

//working fine to print the board
func printBoard(board []int){
		fmt.Println(board[0], "|" , board[1], "|", board[2])
		fmt.Println("---------")
		fmt.Println(board[3], "|" , board[4], "|", board[5])
		fmt.Println("---------")
		fmt.Println(board[6], "|" , board[7], "|", board[8])
		fmt.Println("")
}

//check to see if someone has won
//add check case for board being full and no winner
func checkWinner(board []int)int{
	for i := 0; i < 3; i++ {
		if board [i*3] != 0 && board[i*3] == board[(i*3)+1] && board[i*3] == board[(i*3)+2] {
			return board[i*3]
		}
		if board[i] != 0 && board[i] == board[i+3] && board[i] == board[i+6] {
			return board[i]
		}
	}
	
	if board[0] != 0 && board[0] == board[4] && board[0] == board[8] {
		return board[0]
	}
	
	if board[2] != 0 && board[2] == board[4] && board[2] == board[6] {
		return board[2]
	}
	
	return -1
}

//channels seem to be working but it ruined play logic so just using random numbers
//need to keep this going until someone wins or board is full
func playGame(order int, name string, whoseTurn chan int){
	
	thisPlayer := player{name, order, true}  //each player thread has its own player
	fmt.Println(thisPlayer)

	for {
		turn := <- whoseTurn
		if turn != thisPlayer.order	{
			if turn == 2 {
				whoseTurn <- 1
			} else {whoseTurn <- 2}
		} else if turn == thisPlayer.order {
			for keepTrying := 0; keepTrying != 1; {
				i := rand.Intn(9)
				if gameboard[i] == 0 {
					gameboard[i] = thisPlayer.order
					keepTrying = 1
				}
			}

			printBoard(gameboard)

			if turn == thisPlayer.order {
				if thisPlayer.order == 1 {
					whoseTurn <- 2
				} else { whoseTurn <- 1}
			}
		}
		time.Sleep(1000*time.Millisecond)
		if checkWinner(gameboard) != -1 {
			break
		}
	}
}



func main(){
	runtime.GOMAXPROCS(2) //makes it use up to 2x cpu cores instead of 1 core interleaved

fmt.Println("           )          (                                              )       ")
fmt.Println(" (      ( /(     *   ))\\ )  (       *   )   (       (       *   ) ( /(       ")
fmt.Println(")\\ )   )\\())  ` )  /(()/(  )\\    ` )  /(   )\\      )\\    ` )  /( )\\()) (     ")
fmt.Println("(()/(  ((_)\\    ( )(_))(_)|((_)___ ( )(_)|(((_)(  (((_)___ ( )(_)|(_)\\  )\\   ")
fmt.Println(" /(_))_  ((_)  (_(_()|_)) )\\__|___(_(_()) )\\ _ )\\ )\\__|___(_(_())  ((_)((_)  ")
fmt.Println("(_)) __|/ _ \\  |_   _|_ _((/ __|  |_   _| (_)_\\(_|(/ __|  |_   _| / _ \\| __| ")
fmt.Println("  | (_ | (_) |   | |  | | | (__     | |    / _ \\  | (__     | |  | (_) | _|  ")
fmt.Println("   \\___|\\___/    |_| |___| \\___|    |_|   /_/ \\_\\  \\___|    |_|   \\___/|___| ")
                                                                             


	for i := 0; i < 9; i++ {
		gameboard[i] = 0
	}

	rand.Seed(time.Now().UTC().UnixNano())
	
	gameboard[rand.Intn(9)] = 2  //test code
	printBoard(gameboard)
	whoseTurn := make(chan int, 1) //make channel containing whose turn it is
	whoseTurn <- 1

	go playGame(1, "Bob", whoseTurn) 	//make the two threads
	go playGame(2, "Sue", whoseTurn)

	var input string
	fmt.Scanln(&input) //prevent main from closing b4 other threads fin

	//test code this had to go after fmt.Scanln or would check winner and declare 0 winner earlier
	//	gameboard[2] = 1
	//	gameboard[4] = 1
	//	gameboard[6] = 1

	printBoard(gameboard)
	foo := checkWinner(gameboard)
	if foo == -1 {
		fmt.Println("No winner")
	} else { fmt.Println("the winner is ", foo) }
}