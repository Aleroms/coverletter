package cmd

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Set up or update your user configuration details for the program",
	Long: `The 'config' command allows you to initialize or modify user-specific settings 
for the program. It collects relevant information such as your name, email, and other 
preferences, and stores them in a configuration file for consistent reuse across 
commands. This helps streamline the user experience and ensures your data is readily 
available when generating personalized cover letters.`,
	Run: configure,
}

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


func init(){
	rootCmd.AddCommand(configCmd)
}
func configure(cmd *cobra.Command, args []string) {

	fmt.Println("Please enter name, location, phone number, email, and any external links")
	fmt.Println()

	firstName := userInput("first name")
	lastName := userInput("last name")
	location := userInput("location")
	phoneNumber := userInput("phone number")
	email := userInput("email")
	links := userInputForLinks()

	coverLetterConfiguration := coverLetterConfig{
		FirstName:   firstName,
		LastName:    lastName,
		Location:    location,
		PhoneNumber: phoneNumber,
		Email:       email,
		Links:       links,
	}

	bs, err := json.MarshalIndent(coverLetterConfiguration, "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal configuration:", err)
		return
	}

	if err = os.WriteFile("config.json", bs, 0644); err != nil {
		fmt.Println("Failed to write config file:", err)
		return
	}

	fmt.Println("Configuration saved to config.json successfully.")

}

func userInputForLinks() []ExternalLink {
	links := []ExternalLink{}


	fmt.Println("Please enter external links. text is the label. link is the link to the resource.")
	stop := "y"

	var text string
	var link string

	for stop == "y" {

		text = userInput("text")
		link = userInput("link")

		links = append(links, ExternalLink{Text: text, Link: link})

		fmt.Printf("Another link (y/n)? ")
		stop = userInput("decision")
	}

	return links
}

func userInput(category string) string {
	var input string

	rdr := bufio.NewReader(os.Stdin)
	for input == "" {
		fmt.Printf("Enter %s: ", category)
		uin, err := rdr.ReadString(byte('\n'))

		if err != nil {
			fmt.Println("An error occurred, please try again")
		}

		input = uin
	}
	return input
}
