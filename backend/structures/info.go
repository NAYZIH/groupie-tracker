package structures

type RelationResponse struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Track struct {
	Title   string `json:"title"`
	Preview string `json:"preview"`
}

type DeezerResponse struct {
	Data []Track `json:"data"`
}
