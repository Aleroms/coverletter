package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use: "create",
	Short: "creates the cover letter at output directory",
	Long: `The create command generates a cover letter. It needs the
company name and job position to continue. Then, it outputs the pdf 
cover letter in /output`,
	Run: create,
}

func init(){
	rootCmd.AddCommand(createCmd)
	
	// flags for create command
	createCmd.Flags().StringP("Company", "c", "", "company applying to")
	createCmd.Flags().StringP("Position","p","","position applying to")

	createCmd.MarkFlagRequired("Company")
	createCmd.MarkFlagRequired("Position")
}

func create(cmd *cobra.Command, args []string){
	company, _ := cmd.Flags().GetString("Company")
	position, _ := cmd.Flags().GetString("Position")
	fmt.Printf("company: %s\tposition:%s\n",company, position)

	json.Unmarshal()

	// init PDF
	cfg := config.NewBuilder().
	WithOrientation(orientation.Vertical).
	WithPageSize(pagesize.A4).
	WithBottomMargin(15).
	WithTopMargin(15).
	WithRightMargin(15).
	WithLeftMargin(15).
	Build()

	m := maroto.New(cfg)

	addHeader(m)
}

// addHeader adds the contact information in MLA format
func addHeader(m core.Maroto) {

}