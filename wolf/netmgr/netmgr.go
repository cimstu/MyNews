package netmgr

import "net/http"

type NetWorkManager struct {

}

func (m* NetWorkManager) Init() {

}

func (m* NetWorkManager) Start() {
	http.HandleFunc("/WolfGame", m.HandleGetIn)
	http.ListenAndServe(":9898", nil)
}
