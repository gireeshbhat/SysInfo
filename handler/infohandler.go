package handler

import (
	"github.com/elastic/go-sysinfo"
	"net/http"
)

func Handler(response http.ResponseWriter, request *http.Response) {

}

func h() {
	process, err := sysinfo.Self()
}
