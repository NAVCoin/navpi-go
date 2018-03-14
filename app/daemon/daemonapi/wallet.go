package daemonapi

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/NAVCoin/navpi-go/app/conf"
	"github.com/NAVCoin/navpi-go/app/daemon/daemonrpc"
	"github.com/gorilla/mux"
	"github.com/muesli/crunchy"
)

// InitWalletHandlers sets up handlers for the blockchain rpc interface
func InitWalletHandlers(r *mux.Router, prefix string) {

	namespace := "wallet"
	r.HandleFunc(fmt.Sprintf("/%s/%s/v1/getstakereport", prefix, namespace), getStakeReport).Methods("GET")
	//r.HandleFunc(fmt.Sprintf("/%s/%s/v1/encryptwallet", prefix, namespace), encryptWallet).Methods("GET")

}

// checkPasswordStrength ensures password entered is safe
func checkPasswordStrength(pass string) error {

	validator := crunchy.NewValidator()
	err := validator.Check(pass)

	return err

}

// encryptWallet executes json RPC command and returns...
// func encryptWallet(w http.ResponseWriter, r *http.Request) {

// 	pass := "im the"
// 	apiResp, err := checkPasswordStrength(pass)

// 	if err != nil {
// 		apiResp.Send(w)
// 		return
// 	}

// 	// now execute ecnrypt command...

// 	// respond

// }

// getStakeReport takes writer, request - writes out stake report
func getStakeReport(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, "NAVCoin pi server") // send data to client side
	n := daemonrpc.RpcRequestData{}
	n.Method = "getstakereport"

	resp, err := daemonrpc.RequestDaemon(n, conf.NavConf)

	// Handle errors requesting the daemon
	if err != nil {
		daemonrpc.RpcFailed(err, w, r)
		return
	}

	bodyText, err := ioutil.ReadAll(resp.Body)
	w.WriteHeader(resp.StatusCode)
	w.Write(bodyText)
}
