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
	AddTrainSchemaFormTemplate   = "./web/templates/addtrainschemaform.html"
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
	CookieAge          = 300
)

// Train Schema fields
const (
	TrainName   = "TrainName"
	TrainNumber = "TrainNumber"
	Frequency   = "Frequency"
	Tickets     = "Tickets"
	Stops       = "Stops"

	CheckboxOn = "on"
	Monday     = "Monday"
	Tuesday    = "Tuesday"
	Wednesday  = "Wednesday"
	Thursday   = "Thursday"
	Friday     = "Friday"
	Saturday   = "Saturday"
	Sunday     = "Sunday"

	AvailPrefix = "ticket["
	AvailClass  = "][class]"
	AvailCount  = "][count]"
	AvailFare   = "][fare]"

	StopPrefix   = "stop["
	StopPosition = "][position]"
	StopStation  = "][station]"
	StopArrival  = "][arrival]"
	StopDepart   = "][departure]"

	OriginPos = 1
	DestinPos = 99
)

// Chart fields
const (
	TrainSchemaID = "trainschema_id"
	Date          = "Date"
	Availability  = "Availability"
	TicketIDs     = "traintickets_id"
	ExpireAt      = "expireAt"
)
