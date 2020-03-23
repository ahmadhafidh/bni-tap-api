package helper

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/json"
	"strings"
)

type header struct {
	Typ string `json:"typ"`
	Alg string `json:"alg"`
}

func zip(a1, a2 []string) []string {
	r := make([]string, 2*len(a1))
	for i, e := range a1 {
		r[i*2] = e
		r[i*2+1] = a2[i]
	}
	return r
}

//genHMAC256 generates a hash signature of the encrypted text
func genHMAC256(ciphertext, key []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write([]byte(ciphertext))
	hmac := mac.Sum(nil)
	return hmac
}

// genSignature generate signature jwt
func GenSignature(data []byte) string {
	h := header{"JWT", "HS256"}

	arrF := []string{"+", "/", "="}
	arrT := []string{"-", "_", ""}

	rH, _ := json.Marshal(h)
	hEnc := b64.StdEncoding.EncodeToString([]byte(rH))
	newH := strings.NewReplacer(zip(arrF, arrT)...).Replace(hEnc)

	dEnc := b64.StdEncoding.EncodeToString([]byte(data))
	newD := strings.NewReplacer(zip(arrF, arrT)...).Replace(dEnc)

	headerData := string(newH) + "." + string(newD)
	// Create a new HMAC by defining the hash type and the key (as byte array)
	jwtKey := "e728632f-b55f-4cb8-bdae-ac0ddaabe960"
	hmc := hmac.New(sha256.New, []byte(jwtKey))

	// Write Data to it
	hmc.Write([]byte(headerData))

	// Get result and encode as hexadecimal string
	sha := genHMAC256([]byte(headerData), []byte(jwtKey))

	sign := strings.NewReplacer(zip(arrF, arrT)...).Replace(b64.StdEncoding.EncodeToString(sha))

	jwt := string(headerData) + "." + string(sign)

	return jwt
}
