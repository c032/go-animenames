package animenames

// Anime contains information about an anime file or directory.
type Anime struct {
	Title   string
	Year    int
	Episode int
	Season  int // e.g. `2` in "Nisekoi S2"
	Volume  int
	Batch   *Batch
	Group   string
	CRC32   string

	IsOVA       bool
	IsBD        bool
	HasSpecials bool
}

// Batch describes a batch.
type Batch struct {
	Start int
	End   int
}
