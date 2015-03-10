package main


import(
	"fmt"
	"runtime"
	"time"
	"math/rand"
)


type player struct {
	name string
	order int
	computer bool
}

func printBoard(board []int){
		fmt.Println(board[0], "|" , board[1], "|", board[2])
		fmt.Println("---------")
		fmt.Println(board[3], "|" , board[4], "|", board[5])
		fmt.Println("---------")
		fmt.Println(board[6], "|" , board[7], "|", board[8])
		fmt.Println("")
}

func playGame(order int, name string, whoseTurn chan int, board []int){
	thisPlayer := player{name, order, true}  //each player thread has its own player
	fmt.Println(thisPlayer)
	board[rand.Intn(9)] = 3
	printBoard(board)
	turn := <- whoseTurn
	fmt.Println(turn)
	
}

func main(){
	runtime.GOMAXPROCS(2)

	gameboard := make([]int, 9)		//declare gameboard
	for i := 0; i < 9; i++ {
		gameboard[i] = 0
	}
	printBoard(gameboard)		//test code

	rand.Seed(time.Now().UTC().UnixNano()) //seed random num generator
	
	gameboard[rand.Intn(9)] = 2
	printBoard(gameboard)
	whoseTurn := make(chan int, 1)
	

	name := "bob"

	go playGame(1, name, whoseTurn, gameboard)
	go playGame(2, "sue", whoseTurn, gameboard)
	var input string
	time.Sleep(2*time.Second)
	fmt.Scanln(&input) //prevent main from closing b4 other threads fin

}