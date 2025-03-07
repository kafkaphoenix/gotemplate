package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

// UserResponse represents the JSON response for a user
type UserResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Country   string `json:"country"`
}

// NewGetCmd handles fetching a user by ID via API
func NewGetCmd(logger *zerolog.Logger) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get [id]",
		Short: "Get user by ID",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			userID := args[0]
			url := fmt.Sprintf("http://localhost:8081/users/%s", userID)

			resp, err := http.Get(url)
			if err != nil {
				logger.Error().Err(err).Msg("Failed to send request")
				cmd.PrintErrln("Error: Failed to retrieve user")
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				logger.Error().Int("status", resp.StatusCode).Msg("Failed to retrieve user")
				cmd.PrintErrln("Error: Failed to retrieve user")
				return
			}

			// Parse response
			var user UserResponse
			if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
				logger.Error().Err(err).Msg("Failed to decode response")
				cmd.PrintErrln("Error: Failed to parse user data")
				return
			}

			// Print user details
			fmt.Printf("User ID: %s\nName: %s %s\nNickname: %s\nEmail: %s\nCountry: %s\n",
				user.ID, user.FirstName, user.LastName, user.Nickname, user.Email, user.Country)
		},
	}

	return cmd
}
