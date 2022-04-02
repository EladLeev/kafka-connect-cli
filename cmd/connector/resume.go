package connector

import (
	"fmt"
	"net/http"

	"github.com/mattcolombo/kafka-connect-cli/utilities"
	"github.com/spf13/cobra"
)

// TODO add capability to print outcome that is not JSON (see task restart); Probably this needs to become a common printer function

var ConnectorResumeCmd = &cobra.Command{
	Use:   "resume",
	Short: "short description",
	Long:  "long description",
	Run: func(cmd *cobra.Command, args []string) {
		var path string = "/connectors/" + connectorName + "/resume"
		fmt.Println("making a call to", path) // control statement print - TOREMOVE
		response, err := utilities.DoCallByPath(http.MethodPut, path, nil)
		if err != nil {
			fmt.Printf("The HTTP request failed with error %s\n", err)
		} else {
			utilities.PrintResponseJson(response)
		}
	},
}

func init() {
	ConnectorResumeCmd.Flags().StringVarP(&connectorName, "name", "n", "", "name of the connector to show (required)")
	ConnectorResumeCmd.MarkFlagRequired("name")
}
