package t

import (
	"io"
	"net/http"
	"testing"
)

var (
	HostPath       = "http://localhost:8080/"
	StatusOK       = 200
	StatusOkEmpty  = 204
	StatusError    = 400
	StatusNotFound = 404
	Unknow         = "unknow"
)

func TestGet(t *testing.T) {
	//should succeed
	request, err := http.NewRequest("GET", HostPath+"domains/test.com", nil)
	if err != nil {
		t.Errorf("create request fail : " + err.Error())
	}
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		t.Errorf("request fail : " + err.Error())
	}
	if resp.StatusCode != StatusOK {
		t.Errorf("Get domain status didn't return StatusOK")
	}
	body, _ := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if string(body) != Unknow {
		t.Errorf("Get domain status didn't received the good status, Got %+v and expected %+v", string(body), Unknow)
	}
}

func TestPutDelivered(t *testing.T) {
	//should succeed
	request, err := http.NewRequest("PUT", HostPath+"events/test.com/delivered", nil)
	if err != nil {
		t.Errorf("create request fail : " + err.Error())
	}
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		t.Errorf("request fail : " + err.Error())
	}
	if resp.StatusCode != StatusOkEmpty {
		t.Errorf("Put event didn't return StatusOK")
	}
}

func TestPutBounced(t *testing.T) {
	//should succeed
	request, err := http.NewRequest("PUT", HostPath+"events/test.com/bounced", nil)
	if err != nil {
		t.Errorf("create request fail : " + err.Error())
	}
	resp, err := (&http.Client{}).Do(request)
	if err != nil {
		t.Errorf("request fail : " + err.Error())
	}
	if resp.StatusCode != StatusOkEmpty {
		t.Errorf("Put event didn't return StatusOK")
	}
}
