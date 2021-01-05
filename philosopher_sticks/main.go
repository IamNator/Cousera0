package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){

	var wg sync.WaitGroup
	wg.Add(5)


	philos := make([]*Philo, 5)
	cs := make([]*chopStick, 5)
	chanSlice :=  make([]chan int, 5)

	for i, _ := range chanSlice {
		chanSlice[i] = make(chan int)
	}

	CreateFivePhilo(philos, cs)
	go host(chanSlice)

	for i, philo := range philos {
		go philo.Eat(chanSlice[i], &wg, i)
	}


	 wg.Wait()

}


type chopStick struct {
	sync.Mutex
}

type Philo struct {
	rightCS, leftCS *chopStick
}


func (p *Philo) Eat(ch chan int, wg *sync.WaitGroup, num int){


	for _ = range ch {

		p.leftCS.Lock()
		p.rightCS.Lock()

		fmt.Printf("starting to eat %d\n", num+1)
		time.Sleep(2 * time.Second)
		fmt.Printf("finishing eating %d\n", num+1)

		p.rightCS.Unlock()
		p.leftCS.Unlock()
	}

	wg.Done()
}

func CreateFiveCS(cs []*chopStick){
	for i:=0;i<5;i++{
		cs[i] = new(chopStick)
	}
}

func CreateFivePhilo(p []*Philo, cs []*chopStick){

	CreateFiveCS(cs)
	for i:=0; i<5; i++ {
		p[i] = &Philo{
			cs[i], cs[(i+1)%5],
		}
	}
}

func host(ch []chan int){

	for i:=0; i<3; i++ {

		ch[0] <- i
		ch[3] <- i
		time.Sleep(70 * time.Millisecond)
		ch[1] <- i
		ch[4] <- i
		time.Sleep(70 * time.Millisecond)
		ch[0] <- i
		time.Sleep(70 * time.Millisecond)

	}

	for _,c := range ch {
		close(c)
	}

}

