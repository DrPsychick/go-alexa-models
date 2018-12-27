package alexa

// Resolutions is an Alexa skill Resolution
type Resolutions struct {
	ResolutionPerAuthority []struct {
		Values []struct {
			Value struct {
				Name string `json:"name"`
				ID   string `json:"id"`
			} `json:"value"`
		} `json:"values"`
	} `json:"resolutionsPerAuthority"`
}
