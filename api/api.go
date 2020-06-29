package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	model "github.com/RiccardoBiosas/golang-ethereum-auth/model"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Api struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *Api) Mount() {
	notFound := godotenv.Load()
	if notFound != nil {
		fmt.Println("error loading .env file")
	}
	var err error
	a.DB, err = sql.Open("mysql", os.Getenv("DB_AUTH"))
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
	a.mountRoutes()
}

func (a *Api) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *Api) mountRoutes() {
	a.Router.HandleFunc("/api/auth/login", a.getNonce).Methods("GET")
	a.Router.HandleFunc("/api/auth/login", a.sendSignature).Methods("POST")
	a.Router.HandleFunc("/api/auth/register", a.register).Methods("POST")
}

func (a *Api) register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin, access-control-allow-headers")
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	user.Nonce = "second_nonce"

	err = user.CreateUser(a.DB)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *Api) getNonce(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	pb := r.URL.Query().Get("pb")
	user := model.User{
		PublicKey: pb,
	}
	user.GetUserNonce(a.DB)
	fmt.Println("user nonce now", user.Nonce)
	resp, _ := json.Marshal(user.Nonce)
	w.Write(resp)
}

func (a *Api) sendSignature(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin, access-control-allow-headers")
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	defer r.Body.Close()
	decodedSig, err := hexutil.Decode(user.Signature)
	if err != nil {
		log.Fatal(err)
	}
	if decodedSig[64] != 27 && decodedSig[64] != 28 {
		return
	}
	decodedSig[64] -= 27
	user.GetUserNonce(a.DB)
	nonce := []byte(user.Nonce)
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
	isClientAddressEqualToRecoveredAddress := strings.ToLower(user.PublicKey) == strings.ToLower(recoveredAddress)
	if isClientAddressEqualToRecoveredAddress {
		user.Nonce = "NEW NONCE"
		user.UpdateNonce(a.DB)
	}
	json.NewEncoder(w).Encode(isClientAddressEqualToRecoveredAddress)

}