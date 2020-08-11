package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/avinashmk/goTicketSystem/internal/model"
	"github.com/avinashmk/goTicketSystem/internal/server/session"
	"github.com/avinashmk/goTicketSystem/internal/store"
	"github.com/avinashmk/goTicketSystem/logger"
	"go.mongodb.org/mongo-driver/mongo"
)

// Signin Signin
func Signin(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("Signin()")
	defer logger.Leave.Println("Signin()")

	if r.Method != http.MethodPost {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	gen := model.MakeGeneral(username)
	if user, err := store.FindUser(username); err != nil {
		if err == mongo.ErrNoDocuments { // User doesn't exist.
			t, err := template.ParseFiles("./web/templates/signup.html")
			if err != nil {
				logger.Err.Println("Unable to parse template ", err)
				http.Error(w, "Unable to login: Internal Error Occurred", http.StatusInternalServerError)
			} else {
				gen.Message = "User doesn't exist. Sign Up instead?"
				t.Execute(w, gen)
			}
		} else {
			http.Error(w, "Unable to login: Internal Error Occurred", http.StatusInternalServerError)
		}
	} else {
		gen.Role = user.Role
		if user.Authenticate(password) {
			if _, alreadyActive := session.New(w, gen); alreadyActive {
				t, err := template.ParseFiles("./web/templates/signin.html")
				if err != nil {
					logger.Err.Println("Unable to parse template ", err)
					http.Error(w, "Unable to login: Internal Error Occurred", http.StatusInternalServerError)
				} else {
					gen.Message = "User already active in another session"
					t.Execute(w, gen)
				}
			} else {
				t, err := template.ParseFiles("./web/templates/menu.html")
				if err != nil {
					logger.Err.Println("Unable to parse template ", err)
					http.Error(w, "Unable to login: Internal Error Occurred", http.StatusInternalServerError)
				} else {
					t.Execute(w, model.MakeMainMenu(gen))
				}
			}
		} else {
			t, err := template.ParseFiles("./web/templates/signin.html")
			if err != nil {
				logger.Err.Println("Unable to parse template ", err)
				http.Error(w, "Unable to login: Internal Error Occurred", http.StatusInternalServerError)
			} else {
				gen.Message = "Incorrect Password"
				t.Execute(w, gen)
			}
		}
	}
}
