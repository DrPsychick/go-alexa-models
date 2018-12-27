package response

// Response is the response back to the response speech service
type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Body              Body                   `json:"response"`
}

// NewSpeechResponse builds a simple response session response
func NewSpeechResponse(speech string) Response {
	r := Response{
		Version: "1.0",
		Body: Body{
			ShouldEndSession: true,
		},
	}
	r.Body.OutputSpeech = &Payload{
		Type: "PlainText",
		Text: speech,
	}

	return r
}

// NewDialogDelegateResponse builds a simple response response to advance to the next step
func NewDialogDelegateResponse() Response {
	r := Response{
		Version: "1.0",
		Body: Body{
			ShouldEndSession: false,
		},
	}
	r.Body.Directives = append(r.Body.Directives, Directives{Type: "Dialog.Delegate"})

	return r
}
