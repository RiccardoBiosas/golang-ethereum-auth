package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gorilla/mux"
)

type Api struct {
	Router *mux.Router
}

type clientSignature struct {
	PublicKey string `json:"pb"`
	Signature string `json:"sig"`
}

func (a *Api) Mount() {
	a.Router = mux.NewRouter()
	a.mountRoutes()
}

func (a *Api) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *Api) mountRoutes() {
	a.Router.HandleFunc("/api/auth", a.getNonce).Methods("GET")
	a.Router.HandleFunc("/api/auth", a.sendSignature).Methods("POST")
}

func (a *Api) getNonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode("test string")
}

func (a *Api) sendSignature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin, access-control-allow-headers")

	var clientSig clientSignature
	err := json.NewDecoder(r.Body).Decode(&clientSig)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	defer r.Body.Close()

	decodedSig, err := hexutil.Decode(clientSig.Signature)
	if err != nil {
		log.Fatal(err)
	}

	if decodedSig[64] != 27 && decodedSig[64] != 28 {
		return
	}
	decodedSig[64] -= 27

	nonce := []byte("test string")
	prefixedNonce := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(nonce), nonce)
	hash := crypto.Keccak256Hash([]byte(prefixedNonce))

	recoveredPublicKey, err := crypto.Ecrecover(hash.Bytes(), decodedSig)
	if err != nil {
		log.Fatal(err)
	}

	secp256k1RecoveredPublicKey, err := crypto.UnmarshalPubkey(recoveredPublicKey)
	if err != nil {
		log.Fatal(err)
	}

	recoveredAddress := crypto.PubkeyToAddress(*secp256k1RecoveredPublicKey).Hex()

	isClientAddressEqualToRecoveredAddress := strings.ToLower(clientSig.PublicKey) == strings.ToLower(recoveredAddress)

	json.NewEncoder(w).Encode(isClientAddressEqualToRecoveredAddress)

}
