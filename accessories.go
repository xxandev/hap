package hap

import (
	"github.com/brutella/hap/accessory"
	"github.com/brutella/hap/log"

	"net/http"
)

type AccessoriesPayload struct {
	Accessories []*accessory.A `json:"accessories"`
}

func (srv *Server) GetAccessories(res http.ResponseWriter, req *http.Request) {
	if !srv.isPaired() {
		log.Info.Println("not paired")
		jsonError(res, JsonStatusInsufficientPrivileges)
		return
	}

	var as []*accessory.A
	as = append(as, srv.a)
	as = append(as, srv.as[:]...)

	p := AccessoriesPayload{as}
	log.Debug.Println(toJSON(p))
	jsonOK(res, p)
}