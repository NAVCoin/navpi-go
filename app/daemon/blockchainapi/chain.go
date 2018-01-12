package blockchainapi

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/NAVCoin/navpi-go/app/conf"
	"io/ioutil"
	"log"

	"fmt"
	"github.com/NAVCoin/navpi-go/app/daemon/deamonrpc"
)



var config *conf.Config


// Setup all the handlers for the blockchain rpc interface
func InitHandlers(r *mux.Router, conf *conf.Config, prefix string)  {

	config = conf
	//
	r.HandleFunc(fmt.Sprintf("/%s/blockchain/v1/getblockcount", prefix), getBlockCount).Methods("GET")

	//// not implemented
	//r.HandleFunc("/blockchain/v1/getbestblockhash", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getblock", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getblockchaininfo", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getblockhash", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getblockhashes", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getblockheader", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getchaintips", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getdifficulty", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getmempoolancestors", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getmempoolentry", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getmempoolinfo", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getrawmempool", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/getspentinfo", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/gettxout", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/gettxoutproof", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/gettxoutsetinfo", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/verifychain", api.NotImplemented).Methods("GET")
	//r.HandleFunc("/blockchain/v1/verifytxoutproof", api.NotImplemented).Methods("GET")

}



func getBlockCount(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "NAVCoin pi server") // send data to client side

	log.Println("getBlockCount")

	n := deamonrpc.RpcRequestData{}
	n.Method = "getblockcount"

	resp, err := deamonrpc.RequestDaemon(n, config)

	if err != nil { // Handle errors requesting the daemon
		deamonrpc.RpcFailed(err, w, r)
		return
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write(bodyText)
}


