package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// UserRequest represents the JSON payload for user creation
type UserRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

// NewCreateCmd handles user creation via API
func NewCreateCmd(logger *zerolog.Logger) *cobra.Command {
	var firstName, lastName, nickname, email, country string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create a new user",
		Run: func(cmd *cobra.Command, args []string) {
			user := UserRequest{
				FirstName: firstName,
				LastName:  lastName,
				Nickname:  nickname,
				Email:     email,
				Country:   country,
			}

			// Convert struct to JSON
			jsonData, err := json.Marshal(user)
			if err != nil {
				logger.Error().Err(err).Msg("Failed to serialize request")
				cmd.PrintErrln("Error: Failed to create user request")
				return
			}

			// Send POST request
			resp, err := http.Post("http://localhost:8081/users", "application/json", bytes.NewBuffer(jsonData))
			if err != nil {
				logger.Error().Err(err).Msg("Failed to send request")
				cmd.PrintErrln("Error: Failed to send request")
				return
			}
			defer resp.Body.Close()

			// Handle response
			if resp.StatusCode == http.StatusCreated {
				logger.Info().Msg("User created successfully")
				fmt.Println("User created successfully!")
			} else {
				logger.Error().Int("status", resp.StatusCode).Msg("Failed to create user")
				cmd.PrintErrln("Error: Failed to create user")
			}
		},
	}

	// Flags
	cmd.Flags().StringVarP(&firstName, "first-name", "f", "", "First name (required)")
	cmd.Flags().StringVarP(&lastName, "last-name", "l", "", "Last name (required)")
	cmd.Flags().StringVarP(&nickname, "nickname", "n", "", "Nickname (optional)")
	cmd.Flags().StringVarP(&email, "email", "e", "", "Email address (required)")
	cmd.Flags().StringVarP(&country, "country", "c", "", "Country (required)")
	cmd.MarkFlagRequired("first-name")
	cmd.MarkFlagRequired("last-name")
	cmd.MarkFlagRequired("email")
	cmd.MarkFlagRequired("country")

	return cmd
}
