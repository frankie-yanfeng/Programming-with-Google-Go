package main

import (
	"fmt"
	"sync"
)

//ChopS struct
type ChopS struct{ sync.Mutex }

//Philo struct
type Philo struct {
	leftCS, rightCS *ChopS
	number          int
}

func (p Philo) eat() {
	for {
		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("#%d philosopher is eating\n", p.number)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}
}

func main() {

	c := make(chan int)

	CSticks := make([]*ChopS, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(ChopS)
	}

	philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		if i < (i+1)%5 {
			philos[i] = &Philo{CSticks[i], CSticks[(i+1)%5], i + 1}
		} else {
			philos[i] = &Philo{CSticks[(i+1)%5], CSticks[i], i + 1}
		}
	}

	for i := 0; i < 5; i++ {
		go philos[i].eat()
	}

	<-c
}
