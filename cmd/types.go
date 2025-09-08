package cmd

const (
	configPath = "config.json"
	outputPath = "output/"
	firstName = "FirstName"
	lastName  = "LastName"
	location  = "Location"
	phone     = "Phone"
	email     = "Email"
)

type ExternalLink struct {
	Text, Link string
}

type coverLetterConfig struct {
	FirstName string
	LastName string
	Location string
	PhoneNumber string
	Email string
	Links []ExternalLink 
}