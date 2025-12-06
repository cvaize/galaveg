package auth

import (
	"galaveg/internal/modules/alerts"
	"galaveg/internal/modules/auth"
	"galaveg/internal/modules/errors"
	sessionsModule "galaveg/internal/modules/sessions"
	view "galaveg/internal/modules/view/layouts/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RegisterRequest defines the structure for user registration form data
// It uses Gin's form binding tags for validation
type RegisterRequest struct {
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required,min=6"`
	ConfirmPassword string `form:"confirm_password" binding:"required,min=6"`
}

// Register handles the user registration process
// It serves both GET (display form) and POST (process registration) requests
func (ctr *Controller) Register(c *gin.Context) {
	// Get the current session
	session := sessions.Default(c)

	// If user is already logged in, redirect to panel
	if sessionsModule.ExistsUserId(ctr.ctx.Cfg, session) {
		c.Redirect(http.StatusFound, "/panel")
		return
	}

	// Get service instances from context
	authService := ctr.ctx.Services.Auth
	as := ctr.ctx.Services.App
	ts := ctr.ctx.Services.Translator
	ls := ctr.ctx.Services.Locales

	// Get current locale
	locale := ls.Locale(c, nil)

	// Initialize view data
	viewData := view.RegisterViewData{}
	reqData := RegisterRequest{}
	status := http.StatusOK

	// Handle POST request (form submission)
	if c.Request.Method == "POST" {
		valid := false

		// Bind form data to request struct and validate
		if err := c.ShouldBind(&reqData); err != nil {
			// Translate validation errors to current locale
			errs := ts.TVE(locale, err)

			// Map validation errors to appropriate form fields
			for _, e := range errs {
				if e.Name == "Email" {
					viewData.EmailErrors = append(viewData.EmailErrors, e.GetMessage(ts.T(locale, "page.register.fields.email")))
				} else if e.Name == "Password" {
					viewData.PasswordErrors = append(viewData.PasswordErrors, e.GetMessage(ts.T(locale, "page.register.fields.password")))
				} else if e.Name == "ConfirmPassword" {
					viewData.ConfirmPasswordErrors = append(viewData.ConfirmPasswordErrors, e.GetMessage(ts.T(locale, "page.register.fields.confirm_password")))
				}
			}
		} else {
			// Additional validation: check if passwords match
			if reqData.Password != reqData.ConfirmPassword {
				a := ts.T(locale, "page.register.fields.password")
				attributes := map[string]string{"attribute": a}
				viewData.PasswordErrors = append(viewData.PasswordErrors, "")
				viewData.ConfirmPasswordErrors = append(viewData.ConfirmPasswordErrors, ts.V(locale, "validation.confirmed", attributes))
			} else {
				valid = true
			}
		}

		// If form data is valid, attempt registration
		if valid {
			// Create email and password value objects
			e := authService.Register(auth.NewEmailVO(reqData.Email), auth.NewPasswordVO(reqData.Password))

			if e != nil {
				// Handle specific registration errors
				if e.Code == auth.ErrorRegisterDuplicateUser {
					status = http.StatusBadRequest
					viewData.Errors = append(viewData.Errors, ts.T(locale, auth.TranslateUserIsAlreadyRegistered))
				} else {
					// Handle other errors (typically server errors)
					status = e.Status
					viewData.Errors = append(viewData.Errors, ts.T(locale, errors.Translate500))
				}
			} else {
				// Registration successful - show success alert and redirect to login
				alert := alerts.NewSuccessAlert(ts.T(locale, "alert.register.success"))
				//goland:noinspection GoUnhandledErrorResult
				alerts.AddFlash(session, []alerts.AlertDto{alert})
				c.Redirect(http.StatusFound, "/login")
				return
			}
		}

		// Preserve form field values for re-display
		viewData.EmailValue = reqData.Email
		viewData.PasswordValue = reqData.Password
		viewData.ConfirmPasswordValue = reqData.ConfirmPassword
	}

	// Render the registration page
	d, err := view.NewRegister(c, as, ls, ts, session, &viewData)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.HTML(status, view.TEMPLATE, d)
}
