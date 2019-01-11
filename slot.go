package alexa

// Slot is an Alexa skill slot
type Slot struct {
	Name        string       `json:"name"`
	Value       string       `json:"value"`
	Resolutions *Resolutions `json:"resolutions"`
}
