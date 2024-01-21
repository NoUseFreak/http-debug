package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "http-debug",
	Short: "Debug your http calls",
	Long:  ``,
    RunE: func(cmd *cobra.Command, args []string) error {
        if json, _ := cmd.Flags().GetBool("json"); json {
		    http.HandleFunc("/", printJson)
        } else {
            http.HandleFunc("/", print)
        }
        addr, _ := cmd.Flags().GetString("addr")
		fmt.Fprintf(os.Stderr, "Serving at %s\n", addr)

		return http.ListenAndServe(addr, nil)
	},
}

func Execute() {
	rootCmd.Flags().StringP("addr", "a", ":8080", "Address to listen on")
	rootCmd.Flags().BoolP("json", "j", false, "Print json instead of plain text")
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
