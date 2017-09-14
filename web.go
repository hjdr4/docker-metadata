package metadata

import (
	"net/http"

	"github.com/Sirupsen/logrus"

	"github.com/gorilla/mux"
)

var iface string

func SetIface(_iface string) {
	iface = _iface
}

func ipv4(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	addr, err := GetIpForInterface(iface)
	if err != nil {
		logrus.Errorf(err.Error())
		panic(nil)
	}
	w.WriteHeader(200)
	w.Write([]byte(addr))
}

func router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/latest/meta-data/local-ipv4", ipv4).Methods("GET")
	return r
}

func Serve() {
	CreateDummyInterface()
	err := http.ListenAndServe("169.254.169.254:80", router())
	if err != nil {
		panic(err)
	}
}
