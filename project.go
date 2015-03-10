package main


import(
	"fmt"
	"runtime"
	"time"
	"math/rand"
)


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

//channels borked as heck and it ruined play logic, need to keep this going until someone wins
func playGame(order int, name string, whoseTurn chan int, board []int){
	thisPlayer := player{name, order, true}  //each player thread has its own player
	fmt.Println(thisPlayer)
	

		
	turn := <- whoseTurn

	if turn == thisPlayer.order {
		i := rand.Intn(9)
		if board[i] == 0 {
			board[i] = thisPlayer.order
		}
	}
	if turn == thisPlayer.order {
		if thisPlayer.order == 1 {
			whoseTurn <- 2
		} else { whoseTurn <- 1}
	}
}

func main(){
	runtime.GOMAXPROCS(2) //makes it use 2x cpu cores instead of 1 core interleaved

	gameboard := make([]int, 9)		//declare gameboard
	for i := 0; i < 9; i++ {
		gameboard[i] = 0
	}
	printBoard(gameboard)		//test code

	rand.Seed(time.Now().UTC().UnixNano()) //seed random num generator
	
	gameboard[rand.Intn(9)] = 2  //test code
	printBoard(gameboard)
	whoseTurn := make(chan int, 1) //make channel containing whose turn it is
	whoseTurn <- 1

	go playGame(1, "Bob", whoseTurn, gameboard) 	//make the two threads
	go playGame(2, "Sue", whoseTurn, gameboard)

	var input string
	fmt.Scanln(&input) //prevent main from closing b4 other threads fin
}