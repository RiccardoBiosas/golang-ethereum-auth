package helpers

import (
        "math/rand"
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

