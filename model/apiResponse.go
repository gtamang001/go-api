package model

// APIResponse is the simplified structure sent to the frontend UI
type APIResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
