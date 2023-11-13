package main
import "fmt"
func main () {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = 50
	fmt.Printf("Welcome to our %v booking service\n",conferenceName)
	fmt.Println("We have total ",conferenceTickets, "Tickets and ",remainingTickets ,"are avvailable")
	fmt.Println("Get your tickets here to attend")

}