package ffxiv

// A FFXIV character.
type FFXIVCharacter struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Title        string `json:"title"`
	ServerName   string `json:"server_name"`
	Race         string `json:"race"`
	Clan         string `json:"clan"`
	Gender       string `json:"gender"`
	BirthMonth   int    `json:"birth_month"`
	BirthDay     int    `json:"birth_day"`
	Guardian     string `json:"guardian"`
}
