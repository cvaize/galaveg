package actions

import "galaveg/internal/modules/app"

func CreateResetPasswordLink(as *app.Service, email string) (string, error) {
	refUrl := as.RefUrl()
	refUrl = refUrl.JoinPath("reset-password-confirm")

	// Add query parameters
	q := refUrl.Query()
	q.Set("code", "code123") // TODO: replace with actual secure token
	q.Set("email", email)
	refUrl.RawQuery = q.Encode()

	// The final URL included in the email
	return refUrl.String(), nil
}
