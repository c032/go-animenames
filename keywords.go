package animenames

var aliases = map[string]string{
	"blu-ray": "bd",
	"bluray":  "bd",
	"h.264":   "h264",
	"x264":    "h264",
}

var resolutions = []string{
	"1080p",
	"360p",
	"480p",
	"720p",
}

var quality = []string{
	"10bit",
	"8bit",
	"bd",
	"dvd",
	"tv",
}

var videoCodecs = []string{
	"h264",
	"hevc",
}

var audioCodecs = []string{
	"aac",
	"ac3",
	"flac",
	"opus",
}

var extensions = []string{
	"ass",
	"mkv",
	"mp4",
}

var otherProperties = []string{
	"censored",
	"complete",
	"dual",
	"dub",
	"english",
	"simuldub",
	"uncensored",
}
