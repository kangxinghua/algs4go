package paxos

import (
	"fmt"
	"time"
	"math/rand"
)

type proposer struct {
	name      string
	id        int
	N         int
	V         string
	acceptors []*Acceptor
}

func NewProposer(name string, N int, V string, acceptors []*Acceptor) *proposer {
	p := new(proposer)
	p.name = name
	p.id = N
	p.N = N
	p.V = V
	p.acceptors = acceptors
	return p
}

func (p *proposer) prepare() {

}

func (p *proposer) accept() {

}

func (p *proposer) Run() {
	x := rand.Intn(10)
	time.Sleep(time.Millisecond * time.Duration(x))// 随机暂停

	acceptorCount := len(p.acceptors)
	quorum := (acceptorCount / 2)
	prepareTime := 0;
	for ; ; {
		prepareTime++
		responseCount := 0
		maxN := -1
		maxV := ""
		for i := 0; i < acceptorCount; i++ {
			a := p.acceptors[i]
			pok, AcceptN, AcceptV := a.prepare(p.N)
			//fmt.Printf("%s prepare(%d) %s reture pok=%t,AcceptN=%d,AcceptV=%s \n", p.name, p.N, a.name, pok, AcceptN, AcceptV)
			if (pok) {
				responseCount++
				if (AcceptN > -1&&AcceptN > maxN) {
					maxN = AcceptN
					maxV = AcceptV
				}
			}
		}
		//fmt.Printf("%s N=%d V=%s maxN=%d responseCount=%d \n", p.name, p.N, p.V, maxN, responseCount)
		if (maxN > -1) {
			p.N = maxN
			p.V = maxV
		}
		if (responseCount > quorum) {
			acceptCount := 0;
			for i := 0; i < acceptorCount; i++ {
				a := p.acceptors[i]
				if (a.accept(p.N, p.V)) {
					acceptCount++;
				}
			}
			if (acceptCount >= quorum) {
				fmt.Printf("=====================================================%s N=%d V=%s \n", p.name, p.N, p.V)
				break;
			}
		}
		p.N = (acceptorCount * prepareTime) + p.id;
		//fmt.Printf("%s next N=%d V=%s \n", p.name, p.N, p.V, )
	}
}
