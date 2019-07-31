package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/uniris/Aerospike-Elexir-Port/pkg/query"
)

func main() {

	//Manage stop
	var gracefulStop = make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		s := <-gracefulStop
		log.Println(fmt.Sprintf("RECEIVED SIGNAL: %s", s))
		log.Println("Stopping ...")
		os.Exit(0)
	}()

	//Looping on commands
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		dq, err := query.NewDatabaseQuery(scanner.Text())
		if err != nil {
			fmt.Println(query.NewError("0", err.Error()).ToString())
		} else {
			err := dq.CheckValues()
			if err != nil {
				fmt.Println(query.NewError(dq.ID, err.Error()).ToString())
			} else {
				if dq.Data.Type == "get" {

					b, err := dq.ExecuteGetQuery()
					if err != nil {
						fmt.Println(query.NewError(dq.ID, err.Error()).ToString())
					}
					fmt.Println(b)
				}

				if dq.Data.Type == "put" {
					k, err := dq.ExecutePutQuery()
					if err != nil {
						fmt.Println(query.NewError(dq.ID, err.Error()).ToString())
					}
					fmt.Println(k)
				}
			}
		}
	}

	//Handling errors
	err := scanner.Err()
	if err == io.EOF {
		os.Exit(0)
	}
	if err != nil {
		log.Fatal(err.Error())
	}
}
