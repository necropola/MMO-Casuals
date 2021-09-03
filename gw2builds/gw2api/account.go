package gw2api

import (
	"time"
)

type Account struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Age          time.Duration `json:"age"`
	World        int           `json:"world"`
	Guilds       []string      `json:"guilds"`
	GuildLeader  []string      `json:"guild_leader"`
	Created      time.Time     `json:"created"`
	Access       []string      `json:"access"`
	Commander    bool          `json:"commander"`
	FractalLevel int           `json:"fractal_level"`
	DailyAP      int           `json:"daily_ap"`
	MonthlyAP    int           `json:"monthly_ap"`
	WvWRank      int           `json:"wvw_rank"`
}

func (api *GW2API) Account() (account Account, err error) {
	err = api.fetch("/v2/account", &account)
	return
}
