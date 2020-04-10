package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

const keyEnv = "CONSTELLIX_API_KEY"
const secretEnv = "CONSTELLIX_SECRET_KEY"

func buildSecurityToken(key string, secret string) string {
	millis := time.Now().UnixNano() / 1000000
	timestamp := strconv.FormatInt(millis, 10)
	mac := hmac.New(sha1.New, []byte(secret))
	mac.Write([]byte(timestamp))
	hmacstr := base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return key + ":" + hmacstr + ":" + timestamp
}

func main() {
	curlHeaderFlag := flag.Bool("c", false, "Print as a curl header")
	debugFlag := flag.Bool("d", false, "Output debug information")
	flag.Parse()

	apiKey := os.Getenv(keyEnv)
	secretKey := os.Getenv(secretEnv)
	if *debugFlag {
		fmt.Println(keyEnv, "=", apiKey)
		fmt.Println(secretEnv, "=", secretKey)
	}
	token := buildSecurityToken(apiKey, secretKey)
	if *curlHeaderFlag {
		fmt.Printf("x-cns-security-token: %s", token)
	} else {
		fmt.Printf(token)
	}
}
