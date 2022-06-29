package account

import (
	"fmt"
	"time"

	enc "github.com/AndrewAlizaga/eog_authentication_service/app/security/encryption"
)

type Account struct {
	Id             string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name           string    `json:"name"`
	AccountStatus  string    `json:"account_status"`
	RecentActivity string    `json:"recent_activity"`
	Password       string    `json:"password"`
	CreationDate   time.Time `json:"creation_date"`
	Email          string    `json:"email"`
}

type ServiceAccount struct {
	AccountId string `json:"account_id"`
}

type AccountRequest struct {
	Name     string `json:"name"`
	OwnerId  string `json:"owner_id"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AccountInterface interface {
	GetPassword() string

	GetEmail() string

	ComparePassword(pass string) bool
}

//Interface methods
func (e *Account) GetPassword() string {

	return e.Password

}

func (e *Account) ComparePassword(pass string) bool {
	fmt.Println("CURRENT USER: ", e)
	fmt.Println("CURRENT ENCRYPTED PASS: ", e.Password)
	fmt.Println("NON ENCRYPTED PASS: ", pass)

	hashed := enc.Encrypt(pass)
	fmt.Println("after enc ", hashed)
	return e.Password == hashed
}
