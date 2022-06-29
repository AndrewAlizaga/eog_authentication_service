package encryption

import (
	"crypto/sha512"
	"encoding/base64"
)

const Salt = "Alexander Hamilton My name is Alexander Hamilton And there's a million things I haven't done But just you wait, just you wait"

func Encrypt(data string) string {

	encryptedData := hashData(data, []byte(Salt))

	return encryptedData
}

func hashData(data string, salt []byte) string {

	passwordBytes := []byte(data)

	var hasher = sha512.New()

	passwordBytes = append(passwordBytes, salt...)

	//writing bytes to the hasher
	hasher.Write(passwordBytes)

	// get the hashed data
	hashedData := hasher.Sum(nil)

	//convert hashed data to base64 encoded string
	encodedData := base64.URLEncoding.EncodeToString(hashedData)

	return encodedData
}

// Check if two passwords match
func DataMatch(hashedPassword, currPassword string,
	salt []byte) bool {

	var currPasswordHash = hashData(currPassword, salt)

	return hashedPassword == currPasswordHash
}
