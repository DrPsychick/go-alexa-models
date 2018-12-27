package alexa

// UpdatedIntent is to update the Intent
type UpdatedIntent struct {
	Name               string                 `json:"name,omitempty"`
	ConfirmationStatus ConfirmationStatus     `json:"confirmationStatus,omitempty"`
	Slots              map[string]interface{} `json:"slots,omitempty"`
}
