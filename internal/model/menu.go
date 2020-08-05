package model

import (
	"github.com/avinashmk/goTicketSystem/internal/consts"
)

// Menu place holder for menu objects
type Menu struct {
	Gen     *General
	Options []Option
}

// Option place holder for each Option item
type Option struct {
	Function string
	Name     string
}

// MakeMainMenu the main menu
func MakeMainMenu(gen *General) (menu *Menu) {
	menu = &Menu{
		Gen: gen,
		Options: []Option{
			{
				Function: consts.SearchTrainOptionFunc,
				Name:     consts.SearchTrainOptionName,
			},
			{
				Function: consts.MakeReservOptionFunc,
				Name:     consts.MakeReservOptionName,
			},
			{
				Function: consts.CancelReservOptionFunc,
				Name:     consts.CancelReservOptionName,
			},
			{
				Function: consts.ViewReservOptionFunc,
				Name:     consts.ViewReservOptionName,
			},
		},
	}
	return
}
