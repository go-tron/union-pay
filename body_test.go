package unionPay

import (
	"testing"
)

func TestBody(t *testing.T) {
	a := "accessType=0&bizType=你好呀&amt={a=试一下&c=d}"

	result, err := UnmarshalBody(a)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}

func TestReserved(t *testing.T) {
	a := "e2Rpc2NvdW50QW10PTQwMDB9"

	result, err := GetReserved(a)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("result", result)
}
