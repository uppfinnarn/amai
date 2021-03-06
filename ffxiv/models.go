package ffxiv

type ClassInfo struct {
	Level int `json:"level"`
	ExpAt int `json:"exp_at"`
	ExpOf int `json:"exp_of"`
}

// A FFXIV character.
type FFXIVCharacter struct {
	ID           uint64 `json:"id"`
	Name         string `json:"name"`
	Title        string `json:"title"`
	Server       string `json:"server"`
	Race         string `json:"race"`
	Clan         string `json:"clan"`
	Gender       string `json:"gender"`
	BirthDay     int    `json:"birth_day"`
	BirthMonth   int    `json:"birth_month"`
	Guardian     string `json:"guardian"`
	GrandCompany struct {
		Name string `json:"name"`
		Rank int    `json:"rank"`
	} `json:"grand_company"`
	FreeCompany struct {
		ID   uint64 `json:"id"`
		Name string `json:"name"`
	} `json:"free_company"`
	Stats   map[string]int       `json:"stats"`
	Classes map[string]ClassInfo `json:"classes"`
	Minions []string             `json:"minions"`
	Mounts  []string             `json:"mounts"`
}

// "Focus" information for a Free Company.
type FCFocus struct {
	RolePlaying bool `json:"role_playing"`
	Leveling    bool `json:"leveling"`
	Casual      bool `json:"casual"`
	Hardcore    bool `json:"hardcore"`
	Dungeons    bool `json:"dungeons"`
	Guildhests  bool `json:"guildhests"`
	Trials      bool `json:"trials"`
	Raids       bool `json:"raids"`
	PvP         bool `json:"pvp"`
}

// "Seeking" information for a Free Company.
type FCSeeking struct {
	Tank     bool `json:"tank"`
	Healer   bool `json:"healer"`
	DPS      bool `json:"dps"`
	Crafter  bool `json:"crafter"`
	Gatherer bool `json:"gatherer"`
}

// Estate information for a Free Company
type FCEstate struct {
	Name     string `json:"name"`
	Plot     int    `json:"plot"`
	Ward     int    `json:"ward"`
	District string `json:"district"`
	Size     int    `json:"size"`
	Greeting string `json:"greeting"`
}

// A Free Company.
type FFXIVFreeCompany struct {
	ID           uint64    `json:"id"`
	Name         string    `json:"name"`
	Tag          string    `json:"tag"`
	Description  string    `json:"description"`
	Server       string    `json:"server_name"`
	GrandCompany string    `json:"grand_company"`
	Rank         int       `json:"rank"`
	Focus        FCFocus   `json:"focus"`
	Seeking      FCSeeking `json:"seeking"`
	Recruiting   bool      `json:"recruiting"`
	Estate       FCEstate  `json:"estate"`
}
