package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mfdeux/pushshift/pushshift"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(submissionsCmd)
	submissionsCmd.Flags().StringP("userAgent", "u", "testClient/0.1.0", "user agent of client")
	submissionsCmd.Flags().StringP("query", "q", "", "query to filter submissions")
}

var submissionsCmd = &cobra.Command{
	Use: "submissions",
	Run: func(cmd *cobra.Command, args []string) {
		userAgent, _ := cmd.Flags().GetString("userAgent")
		query, _ := cmd.Flags().GetString("query")
		client := pushshift.NewClient(userAgent)
		q := &pushshift.SubmissionQuery{}
		err := json.Unmarshal([]byte(query), q)
		if err != nil {
			log.Fatal(err)
		}
		submissions, err := client.GetSubmissions(q)
		if err != nil {
			log.Fatal(err)
		}
		output, err := json.Marshal(submissions)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(output))
	},
}
