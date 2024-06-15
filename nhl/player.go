package nhl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Player struct {
	PlayerID            *int                       `json:"playerId,omitempty"`
	IsActive            *bool                      `json:"isActive,omitempty"`
	CurrentTeamId       *int                       `json:"currentTeamId,omitempty"`
	CurrentTeamAbbrev   *string                    `json:"currentTeamAbbrev,omitempty"`
	FullTeamName        *Player_FullTeamName       `json:"fullTeamName,omitempty"`
	FirstName           *Player_FirstName          `json:"firstName,omitempty"`
	LastName            *Player_LastName           `json:"lastName,omitempty"`
	TeamLogo            *string                    `json:"teamLogo,omitempty"`
	SweaterNumber       *int                       `json:"sweaterNumber,omitempty"`
	Position            *string                    `json:"position,omitempty"`
	Headshot            *string                    `json:"headshot,omitempty"`
	HeroImage           *string                    `json:"heroImage,omitempty"`
	HeightInInches      *int                       `json:"heightInInches,omitempty"`
	HeightInCentimeters *int                       `json:"heightInCentimeters,omitempty"`
	WeightInPounds      *int                       `json:"weightInPounds,omitempty"`
	WeightInKilograms   *int                       `json:"weightInKilograms,omitempty"`
	BirthDate           *string                    `json:"birthDate,omitempty"`
	BirthCity           *Player_BirthCity          `json:"birthCity,omitempty"`
	BirthStateProvince  *Player_BirthStateProvince `json:"playerStateProvince,omitempty"`
	BirthCountry        *string                    `json:"birthCountry,omitempty"`
	ShootsCatches       *string                    `json:"shootsCatches,omitempty"`
	DraftDetails        *Player_DraftDetails       `json:"draftDetails,omitempty"`
	PlayerSlug          *string                    `json:"playerSlug,omitempty"`
	InTop100AllTime     *int                       `json:"inTop100AllTime,omitempty"`
	InHHOF              *int                       `json:"inHHOF,omitempty"`
	FeaturedStats       *Player_FeaturedStats      `json:"featuredStats,omitempty"`
	CareerTotals        *Player_CareerTotals       `json:"careerTotals,omitempty"`
	ShopLink            *string                    `json:"shopLink,omitempty"`
	TwitterLink         *string                    `json:"twitterLink,omitempty"`
	WatchLink           *string                    `json:"watchLink,omitempty"`
	// Last5Games _ `json:"last5Games"`
	SeasonTotals []*Player_Total `json:"seasonTotals,omitempty"`
	// Awards _ `json:"awards"`
	// CurrentTeamRoster _ `json:"currentTeamRoster"`
}

type Player_FullTeamName struct {
	Default *string `json:"default,omitempty"`
	FR      *string `json:"fr,omitempty"`
	CS      *string `json:"cs,omitempty"`
	DE      *string `json:"de,omitempty"`
	FI      *string `json:"fi,omitempty"`
	SK      *string `json:"sk,omitempty"`
	SV      *string `json:"sk,omitempty"`
}

func (p *Player) GetTeamName() string {
	return p.GetFullTeamName().GetDefault()
}

type Player_FirstName struct {
	Default *string `json:"default,omitempty"`
}

// func (p *Player) GetFirstName() string {
// 	return p.GetFirstName().GetDefault()
// }

type Player_LastName struct {
	Default *string `json:"default,omitempty"`
}

// func (p *Player) GetLastName() string {
// 	return p.LastName.Default
// }

type Player_BirthCity struct {
	Default *string `json:"default,omitempty"`
}

// func (p *Player) GetBirthCity() string {
// 	return p.BirthCity.Default
// }

type Player_BirthStateProvince struct {
	Default *string `json:"default,omitempty"`
}

// func (p *Player) GetBirthStateProvince() string {
// 	return p.BirthStateProvince.Default
// }

type Player_DraftDetails struct {
	Year        *int    `json:"year,omitempty"`
	TeamAbbrev  *string `json:"teamAbbrev,omitempty"`
	Round       *int    `json:"round,omitempty"`
	PickInRound *int    `json:"pickInRound,omitempty"`
	OverallPick *int    `json:"overallPick,omitempty"`
}

type Player_FeaturedStats struct {
	Season         *int                   `json:"season,omitempty"`
	RegularSesason *Player_FeaturedTotals `json:"regularSeason,omitempty"`
	Playoffs       *Player_FeaturedTotals `json:"playoffs,omitempty"`
}

type Player_FeaturedTotals struct {
	SubSeason *Player_Total `json:"subSeason,omitempty"`
	Career    *Player_Total `json:"career,omitempty"`
}

type Player_CareerTotals struct {
	RegularSeason *Player_Total `json:"regularSeason,omitempty"`
	Playoffs      *Player_Total `json:"playoffs,omitempty"`
}

type Player_Total struct {
	Assists            *int                 `json:"assists,omitempty"`
	AvgToi             *string              `json:"avgToi,omitempty"`
	FaceoffWinningPctg *float64             `json:"faceoffWinningPctg,omitempty"`
	GameTypeId         *int                 `json:"gameTypeId,omitempty"`
	GameWinningGoals   *int                 `json:"gameWinningGoals,omitempty"`
	GamesPlayed        *int                 `json:"gamesPlayed,omitempty"`
	Goals              *int                 `json:"goals,omitempty"`
	LeagueAbbrev       *string              `json:"leagueAbbrev,omitempty"`
	OTGoals            *int                 `json:"otGoals,omitempty"`
	PIM                *int                 `json:"pim,omitempty"`
	PlusMinus          *int                 `json:"plusMinus,omitempty"`
	Points             *int                 `json:"points,omitempty"`
	PowerPlayGoals     *int                 `json:"powerPlayGoals,omitempty"`
	PowerPlayPoints    *int                 `json:"powerPlayPoints,omitempty"`
	Season             *int                 `json:"season,omitempty"`
	Sequence           *int                 `json:"sequence,omitempty"`
	ShootingPctg       *float64             `json:"shootingPctg,omitempty"`
	ShorthandedGoals   *int                 `json:"shorthandedGoals,omitempty"`
	ShorthandedPoints  *int                 `json:"shorthandedPoints,omitempty"`
	Shots              *int                 `json:"shots,omitempty"`
	TeamName           *Player_FullTeamName `json:"teamName,omitempty"`
}

func GetPlayer(id int) (*Player, error) {
	url := fmt.Sprintf("%s/v1/player/%d/landing", baseURL, id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GetPlayer: %v", err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var p Player
	if err := json.Unmarshal(body, &p); err != nil {
		return &p, fmt.Errorf("GetPlayer: %v", err)
	}

	return &p, nil
}
