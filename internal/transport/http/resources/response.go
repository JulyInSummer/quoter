package resources

type QuoteResponse struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}
