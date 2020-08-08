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

// POST Action vs UI Names
const (
	SignInFunc  = "/signin"
	SignUpFunc  = "/signup"
	SignOffFunc = "/signoff"

	SearchTrainPostAction = "/searchtrain"
	SearchTrainOptionName = "Search Trains"

	MakeReservPostAction = "/makereservation"
	MakeReservOptionName = "Make a Reservation"

	CancelReservPostAction = "/cancelreservation"
	CancelReservOptionName = "Cancel a Reservation"

	ViewReservPostAction = "/viewreservation"
	ViewReservOptionName = "View Reservations"

	// Admin specific Options
	AddTrainSchemaFormPostAction = "/addtrainschemaform"
	AddTrainSchemaPostAction     = "/addtrainschema"
	AddTrainSchemaOptionName     = "Add Train Schema"

	RemoveTrainSchemaPostAction = "/removetrainschema"
	RemoveTrainSchemaOptionName = "Remove Train Schema"

	ViewTrainSchemaPostAction = "/viewtrainschema"
	ViewTrainSchemaOptionName = "View Train Schema"

	UpdateTrainSchemaPostAction = "/updatetrainschema"
	UpdateTrainSchemaOptionName = "Update Train Schema"
)

// Server specific
const (
	SessionTokenCookie = "session_token"
	UserIDCookie       = "user_id"
)

// Web Templates, Statics locations
const (
	AddTrainSchemaFormTemplate = "./web/templates/addtrainschemaform.html"
)
