package domain

// BattleLog is a representation of the latest matches
type BattleLog struct {
	Type       string `json:"type"`
	BattleTime string `json:"battleTime"`
	Arena      struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"arena"`
	GameMode struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"gameMode"`
	DeckSelection string `json:"deckSelection"`
	Team          []struct {
		Tag              string `json:"tag"`
		Name             string `json:"name"`
		StartingTrophies int    `json:"startingTrophies"`
		TrophyChange     int    `json:"trophyChange"`
		Crowns           int    `json:"crowns"`
		Clan             struct {
			Tag     string `json:"tag"`
			Name    string `json:"name"`
			BadgeID int    `json:"badgeId"`
		} `json:"clan"`
		Cards []struct {
			Name     string `json:"name"`
			Level    int    `json:"level"`
			MaxLevel int    `json:"maxLevel"`
			IconUrls []struct {
				Medium string `json:"medium"`
			} `json:"iconUrls"`
		} `json:"cards"`
	} `json:"team"`
	Opponent []struct {
		Tag              string `json:"tag"`
		Name             string `json:"name"`
		StartingTrophies int    `json:"startingTrophies"`
		TrophyChange     int    `json:"trophyChange"`
		Crowns           int    `json:"crowns"`
		Clan             struct {
			Tag     string `json:"tag"`
			Name    string `json:"name"`
			BadgeID int    `json:"badgeId"`
		} `json:"clan"`
		Cards []struct {
			Name     string `json:"name"`
			Level    int    `json:"level"`
			MaxLevel int    `json:"maxLevel"`
			IconUrls []struct {
				Medium string `json:"medium"`
			} `json:"iconUrls"`
		} `json:"cards"`
	} `json:"opponent"`
}
