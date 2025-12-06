package auth

import (
	"galaveg/internal/modules/auth"
	"galaveg/internal/modules/errors"
	sessionsModule "galaveg/internal/modules/sessions"
	view "galaveg/internal/modules/view/layouts/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// LoginRequest defines the structure for the login form data.
// It includes validation tags to ensure the email is valid and the password meets requirements.
type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required,min=6"`
}

// Login handles the user authentication process.
// It serves the login page (GET) and processes the login credentials (POST).
func (ctr *Controller) Login(c *gin.Context) {
	// Get the current session
	session := sessions.Default(c)

	// Check if the user is already authenticated.
	// If the user ID exists in the session, redirect them to the panel immediately.
	if sessionsModule.ExistsUserId(ctr.ctx.Cfg, session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}

	// Initialize necessary services from the application context.
	authService := ctr.ctx.Services.Auth
	as := ctr.ctx.Services.App
	ts := ctr.ctx.Services.Translator
	ls := ctr.ctx.Services.Locales

	// Determine the current locale and initialize view data structures.
	locale := ctr.ctx.Services.Locales.Locale(c, nil)
	viewData := view.LoginViewData{}
	reqData := LoginRequest{}
	status := http.StatusOK

	// Handle the form submission (POST request).
	if c.Request.Method == "POST" {
		// Attempt to bind the incoming form data to the LoginRequest struct.
		if err := c.ShouldBind(&reqData); err != nil {
			// If binding fails, translate the validation errors for the UI.
			errs := ts.TVE(locale, err)

			for _, e := range errs {
				if e.Name == "Email" {
					viewData.EmailErrors = append(viewData.EmailErrors, e.GetMessage(ts.T(locale, "page.login.fields.email")))
				} else if e.Name == "Password" {
					viewData.PasswordErrors = append(viewData.PasswordErrors, e.GetMessage(ts.T(locale, "page.login.fields.password")))
				}
			}
		} else {
			// If binding succeeds, attempt to authenticate the user via the AuthService.
			userId, e := authService.Login(auth.NewEmailVO(reqData.Email), auth.NewPasswordVO(reqData.Password))
			if e != nil {
				// Handle authentication errors.
				status = e.Status
				if status >= 500 {
					viewData.Errors = append(viewData.Errors, ts.T(locale, errors.Translate500))
				} else {
					// Provide specific feedback if the user is not found or credentials are invalid.
					if e.Code == auth.ErrorLoginUserNotFound {
						viewData.Errors = append(viewData.Errors, ts.T(locale, auth.TranslateUserHasNotYetRegistered))
					} else {
						viewData.Errors = append(viewData.Errors, ts.T(locale, auth.TranslateCredentialsInvalid))
					}
				}
			} else {
				// Authentication successful: Register the user in the session.
				e = sessionsModule.Login(ctr.ctx.Cfg, session, userId)
				if e != nil {
					status = e.Status
					viewData.Errors = append(viewData.Errors, ts.T(locale, errors.Translate500))
				} else {
					// Redirect to the homepage upon successful login.
					c.Redirect(http.StatusFound, "/")
					return
				}
			}
		}

		// Preserve the entered values in the form in case of failure.
		viewData.EmailValue = reqData.Email
		viewData.PasswordValue = reqData.Password
	}

	// Prepare the view object with the gathered data.
	d, err := view.NewLogin(c, as, ls, ts, session, &viewData)
	if err != nil {
		// If the view generation fails, abort with an internal server error.
		//goland:noinspection GoUnhandledErrorResult
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	// Render the login HTML template.
	c.HTML(status, view.TEMPLATE, d)
}
