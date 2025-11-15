package models

type Result struct {
	Password   string
	Username   string
	URI        string
	ItemName   string
	PwnedCount uint64
	Prefix     string
	Suffix     string
	Hash       string
	Pwned      bool
	WorkerID   string
}
