package paxos

import (
	"math/rand"
)

type Acceptor struct {
	name          string
	ResN, AcceptN int
	AcceptV       string
}

func NewAcceptor(name string) *Acceptor {
	a := new(Acceptor)
	a.ResN = -1
	a.AcceptN = -1
	a.name = name
	return a
}

func randResponse() bool {
	//%50 概率调用失败
	b := 100 - rand.Intn(100) > 50
	//fmt.Printf("is call Success= %t \n", b)
	return b
}

func (a *Acceptor) prepare(N int) (bool, int, string) {
	if (randResponse()) {
		return false, -1, ""
	}
	if (N > a.ResN) {
		a.ResN = N
		return true, a.AcceptN, a.AcceptV
	} else {
		return false, -1, ""
	}
}

func (a *Acceptor) accept(N int, V string) bool {
	if (randResponse()) {
		return false
	}
	if (a.AcceptV == "") {
		a.AcceptN = N
		a.AcceptV = V
		return true;
	} else if (N < a.ResN&&V == a.AcceptV) {
		a.AcceptN = N
		a.AcceptV = V
		return true;
	} else {
		return false
	}
	return false
}