package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/jlucktay/go-getter/pkg/getter"
	"github.com/jlucktay/go-getter/pkg/minmaxer"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	file, errOpen := os.Open("SampleListOfUrls.txt")
	if errOpen != nil {
		log.Fatalf("error opening file: %v", errOpen)
	}
	defer file.Close()

	mm := minmaxer.New()
	gg := getter.New(30)
	wg := &sync.WaitGroup{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			mm.Add(gg.Get(url))
		}(scanner.Text())
	}
	wg.Wait()

	if errScan := scanner.Err(); errScan != nil {
		log.Fatalf("error while scanning: %v", errScan)
	}

	p := message.NewPrinter(language.English)
	fmt.Println("Content-Length results:")
	p.Printf("Minimum of %d and maximum of %d (total: %d) from %d URLs.\n",
		mm.Minimum(), mm.Maximum(), mm.Sum(), mm.Count())
}
