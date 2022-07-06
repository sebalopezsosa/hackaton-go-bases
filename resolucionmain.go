package main

import (
	"fmt"
	"strconv"

	"hackaton-go-bases/internal/file"
	"hackaton-go-bases/internal/service"
)

const (
	READ   = "READ"
	CREATE = "CREATE"
	UPDATE = "UPDATE"
	DELETE = "DELETE"
)

func main() {
	var tickets []service.Ticket
	f := &file.File{Path: "./tickets.csv"}
	read, err := f.Read()
	if err != nil {
		panic(err)
	}

	// SERVICE BOOKINGS
	tickets = getTickets(read)
	s := service.NewBookings(tickets)

	var (
		OPERACION string = "DELETE"
		UpdateId  int    = 1001
		ReadId    int    = 2
		DeleteId  int    = 1007
	)

	switch OPERACION {
	case READ:
		ticket, err := s.Read(ReadId)
		if err != nil {
			fmt.Println(err)
			return
		}
		ShowInfo(ticket)
	case CREATE:
		ticketC := service.Ticket{
			Names:       "Herna Torres",
			Email:       "hernan@email.com",
			Destination: "Destino **",
			Date:        "18:90",
			Price:       112,
		}
		ticketC, err = s.Create(ticketC)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = f.Write(ticketC, CREATE)
		if err != nil {
			fmt.Println(err)
			return
		}
	case UPDATE:
		tUpdate := service.Ticket{
			Names:       "Nombre Actualizado",
			Email:       "emailActualizado@email.com",
			Destination: "Buenos Aires UPDATE",
			Date:        "20:22",
			Price:       33,
		}
		ticketUp, err := s.Update(UpdateId, tUpdate)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = f.Write(ticketUp, UPDATE)
		if err != nil {
			fmt.Println(err)
			return
		}
	case DELETE:
		id, err := s.Delete(DeleteId)
		if err != nil {
			fmt.Println(err)
			return
		}
		err = f.Write(service.Ticket{Id: id}, DELETE)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ShowInfo(ticket service.Ticket) {
	fmt.Printf("\n*** N° %d *****\n Names: %s\nEmail: %s\nDestination: %s\nDate: %s\nPrice: %d\n\n", ticket.Id, ticket.Names, ticket.Email, ticket.Destination, ticket.Date, ticket.Price)
}

func getTickets(lines [][]string) []service.Ticket {
	var tickets []service.Ticket
	for _, value := range lines {
		id, err := strconv.Atoi(value[0])
		if err != nil {
			panic(err)
		}

		price, err := strconv.Atoi(value[5])
		if err != nil {
			panic(err)
		}

		ticket := service.Ticket{
			Id:          id,
			Names:       value[1],
			Email:       value[2],
			Destination: value[3],
			Date:        value[4],
			Price:       price,
		}

		tickets = append(tickets, ticket)
	}

	return tickets
}
