package response

import "bitbucket.org/solo_works/carrieralexalambda/alexa"

// Response is the response back to the Alexa speech service
type Response struct {
	Version           string                 `json:"version"`
	SessionAttributes map[string]interface{} `json:"sessionAttributes,omitempty"`
	Body              Body                   `json:"response"`
}

// NewSpeechResponse builds a simple Alexa session response
func NewSpeechResponse(speech string) alexa.Response {
	r := alexa.Response{
		Version: "1.0",
		Body: alexa.ResBody{
			ShouldEndSession: true,
		},
	}
	r.Body.OutputSpeech = &alexa.Payload{
		Type: "PlainText",
		Text: speech,
	}

	return r
}

// NewDialogDelegateResponse builds a simple Alexa response to advance to the next step
func NewDialogDelegateResponse() alexa.Response {
	r := alexa.Response{
		Version: "1.0",
		Body: alexa.ResBody{
			ShouldEndSession: false,
		},
	}
	r.Body.Directives = append(r.Body.Directives, alexa.Directives{Type: "Dialog.Delegate"})

	return r
}
