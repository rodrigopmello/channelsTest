package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

func main() {
	args := os.Args[:2]
	fmt.Println("File to be processed: ", args[1])
	reader, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	scanner := bufio.NewScanner(reader)
	start := time.Now()

	var w8 sync.WaitGroup
	for scanner.Scan() {
		w8.Add(1)
		go processLine(scanner.Text(), &w8)
	}
	w8.Wait()

	fmt.Println("Total execution time:", time.Since(start))

}
