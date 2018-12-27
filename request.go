package alexa

// Request is an Alexa skill request
// see https://developer.amazon.com/docs/custom-skills/request-and-response-json-reference.html#request-format
type Request struct {
	Version string  `json:"version"`
	Session Session `json:"session"`
	Context Context `json:"context"`
	Body    struct {
		Type        string `json:"type"`
		RequestID   string `json:"requestId"`
		Timestamp   string `json:"timestamp"`
		Locale      string `json:"locale"`
		Intent      Intent `json:"intent,omitempty"`
		Reason      string `json:"reason,omitempty"`
		DialogState string `json:"dialogState,omitempty"`
	} `json:"request"`
}
