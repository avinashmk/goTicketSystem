package types

// Users specifies users collection schema relative to MongoDB
type Users struct {
	username string
	pwd      string
	role     string
}
