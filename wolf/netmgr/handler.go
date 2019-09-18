package netmgr

import "net/http"

func (m* NetWorkManager) HandleGetIn(w http.ResponseWriter, r* http.Request) {
	r.ParseForm()
	user_name := r.Form.Get("user-name")
	user_id := r.Form.Get("user-id")


}