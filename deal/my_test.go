package deal

import (
	"testing"
	"time"
)

func TestDo(t *testing.T) {
	r := NewRRoot(3, 3)
	_ = r.do()

	time.Sleep(time.Second * 10)
	//dr := DealRoot{}
	//
	//get := &GetDealInfo{}
	//get.Url = "http://www.baidu.com"
	//get.EndSignal = make(chan uint8)
	//get.DoSignal = make(chan uint8)
	//get.hcli = *httpclient.NewClient(http.DefaultClient)
	//
	//chain := make([]IDealInfo,0)
	//chain = append(chain,get)
	//
	//dr.DealChain = chain
	//
	//for _,v := range dr.DealChain {
	//	err := v.pre()
	//	if err!=nil {
	//		fmt.Println(err)
	//	}
	//}
	//
	//for _,v := range dr.DealChain {
	//	err := v.goDoing()
	//	if err!=nil {
	//		fmt.Println(err)
	//	}
	//}
}
