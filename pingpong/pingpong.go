package main

import (
	"fmt"
	"math/rand"
	"sync"
)

const maxScore = 11

func main() {
	ch := make(chan string)
	wg := sync.WaitGroup{}
	p1, p2 := "Player 1", "Player 2"
	results := map[string]int{
		p1: 0,
		p2: 0,
	}

	for results[p1] < maxScore && results[p2] < maxScore {
		wg.Add(2)
		go round(p1, ch, &wg)
		go round(p2, ch, &wg)
		ch <- "begin"
		wg.Wait()
		results[<-ch]++
		fmt.Printf("%s - %v; %s - %v\n", p1, results[p1], p2, results[p2])
	}
}

func round(p string, ch chan string, wg *sync.WaitGroup) {
	for {
		m := <-ch

		if m == "stop" {
			wg.Done()
			ch <- p
			return
		}

		if m != "begin" && rand.Intn(9) < 2 {
			ch <- "stop"
			wg.Done()
			return
		}

		if m == "ping" {
			m = "pong"
		} else {
			m = "ping"
		}
		fmt.Printf("%s: %s!\n", p, m)

		ch <- m
	}
}
