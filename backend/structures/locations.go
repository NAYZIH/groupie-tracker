package structures

type Location struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
	Image     string   `json:"image"`
	Name      string   `json:"name"`
}
