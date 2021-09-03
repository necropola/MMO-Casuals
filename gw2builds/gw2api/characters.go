package gw2api

func (api *GW2API) Characters() (chars []string, err error) {
	err = api.fetch("/v2/characters", &chars)
	return
}

type CharacterCore struct {
	Name       string `json:"name"`
	Race       string `json:"race"`
	Gender     string `json:"gender"`
	Profession string `json:"profession"`
	Level      int    `json:"level"`
	GuildID    string `json:"guild"`
	Age        int    `json:"age"`
	Created    string `json:"created"`
	Deaths     int    `json:"deaths"`
	Title      int    `json:"title"`
}

func (api *GW2API) CharacterCore(character string) (core CharacterCore, err error) {
	err = api.fetch("/v2/characters/"+character+"/core", &core)
	return
}
