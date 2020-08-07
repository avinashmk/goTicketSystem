package session

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"time"

	"github.com/avinashmk/goTicketSystem/internal/consts"
	"github.com/avinashmk/goTicketSystem/internal/model"
	"github.com/avinashmk/goTicketSystem/logger"

	"golang.org/x/crypto/bcrypt"
)

// sessionData contains sessions sessionData
type sessionData struct {
	gen   *model.General
	token string
}

var (
	// ActiveSessions contains required sessionData for all sessions against userIDs
	ActiveSessions = make(map[string]sessionData)
)

// SetToken sets Tokens
func SetToken(w http.ResponseWriter, user *model.General) (success bool) {
	logger.Enter.Println("SetToken()")
	defer logger.Leave.Println("SetToken()")

	if _, alreadyExists := ActiveSessions[user.UserID]; alreadyExists {
		success = false
	} else {
		now := time.Now().String()
		id := now + user.UserID
		token := generateToken(id)
		expiry := time.Now().Add(121 * time.Second)
		http.SetCookie(w, &http.Cookie{
			Name:    consts.SessionTokenCookie,
			Value:   token,
			Expires: expiry,
		})
		http.SetCookie(w, &http.Cookie{
			Name:    consts.UserIDCookie,
			Value:   user.UserID,
			Expires: expiry,
		})
		ActiveSessions[user.UserID] = sessionData{
			gen:   user,
			token: token,
		}
		success = true
	}
	return
}

// GetToken sets Tokens
func GetToken(r *http.Request) (gen *model.General, httpStatus int) {
	logger.Enter.Println("GetToken()")
	defer logger.Leave.Println("GetToken()")
	httpStatus = http.StatusOK
	userCookie, err := r.Cookie(consts.UserIDCookie)
	if err == nil {
		if record, found := ActiveSessions[userCookie.Value]; found {
			tokenCookie, err1 := r.Cookie(consts.SessionTokenCookie)
			if err1 == nil {
				if tokenCookie.Value == record.token {
					gen = record.gen
				}
			} else {
				if err == http.ErrNoCookie {
					// If the cookie is not set, return an unauthorized status
					httpStatus = http.StatusUnauthorized
				} else {
					// For any other type of error, return a bad request status
					httpStatus = http.StatusBadRequest
				}
			}
		} else {
			httpStatus = http.StatusUnauthorized
		}
	} else {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			httpStatus = http.StatusUnauthorized
		} else {
			// For any other type of error, return a bad request status
			httpStatus = http.StatusBadRequest
		}
	}
	return
}

// UpdateToken sets Tokens
func UpdateToken() {
	logger.Enter.Println("UpdateToken()")
	defer logger.Leave.Println("UpdateToken()")
}

// RemoveToken sets Tokens
func RemoveToken() {
	logger.Enter.Println("RemoveToken()")
	defer logger.Leave.Println("RemoveToken()")
}

func generateToken(id string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(id), bcrypt.DefaultCost)
	if err != nil {
		logger.Err.Println("Unable to generate hash.")
	}
	// fmt.Println("Hash to store:", string(hash))

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}
