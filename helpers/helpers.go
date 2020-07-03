package helpers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	randomStrBytes := make([]byte, length)
	for i := range randomStrBytes {
		randomStrBytes[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(randomStrBytes)
}

func EnableGetRequestsCors(w *http.ResponseWriter) {
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func EnablePostRequestsCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Content-Type", "application/json")
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, access-control-allow-origin, access-control-allow-headers")
}

func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	resp, _ := json.Marshal(payload)

	w.WriteHeader(statusCode)
	w.Write(resp)
}
