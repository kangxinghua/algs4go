package main

import (
	"paxos"
	"strconv"
	"fmt"
	"sync"
)

func main() {
	prepareCount := 11
	acceptors := make([]*paxos.Acceptor, 0)
	var wg sync.WaitGroup
	wg.Add(prepareCount)
	for i := 0; i < prepareCount; i++ {
		acceptors = append(acceptors, paxos.NewAcceptor("Acceptor_" + strconv.Itoa(i)))
	}
	for i := 0; i < prepareCount; i++ {
		p := paxos.NewProposer("Proposer_" + strconv.Itoa(i), i, "P_" + strconv.Itoa(i), acceptors)
		go func() {
			p.Run()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("success!!")
}