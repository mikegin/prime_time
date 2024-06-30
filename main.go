package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"math"
	"net"
)

const (
	HOST = "0.0.0.0"
	PORT = "8080"
	TYPE = "tcp"
)

type Input struct {
	Method string  `json:"method"`
	Number float64 `json:"number"`
}

type Output struct {
	Method string `json:"method"`
	Prime  bool   `json:"prime"`
}

func main() {
	fmt.Println("Starting Prime Time Server...")
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
		return
	}
	defer listen.Close()
	fmt.Println("Server listening on", HOST+":"+PORT)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %s", err)
			continue
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading data:", err)
			return
		}

		var input Input
		err = json.Unmarshal([]byte(message), &input)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			return
		}

		var output Output
		output.Method = "isPrime"
		inputIntNumber := int(input.Number)
		hasDecimalValues := float64(inputIntNumber)-input.Number != 0

		if hasDecimalValues {
			output.Prime = false
		} else {
			output.Prime = isPrime(inputIntNumber)
		}

		o, err := json.Marshal(output)
		if err != nil {
			fmt.Println("Error marshaling JSON:", err)
			return
		}
		o = append(o, byte('\n'))
		conn.Write(o)

		fmt.Printf("Received: %v\n", input)
	}
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value))/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
