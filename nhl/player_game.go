package nhl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type gamelog_resp struct {
	SeasonId           int `json:"seasonId"`
	GameTypeId         int `json:"gameTypeId"`
	PlayerStatsSeasons []struct {
		Season    int   `json:"season"`
		GameTypes []int `json:"gameTypes"`
	} `json:"playerStatsSeasons"`
	GameLog []*Player_Game `json:"gameLog"`
}

type Player_Game struct {
	GameID             *int                    `json:"gameId,omitempty"`
	TeamAbbrev         *string                 `json:"teamAbbrev,omitempty"`
	HomeRoadFlag       *string                 `json:"homeRoadFlag,omitempty"`
	GameDate           *string                 `json:"gameDate,omitempty"`
	Goals              *int                    `json:"goals,omitempty"`
	Assists            *int                    `json:"assists,omitempty"`
	CommonName         *Player_Game_CommonName `json:"commonName,omitempty"`
	OpponentCommonName *Player_Game_CommonName `json:"opponentCommonName,omitempty"`
	OpponentAbbrev     *string                 `json:"opponentAbbrev,omitempty"`
	PIM                *int                    `json:"pim,omitempty"`
	TOI                *string                 `json:"toi,omitempty"`

	// Skater specififc fields

	Points            *int `json:"points,omitempty"`
	PlusMinus         *int `json:"plusMinus,omitempty"`
	PowerPlayGoals    *int `json:"powerPlayGoals,omitempty"`
	PowerPlayPoints   *int `json:"powerPlayPoints,omitempty"`
	GameWinningGoals  *int `json:"gameWinningGoals,omitempty"`
	OTGoals           *int `json:"otGoals,omitempty"`
	Shots             *int `json:"shots,omitempty"`
	Shifts            *int `json:"shifts,omitempty"`
	ShorthandedGoals  *int `json:"shorthandedGoals,omitempty"`
	ShorthandedPoints *int `json:"shorthandedPoints,omitempty"`

	// Goalie specific fields

	GamesStarted *int     `json:"gamesStarted,omitempty"`
	Decision     *string  `json:"decision,omitempty"`
	ShotsAgainst *int     `json:"shotsAgainst,omitempty"`
	GoalsAgainst *int     `json:"goalsAgainst,omitempty"`
	SavePctg     *float64 `json:"savePctg,omitempty"`
	Shutouts     *int     `json:"shutouts,omitempty"`
}

type Player_Game_CommonName struct {
	Default *string `json:"default,omitempty"`
}

func (p *Player) GetGameLog(season, seasonType int) ([]*Player_Game, error) {
	return GetGameLog(p.GetPlayerId(), season, seasonType)
}

func GetGameLog(playerId, season, seasonType int) ([]*Player_Game, error) {
	url := fmt.Sprintf("%s/player/%d/game-log/%d/%d",
		baseURL,
		playerId,
		season,
		seasonType,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GetGameLog: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var gl_resp gamelog_resp
	if err := json.Unmarshal(body, &gl_resp); err != nil {
		return nil, fmt.Errorf("GetGameLog: %v", err)
	}

	return gl_resp.GameLog, nil
}
