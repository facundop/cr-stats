package domain

// Player is a representation of a player
type Player struct {
	Tag      string `json:"tag"`
	Name     string `json:"name"`
	ExpLevel int    `json:"expLevel"`
	Trophies int    `json:"trophies"`
	Arena    struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"arena"`
	BestTrophies          int    `json:"bestTrophies"`
	Wins                  int    `json:"wins"`
	Losses                int    `json:"losses"`
	BattleCount           int    `json:"battleCount"`
	ThreeCrownWins        int    `json:"threeCrownWins"`
	ChallengeCardsWon     int    `json:"challengeCardsWon"`
	ChallengeMaxWins      int    `json:"challengeMaxWins"`
	TournamentCardsWon    int    `json:"tournamentCardsWon"`
	TournamentBattleCount int    `json:"tournamentBattleCount"`
	Role                  string `json:"role"`
	Donations             int    `json:"donations"`
	DonationsReceived     int    `json:"donationsReceived"`
	TotalDonations        int    `json:"totalDonations"`
	WarDayWins            int    `json:"warDayWins"`
	ClanCardsCollected    int    `json:"clanCardsCollected"`
	Clan                  struct {
		Tag     string `json:"tag"`
		Name    string `json:"name"`
		BadgeID int    `json:"badgeId"`
	} `json:"clan"`
	LeagueStatistics struct {
		CurrentSeason struct {
			ID           string `json:"id"`
			Trophies     int    `json:"trophies"`
			BestTrophies int    `json:"bestTrophies"`
		} `json:"currentSeason"`
		PreviousSeason struct {
			ID           string `json:"id"`
			Trophies     int    `json:"trophies"`
			BestTrophies int    `json:"bestTrophies"`
		} `json:"previousSeason"`
		BestSeason struct {
			ID           string `json:"id"`
			Trophies     int    `json:"trophies"`
			BestTrophies int    `json:"bestTrophies"`
		} `json:"bestSeason"`
	} `json:"leagueStatistics"`
	Achievements []struct {
		Name   string `json:"name"`
		Stars  int    `json:"stars"`
		Value  int    `json:"value"`
		Target int    `json:"target"`
		Info   string `json:"info"`
	} `json:"achievements"`
	Cards []struct {
		Name     string `json:"name"`
		Level    int    `json:"level"`
		MaxLevel int    `json:"maxLevel"`
		Count    int    `json:"count"`
		IconUrls struct {
			Medium string `json:"medium"`
		} `json:"iconUrls"`
	} `json:"cards"`
	CurrentFavouriteCard struct {
		Name     string `json:"name"`
		ID       int    `json:"id"`
		MaxLevel int    `json:"maxLevel"`
		IconUrls struct {
			Medium string `json:"medium"`
		} `json:"iconUrls"`
	} `json:"currentFavouriteCard"`
}
