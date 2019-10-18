package main

import (
	"fmt"
	"time"
)

var flag = false

//15713044976 03095800
//15713044975 36093800

//15713056472 00060900
//15713056472 42064500
type Mk struct {
	c chan int
}

type IRun interface {
	run() error
}

type HttpVisit struct {
	Url        string `json:"url"`
	Method     string `json:"method"`
	Body       string
	HeaderType string
}

type IHttp struct {
	Url        string `json:"url"`
	HeaderType string
}

func (ih *IHttp) do() error {
	fmt.Println("iHttp")
	return nil
}

type HttpPost struct {
	IHttp
	Body string
}

func (hp *HttpPost) do() error {
	fmt.Println("http post")
	return nil
}

func (hv *HttpVisit) run() error {

	return nil
}

type VisitInfo struct {
	runSignal chan bool
	runLink   []HttpVisit
	runTimes  int
}

func main() {
	//fmt.Println("hello world")
	//for i:=0;i<3000;i++ {
	//	pre()
	//}
	//flag = true
	//time.Sleep(time.Minute)

	hp := HttpPost{}
	hp.do()

	//ms := make([]*Mk,1000)
	//
	//for i:=0;i<1000;i++ {
	//	ms[i] = &Mk{
	//		c:make(chan int,0),
	//	}
	//	p2(ms[i])
	//}
	//
	//n := time.Now().UnixNano()
	//for _,m := range ms {
	//	m.c<-1
	//}
	//fmt.Println(time.Now().UnixNano()-n)
}

func p2(m *Mk) {
	go func() {
		select {
		case <-m.c:
			for i := 0; i < 10; i++ {
				fmt.Println(time.Now().UnixNano())
			}
		}
	}()
}

func pre() {
	hf := 0
	go func() {
		for {
			if flag {
				fmt.Println("???", time.Now().UnixNano())
				hf++
			}
			if hf == 10 {
				return
			}
		}
	}()
}
