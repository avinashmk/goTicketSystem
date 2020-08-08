package model

import (
	"github.com/avinashmk/goTicketSystem/internal/consts"
)

// General general info for all templates
type General struct {
	UserID  string
	Role    string
	Message string
}

// MakeGeneral gives a general obj
func MakeGeneral(params ...string) (gen General) {
	gen.UserID = params[0]
	if len(params) > 1 {
		gen.Role = params[1]
	} else {
		gen.Role = consts.UserRole
		return
	}
	if len(params) > 2 {
		gen.Message = params[2]
	}
	return
}
