package consts

// Roles
const (
	UserRole  = "User"
	AdminRole = "Admin"
)

// Status
const (
	ActiveStatus = "Active"
	LockedStatus = "Locked"
)

// Handlers vs Names
const (
	SignInFunc = "/signin"
	SignUpFunc = "/signup"

	SearchTrainOptionFunc = "/searchtrain"
	SearchTrainOptionName = "Search Trains"

	MakeReservOptionFunc = "/makereservation"
	MakeReservOptionName = "Make a Reservation"

	CancelReservOptionFunc = "/cancelreservation"
	CancelReservOptionName = "Cancel a Reservation"

	ViewReservOptionFunc = "/viewreservation"
	ViewReservOptionName = "View my Reservations"
)
