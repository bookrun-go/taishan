package deal

import (
	"fmt"
	"github.com/myadamtest/gobase/httpclient"
	"net/http"
	"time"
)

type RRoot struct {
	CircleCount     int // 循环次数
	ConcurrentCount int // 并发数
	Drs             []*DealRoot
}

func NewRRoot(circle, concurrent int) *RRoot {
	r := &RRoot{}
	r.CircleCount = circle
	r.ConcurrentCount = concurrent

	r.Drs = NewDealRoots(concurrent)
	return r
}

func (r RRoot) do() error {

	for i := 0; i < r.CircleCount; i++ {
		for _, dr := range r.Drs {
			for _, c := range dr.DealChain {
				c.pre()
			}
		}

		for _, dr := range r.Drs {
			for _, c := range dr.DealChain {
				go c.goDoing()
			}
		}

		for _, dr := range r.Drs {
			for _, c := range dr.DealChain {
				_ = c.end()
			}
		}

		fmt.Println(">>>>>")
	}
	return nil
}

type DealRoot struct {
	DealChain []IDealInfo       // 处理链。
	Val       map[string]string // 全局变量
}

func NewDealRoots(concurrentCount int) []*DealRoot {
	if concurrentCount <= 0 {
		return nil
	}

	drs := make([]*DealRoot, concurrentCount)
	for i := 0; i < concurrentCount; i++ {
		drs[i] = &DealRoot{}
		get := &GetDealInfo{}
		get.Url = "http://www.baidu.com"
		get.EndSignal = make(chan uint8)
		get.DoSignal = make(chan uint8)
		get.hcli = *httpclient.NewClient(http.DefaultClient)
		get.Result = DealResult{}

		drs[i].DealChain = []IDealInfo{get} // 链条怎么形成，需要商榷。
	}

	return drs
}

type IDealInfo interface {
	pre() error
	goDoing() error
	end() error
}

type AbsoluteDealInfo struct {
	Name      string
	DoSignal  chan uint8
	EndSignal chan uint8
	BackKey   []string
	Result    DealResult
	startTime int64
	endTime   int64
}

func (ad *AbsoluteDealInfo) pre() error {
	panic("this is absolute need to impl")
}

func (ad *AbsoluteDealInfo) ready(f func() error) {
	go func() {
		select {
		case <-ad.DoSignal:
			{
				fmt.Println("<<<<<<<<<<<<<<<")
				err := f()
				if err != nil {
					fmt.Println(err)
				}
				ad.EndSignal <- 1
			}
		}
	}()
}

func (ad *AbsoluteDealInfo) goDoing() error {
	ad.startTime = time.Now().UnixNano()
	ad.DoSignal <- 1
	return nil // fixme 是否不用返回error
}

func (ad *AbsoluteDealInfo) end() error {
	<-ad.EndSignal
	ad.endTime = time.Now().UnixNano()
	ad.Result.ExpendTime = int(ad.endTime - ad.startTime)
	fmt.Println(ad.Result.ExpendTime)
	return nil // fixme 是否不用返回error
}
