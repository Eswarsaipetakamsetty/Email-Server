package emailmodels

type EmailRequest struct {
	To      string `json:"to" validate:"required,email"`
	Subject string `json:"subject"`
}
