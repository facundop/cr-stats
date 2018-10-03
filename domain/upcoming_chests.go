package domain

// UpcomingChests is a representation of the upcoming chests for a player
type UpcomingChests struct {
	Items []struct {
		Index int    `json:"index"`
		Name  string `json:"name"`
	} `json:"items"`
}
