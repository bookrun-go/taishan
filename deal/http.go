package deal

import (
	"bytes"
	"fmt"
	"github.com/myadamtest/gobase/httpclient"
	"io"
)

type HttpDealInfo struct {
	AbsoluteDealInfo
	Url    string
	Header map[string]string
	hcli   httpclient.Client
}

type PostDealInfo struct {
	HttpDealInfo
	Body        string
	ContentType string
	BodyReader  io.Reader
}

func (pd *PostDealInfo) sendRequest() error {
	resp := pd.hcli.Post(pd.Url, pd.ContentType, pd.BodyReader)
	fmt.Println(resp.ToString())
	return nil
}

func (pd *PostDealInfo) pre() error {
	pd.BodyReader = bytes.NewReader([]byte(pd.Body))
	pd.ContentType = pd.Header["Content-Type"]
	if pd.ContentType == "application/json" {
		pd.ContentType = pd.Header["Content-Type"]
	}

	pd.AbsoluteDealInfo.ready(pd.sendRequest)
	return nil
}

type GetDealInfo struct {
	HttpDealInfo
}

func (gd *GetDealInfo) sendRequest() error {
	_ = gd.hcli.Get(gd.Url)
	//fmt.Println(resp.ToString())
	return nil
}

func (gd *GetDealInfo) pre() error {
	gd.AbsoluteDealInfo.ready(gd.sendRequest)
	return nil
}
