package auth

import (
	"galaveg/app/dto"
	view "galaveg/app/view/layouts/auth"
	auth "galaveg/app/view/layouts/email/reset_password"
	"galaveg/utils/logger"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResetPasswordRequest struct {
	Email string `form:"email" binding:"required,email"`
}

// ResetPassword handles the "Forgot Password" flow.
//
// The method performs the following steps:
//  1. Redirects the user if they are already authenticated.
//  2. Validates the submitted email form.
//  3. Generates a password-reset confirmation URL.
//  4. Renders the email HTML template.
//  5. Sends the reset email to the user.
//  6. Stores a flash notification about the result.
//  7. Renders the Reset Password page.
//
// All errors are gracefully handled and displayed to the user.
func (ctr *Controller) ResetPassword(c *gin.Context) {
	session := sessions.Default(c)

	// -------------------------------------------------------------
	// 1. Do not allow password reset for authenticated users.
	// -------------------------------------------------------------
	if ctr.ctx.SS.ExistsUserId(session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}

	// Resolve the user's locale (from cookie, user profile, or default)
	locale := ctr.ctx.AS.Locale(c, nil)

	// View model for the HTML page
	viewData := view.ResetPasswordViewData{}

	// Struct for POST form data
	reqData := ResetPasswordRequest{}

	// Default response status
	status := http.StatusOK

	// -------------------------------------------------------------
	// 2. Handle POST request (user submitted the form)
	// -------------------------------------------------------------
	if c.Request.Method == "POST" {

		// Validate input using Gin + custom translator
		if err := c.ShouldBind(&reqData); err != nil {

			// Convert validation errors to localized messages
			errs := ctr.ctx.TS.TVE(locale, err)
			for _, e := range errs {
				if e.Name == "Email" {
					viewData.EmailErrors = append(
						viewData.EmailErrors,
						e.GetMessage(ctr.ctx.TS.T(locale, "page.reset_password.fields.email")),
					)
				}
			}

		} else {

			// -------------------------------------------------------------
			// 3. Generate the reset-password confirmation URL
			// -------------------------------------------------------------

			refUrl := ctr.ctx.AS.RefUrl()
			refUrl = refUrl.JoinPath("reset-password-confirm")

			// Add query parameters
			q := refUrl.Query()
			q.Set("code", "code123") // TODO: replace with actual secure token
			q.Set("email", reqData.Email)
			refUrl.RawQuery = q.Encode()

			// The final URL included in the email
			url := refUrl.String()

			// -------------------------------------------------------------
			// 4. Render email HTML template
			// -------------------------------------------------------------
			emailViewData := auth.ViewData{ResetPasswordLink: url}

			tplData, err1 := auth.New(c, ctr.ctx, &emailViewData)
			if err1 != nil {
				// Template context creation failed — internal error
				c.AbortWithError(http.StatusInternalServerError, err1)
				return
			}

			// Render email HTML
			html, err2 := ctr.ctx.TplS.Html(auth.TEMPLATE, tplData)
			if err2 != nil {
				c.AbortWithError(http.StatusInternalServerError, err2)
				return
			}

			// Plain-text fallback version of the email
			txt := url

			// Recipient
			to := reqData.Email

			// Email subject (localized)
			subject := ctr.ctx.TS.T(locale, "mail.reset_password.subject")

			// -------------------------------------------------------------
			// 5. Send the email
			// -------------------------------------------------------------
			sendErr := ctr.ctx.MS.SyncSendSimpleEmail(to, subject, html, txt)
			if sendErr != nil {

				// Log technical details, show generic error to user
				logger.Errorf("(500) Controller.ResetPassword.SyncSendSimpleEmail: %v", sendErr)

				viewData.Errors = append(
					viewData.Errors,
					ctr.ctx.TS.T(locale, "alert.reset_password.fail"),
				)

			} else {

				// -------------------------------------------------------------
				// 6. Email sent successfully → store flash notification
				// -------------------------------------------------------------
				alert := dto.NewSuccessAlert(
					ctr.ctx.TS.T(locale, "alert.reset_password.success"),
				)

				// Flash message will be displayed on /login
				ctr.ctx.AlS.AddFlash(session, []dto.Alert{alert})

				// Redirect to login page
				c.Redirect(http.StatusFound, "/login")
				return
			}
		}

		// Preserve the user-entered email after validation errors
		viewData.EmailValue = reqData.Email
	}

	// -------------------------------------------------------------
	// 7. Render the Reset Password page
	// -------------------------------------------------------------
	d, err := view.NewResetPassword(c, ctr.ctx, session, &viewData)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(status, view.TEMPLATE, d)
}
