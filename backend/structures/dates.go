package structures

type Dates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
	Image string   `json:"image"`
	Name  string   `json:"name"`
}
