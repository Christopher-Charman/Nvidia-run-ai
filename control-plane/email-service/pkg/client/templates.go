package client

type AboutToExpire struct {
	Name string `json:"name,omitempty"`
}

type AccountCreationFailed struct {
	Name string `json:"name,omitempty"`
}

type InternalTenantCreationFailed struct {
	TrialUserEmail string `json:"trial_user_email,omitempty"`
	TenantName     string `json:"tenant_name,omitempty"`
	ErrorMsg       string `json:"error_msg,omitempty"`
}

type FirstSign struct {
	Name         string `json:"name,omitempty"`
	SignInLink   string `json:"sign_in_link,omitempty"`
	TrialEndDate string `json:"trial_end_date,omitempty"`
}

type WelcomeToRunai struct {
	Name      string `json:"name,omitempty"`
	Username  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	TrialLink string `json:"trial_link,omitempty"`
}

type TrialExpired struct {
	Name string `json:"name,omitempty"`
}

type ContactSupport struct {
	Message  string `json:"message,omitempty"`
	User     string `json:"user,omitempty"`
	TenantId string `json:"tenant_id,omitempty"`
}
