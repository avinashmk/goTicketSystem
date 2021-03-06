package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/avinashmk/goTicketSystem/internal/model"
	"github.com/avinashmk/goTicketSystem/internal/server/session"
	"github.com/avinashmk/goTicketSystem/internal/store"
	"github.com/avinashmk/goTicketSystem/logger"
)

// Signup Signup
func Signup(w http.ResponseWriter, r *http.Request) {
	logger.Enter.Println("Signup()")
	defer logger.Leave.Println("Signup()")

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
	confirmpassword := r.PostFormValue("confirmpassword")
	gen := model.MakeGeneral(username)
	if password != confirmpassword {
		t, err := template.ParseFiles("./web/templates/signup.html")
		if err != nil {
			logger.Err.Println("Unable to parse template ", err)
			http.Error(w, "Unable to login: Internal Error Occurred", http.StatusInternalServerError)
		} else {
			gen.Message = "Error Mis-match in passwords"
			t.Execute(w, &gen)
		}
	} else {
		if user, err := store.NewUser(username, password); err != nil {
			logger.Err.Println("Unable to create new UserDoc object: ", err)
			http.Error(w, "Unable to Register: Internal Error Occurred", http.StatusInternalServerError)
		} else {
			if user.RegisterUser() {
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
				logger.Err.Println("Unable to create new UserDoc object: ", err)
				http.Error(w, "Unable to Register: Internal Error Occurred", http.StatusInternalServerError)
			}
		}
	}
}
