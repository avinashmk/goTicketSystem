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
				Function: consts.SearchTrainPostAction,
				Name:     consts.SearchTrainOptionName,
			},
			{
				Function: consts.MakeReservPostAction,
				Name:     consts.MakeReservOptionName,
			},
			{
				Function: consts.CancelReservPostAction,
				Name:     consts.CancelReservOptionName,
			},
			{
				Function: consts.ViewReservPostAction,
				Name:     consts.ViewReservOptionName,
			},
		},
	}

	if gen.Role == consts.AdminRole {
		menu.Options = append(menu.Options,
			Option{
				Function: consts.AddTrainSchemaPostAction,
				Name:     consts.AddTrainSchemaOptionName,
			},
			Option{
				Function: consts.RemoveTrainSchemaPostAction,
				Name:     consts.RemoveTrainSchemaOptionName,
			},
			Option{
				Function: consts.ViewTrainSchemaPostAction,
				Name:     consts.ViewTrainSchemaOptionName,
			},
			Option{
				Function: consts.UpdateTrainSchemaPostAction,
				Name:     consts.UpdateTrainSchemaOptionName,
			},
		)
	}
	return
}
