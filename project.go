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

//check to see if someone has won...could maybe do cleaner with switch or something else but this was easier
func checkWinner(board []int)int{
	value := 0

	if board[0] == board[1]  {
		if board[1] == board[2]{
			if board[0] != 0 {
			value = board[0]		
			}
		}
	}
	if board[3] == board[4]  {
		if board[4] == board[5]{
			if board[3] != 0 {
			value = board[3]		
			}
	}
}
	if board[6] == board[7]  {
		if board[7] == board[8]{
			if board[6] != 0 {
				value = board[6]	
			}
		}
	}
	if board[0] == board[3]  {
		if board[3] == board[6]{
			if board[0] != 0 {
				value = board[0]		
			}
		}	
	}
	if board[1] == board[4]  {
		if board[4] == board[7]{
			if board[1] != 0 {
				value = board[1]		
			}	
		}
	}
	if board[2] == board[5]  {
		if board[5] == board[8]{
			if board[2] != 0 {
				value = board[2]		
			}
		}
	}
	if board[0] == board[4]  {
		if board[4] == board[8]{
			if board[0] != 0 {
				value = board[0]		
			}
		}
	}
	if board[2] == board[4]  {
		if board[4] == board[6]{
			if board[2] != 0 {
			value = board[2]		
			}
		}
	}


 return value
}

//channels borked as heck and it ruined play logic, need to keep this going until someone wins
func playGame(order int, name string, whoseTurn chan int, board []int){
	thisPlayer := player{name, order, true}  //each player thread has its own player
	fmt.Println(thisPlayer)
		
	turn := <- whoseTurn
	fmt.Println(turn)

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

	//test code this had to go after fmt.Scanln or would check winner and declare 0 winner earlier
	gameboard[2] = 1
	gameboard[4] = 1
	gameboard[6] = 1

	printBoard(gameboard)
	foo := checkWinner(gameboard)
	if foo == 0 {
		fmt.Println("No winner")
	} else { fmt.Println("the winner is ", foo) }
}