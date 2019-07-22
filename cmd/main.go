package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/uniris/Aerospike-Elexir-Port/pkg/query"
)

func main() {
	log.Println("Golang erlang Port Running ....")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(fmt.Sprintf("Entered text is: %s", scanner.Text()))
		dq, err := query.NewDatabaseQuery(scanner.Text())
		if err != nil {
			log.Println(err)
		} else {
			if dq.Type == "get" {
				b, err := dq.ExecuteGetQuery()
				if err != nil {
					log.Println(err)
				}
				log.Println(b)
			}

			if dq.Type == "put" {
				k, err := dq.ExecutePutQuery()
				if err != nil {
					log.Println(err)
				}
				log.Println(k)
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}
