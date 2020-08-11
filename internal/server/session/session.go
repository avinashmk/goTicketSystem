package session

import (
	"crypto/md5"
	"encoding/hex"
	"net/http"
	"sync"
	"time"

	"github.com/avinashmk/goTicketSystem/internal/consts"
	"github.com/avinashmk/goTicketSystem/internal/model"
	"github.com/avinashmk/goTicketSystem/logger"
	"golang.org/x/crypto/bcrypt"
)

// Session contains sessions Session
type Session struct {
	Gen        model.General
	tokenValue string
}

var (
	// activeSessions contains required Session for all sessions against userIDs
	activeSessions = make(map[string]Session)
	asMux          sync.Mutex
)

// New creates new Session
func New(w http.ResponseWriter, g model.General) (s Session, alreadyActive bool) {
	logger.Enter.Println("NewSession()")
	defer logger.Leave.Println("NewSession()")

	if _, alreadyActive = activeSessions[g.UserID]; alreadyActive {
		return
	}

	s = Session{
		Gen:        g,
		tokenValue: generateToken(g.UserID),
	}
	s.setCookies(w)

	asMux.Lock()
	activeSessions[g.UserID] = s
	asMux.Unlock()
	return
}

// Get sets Tokens
func Get(r *http.Request) (s Session, httpStatus int) {
	logger.Enter.Println("Get()")
	defer logger.Leave.Println("Get()")
	httpStatus = http.StatusOK

	if userCookie, err := r.Cookie(consts.UserIDCookie); err == nil {
		logger.Debug.Println("User from Cookie: ", userCookie.Value)

		var found bool
		asMux.Lock()
		s, found = activeSessions[userCookie.Value]
		asMux.Unlock()
		if found {
			logger.Debug.Println("activeSessions found: ", s.tokenValue)

			if tokenCookie, err1 := r.Cookie(consts.SessionTokenCookie); err1 == nil {
				logger.Debug.Println("Token from cookie: ", tokenCookie.Value)

				if tokenCookie.Value == s.tokenValue {
					// Success! Authentic request.
					logger.Debug.Println("Got session token for: ", s.Gen.UserID)
				}
			} else {
				if err1 == http.ErrNoCookie {
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

// Refresh sets Tokens
func (s *Session) Refresh(w http.ResponseWriter) (success bool) {
	logger.Enter.Println("Refresh()")
	defer logger.Leave.Println("Refresh()")
	logger.Debug.Println("Refresh session token for: ", s.Gen.UserID)

	refreshSession := Session{
		Gen:        s.Gen,
		tokenValue: generateToken(s.Gen.UserID),
	}
	refreshSession.setCookies(w)

	asMux.Lock()
	activeSessions[s.Gen.UserID] = refreshSession
	asMux.Unlock()

	return true
}

// Close clears all session data and deletes cookies.
func (s *Session) Close(w http.ResponseWriter) {
	logger.Enter.Println("Close()")
	defer logger.Leave.Println("Close()")

	asMux.Lock()
	delete(activeSessions, s.Gen.UserID)
	asMux.Unlock()

	http.SetCookie(w, &http.Cookie{
		Name:    consts.SessionTokenCookie,
		Value:   s.tokenValue,
		Expires: time.Now().Add(1 * time.Second),
		MaxAge:  1,
	})
	http.SetCookie(w, &http.Cookie{
		Name:    consts.UserIDCookie,
		Expires: time.Now().Add(1 * time.Second),
		MaxAge:  1,
	})
}

func (s *Session) setCookies(w http.ResponseWriter) {
	logger.Enter.Println("setCookies()")
	defer logger.Leave.Println("setCookies()")

	expiry := time.Now().Add(consts.CookieAge * time.Second)
	http.SetCookie(w, &http.Cookie{
		Name:     consts.SessionTokenCookie,
		Value:    s.tokenValue,
		Expires:  expiry,
		MaxAge:   consts.CookieAge,
		HttpOnly: true,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     consts.UserIDCookie,
		Value:    s.Gen.UserID,
		Expires:  expiry,
		MaxAge:   consts.CookieAge,
		HttpOnly: true,
	})
	logger.Debug.Println("Cookie set: ", s.tokenValue)
}

func generateToken(userid string) string {
	logger.Enter.Println("generateToken()")
	defer logger.Leave.Println("generateToken()")

	now := time.Now().String()
	id := userid + now
	hash, err := bcrypt.GenerateFromPassword([]byte(id), bcrypt.DefaultCost)
	if err != nil {
		logger.Err.Println("Unable to generate hash.")
	}
	hasher := md5.New()
	hasher.Write(hash)
	t := hex.EncodeToString(hasher.Sum(nil))
	logger.Debug.Println("Generated token: ", t)
	return t
}
