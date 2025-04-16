package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var configUpdateCmd = &cobra.Command{
	Use: "update",
	Short: "Update fields in the existing config",
	Run: UpdateConfig,	
}

func init(){
	configCmd.AddCommand(configUpdateCmd)

	

	configUpdateCmd.Flags().StringP(firstName, "f", "", "Update your first name")
	configUpdateCmd.Flags().StringP(lastName, "l", "", "Update your last name")
	configUpdateCmd.Flags().StringP(location, "c", "", "Update your location")
	configUpdateCmd.Flags().StringP(phone, "p", "", "Update your phone number")
	configUpdateCmd.Flags().StringP(email, "e", "", "Update your email address")
}

const (
	firstName = "FirstName"
	lastName  = "LastName"
	location  = "Location"
	phone     = "Phone"
	email     = "Email"
)



func UpdateConfig(cmd *cobra.Command, args []string){
	

	bs, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println("Failed to read config:", err)
		return
	}

	var cfg coverLetterConfig
	if err := json.Unmarshal(bs, &cfg); err != nil {
		fmt.Println("Failed to read config:", err)
	}

	updateField := func(flag string, target *string) {
		if cmd.Flags().Changed(flag) {
			val, _ := cmd.Flags().GetString(flag)
			*target = strings.TrimSpace(val)
		}
	}

	updateField(firstName, &cfg.FirstName)
	updateField(lastName, &cfg.LastName)
	updateField(location, &cfg.Location)
	updateField(phone, &cfg.PhoneNumber)
	updateField(email, &cfg.Email)

	// Save updated config
	newBs, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		fmt.Println("Failed to encode updated config:", err)
		return
	}

	if err := os.WriteFile(configPath, newBs, 0644); err != nil {
		fmt.Println("Failed to write updated config:", err)
		return
	}

	fmt.Println("Config updated successfully!")
}