package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/common-nighthawk/go-figure"
)

//scan a layer 4 port via tcp or udp
func gate(protocol, hostname string, port int) bool {
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 10*time.Second)

	if err != nil {
		return false
	}
	defer conn.Close()
	fmt.Printf("#         Port %d: %v Open\n", port, protocol)
	return true
}

func main() {
	fmt.Println()
	title := figure.NewFigure("(GATE)", "", true)
	subtitle := figure.NewFigure("Simple Port Scanner", "", true)
	title.Print()
	subtitle.Print()
	fmt.Println()
	var target string
	var protocol string
	fmt.Print("Enter target ip: ")
	fmt.Scanf("%s", &target)
	fmt.Print("Enter protocolo (tcp/udp): ")
	fmt.Scanf("%s", &protocol)
	fmt.Println()
	fmt.Printf("Scanning 65535 %v ports on target %v ===>\n\n", protocol, target)
	fmt.Println("##########################################")
	start := time.Now()
	var wg sync.WaitGroup
	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			gate(protocol, target, i)
		}(i)
	}
	wg.Wait()
	duration := time.Since(start)
	fmt.Println("##########################################")
	fmt.Println("\nRuntime in seconds: ", duration.Seconds())
}
