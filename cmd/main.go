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
			fmt.Println("[ERROR] " + err.Error())
		} else {
			if dq.Type == "get" {
				b, err := dq.ExecuteGetQuery()
				if err != nil {
					fmt.Println("[ERROR] " + err.Error())
				}
				fmt.Println(b)
			}

			if dq.Type == "put" {
				k, err := dq.ExecutePutQuery()
				if err != nil {
					fmt.Println("[ERROR] " + err.Error())
				}
				fmt.Println(k)
			}
		}
	}

	//Handeling errors
	err := scanner.Err()
	if err == io.EOF {
		os.Exit(0)
	}
	if err != nil {
		fmt.Println("[ERROR] " + err.Error())
	}
}
