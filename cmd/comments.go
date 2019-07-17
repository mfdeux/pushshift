package cmd

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/mfdeux/pushshift/pushshift"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(commentsCmd)
	commentsCmd.Flags().StringP("userAgent", "u", "testClient/0.1.0", "user agent of client")
	commentsCmd.Flags().StringP("query", "q", "", "query to filter comments")
}

var commentsCmd = &cobra.Command{
	Use: "comments",
	Run: func(cmd *cobra.Command, args []string) {
		userAgent, _ := cmd.Flags().GetString("userAgent")
		query, _ := cmd.Flags().GetString("query")
		client := pushshift.NewClient(userAgent)
		q := &pushshift.CommentQuery{}
		err := json.Unmarshal([]byte(query), q)
		if err != nil {
			log.Fatal(err)
		}
		comments, err := client.GetComments(q)
		if err != nil {
			log.Fatal(err)
		}
		output, err := json.Marshal(comments)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(output))
	},
}
