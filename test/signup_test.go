package test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestSuccessFullSignUp(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 5)
	rand.Read(b)
	email := fmt.Sprintf("%x", b) + "@gmail.com"
	fmt.Println(email)
}

func TestFailureFullSignUp(t *testing.T) {

	rand.Seed(time.Now().UnixNano())
	b := make([]byte, 5)
	rand.Read(b)
	email := fmt.Sprintf("%x", b) + "gmail.com"
	fmt.Println(email)

}
