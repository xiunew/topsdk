package topsdk_test

import (
	"fmt"
	"testing"
	"github.com/topsdk"
)

func TestWirelessShareTpwdQueryParam(t *testing.T) {
	topClient := topsdk.NewTopClient("", "")
	var p = topsdk.WirelessShareTpwdQueryParam{}
	p.PasswordContent = "￥3ggP09P71CF￥"
	fmt.Println("===== WirelessShareTpwdQueryParam_test =====")
	resp, err := topClient.Request(p)
	fmt.Println("resp:", resp, "err:", err)
	t.Log(resp, err)
}
