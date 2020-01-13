package animenames_test

import (
	"testing"

	"git.wark.io/lib/animenames-go"
)

var parserTests = map[string]*animenames.Anime{
	"[HorribleSubs] Himouto! Umaru-chan - 01 [720p].mkv": &animenames.Anime{
		Title:   "Himouto! Umaru-chan",
		Episode: 1,
		Group:   "HorribleSubs",
	},
	"Himouto! Umaru-chan【01】【GB】【720P】【MP4】": &animenames.Anime{
		Title:   "Himouto! Umaru-chan",
		Episode: 1,
		Group:   "GB",
	},
	"[FFF] Working!!! - 01 [720p][348B33FB].mkv": &animenames.Anime{
		Title:   "Working!!!",
		Episode: 1,
		Group:   "FFF",
		CRC32:   "348B33FB",
	},
	"Fate/Stay Night: Unlimited Blade Works (2015)": &animenames.Anime{
		Title: "Fate/Stay Night: Unlimited Blade Works",
		Year:  2015,
	},
	"[UTW-Mazui-MK] Toaru Majutsu no Index Movie - Endymion no Kiseki [BD 1080p Hi10p Dual Audio-FLAC][9e89d1ac].mkv": &animenames.Anime{
		Title: "Toaru Majutsu no Index Movie - Endymion no Kiseki",
		Group: "UTW-Mazui-MK",
		CRC32: "9e89d1ac",

		IsBD: true,
	},
	"(project-gxs)_Mekakucity_Actors_(10bit_BD_1080p)": &animenames.Anime{
		Title: "Mekakucity Actors",
		Group: "project-gxs",
		IsBD:  true,
	},
	"(project-gxs)_Shimoneta_01_(10bit_720p).mkv": &animenames.Anime{
		Title:   "Shimoneta",
		Episode: 1,
		Group:   "project-gxs",
	},
	"(project-gxs)_Charlotte_-_01_(10bit_720p).mkv": &animenames.Anime{
		Title:   "Charlotte",
		Episode: 1,
		Group:   "project-gxs",
	},
	"[CabbageSubs] Himouto! Umaru-chan - 01 [720p] [549C0C38].mkv": &animenames.Anime{
		Title:   "Himouto! Umaru-chan",
		Episode: 1,
		Group:   "CabbageSubs",
		CRC32:   "549C0C38",
	},
	"[HorribleSubs] Haiyore! Nyaruko-san W - 01-12 [1080p]": &animenames.Anime{
		Title: "Haiyore! Nyaruko-san W",
		Group: "HorribleSubs",
		Batch: &animenames.Batch{
			Start: 1,
			End:   12,
		},
	},
	"[Glitch] Haiyore! Nyaruko-san F - OVA (BD 1280x720 x264 AAC).mkv": &animenames.Anime{
		Title: "Haiyore! Nyaruko-san F",
		Group: "Glitch",
		IsOVA: true,
		IsBD:  true,
	},
	"[MD] Haiyore! Nyaruko-san W - Vol.01 (1920x1080 Blu-ray FLAC)": &animenames.Anime{
		Title:  "Haiyore! Nyaruko-san W",
		Volume: 1,
		Group:  "MD",
		IsBD:   true,
	},
	"[GS] Hibike! Euphonium Vol.1 (BD 1080p 10bit FLAC)": &animenames.Anime{
		Title:  "Hibike! Euphonium",
		Volume: 1,
		Group:  "GS",
		IsBD:   true,
	},
	"[Pn8] High School DxD BorN S03E12 Any Time, For All Time! [Simuldub] [720p][10bit] [Dual]": &animenames.Anime{
		Title:   "High School DxD BorN",
		Episode: 12,
		Season:  3,
		Group:   "Pn8",
	},
	"[EveTaku] Nisekoi 01 [720p-Hi10P] [8FEC89B6].mkv": &animenames.Anime{
		Title:   "Nisekoi",
		Episode: 1,
		Group:   "EveTaku",
		CRC32:   "8FEC89B6",
	},
	"[FFF] Nisekoi S2 - 01 [E0D0C713].mkv": &animenames.Anime{
		Title:   "Nisekoi",
		Episode: 1,
		Season:  2,
		Group:   "FFF",
		CRC32:   "E0D0C713",
	},
	"[FuniOCR] Hetalia - The World Twinkle - 03.ass": &animenames.Anime{
		Title:   "Hetalia - The World Twinkle",
		Episode: 3,
		Group:   "FuniOCR",
	},
	"High School DxD Born English Dub Uncensored 1-12 720p Complete": &animenames.Anime{
		Title: "High School DxD Born",
		Batch: &animenames.Batch{
			Start: 1,
			End:   12,
		},
	},
	"(project-gxs)_Dragon_Ball_Super_-_003v3_(10bit_720p).mkv": &animenames.Anime{
		Title:   "Dragon Ball Super",
		Episode: 3,
		Group:   "project-gxs",
	},
	"[DeadFish] Working!! - S1 & S2 [BD][720p][MP4][AAC]": &animenames.Anime{
		Title: "Working!!",
		Group: "DeadFish",
		IsBD:  true,
	},
	"[Senketsu Subs] Joukamachi no Dandelion - 03.ass (v2)": &animenames.Anime{
		Title:   "Joukamachi no Dandelion",
		Episode: 3,
		Group:   "Senketsu Subs",
	},
	"[Senketsu Rips] Nagato Yuki-chan no Shoushitsu - 16.ass (END)": &animenames.Anime{
		Title:   "Nagato Yuki-chan no Shoushitsu",
		Episode: 16,
		Group:   "Senketsu Rips",
	},
	"[sushit] GATE - Thus, the Self Defense Force Fought There - 03 (720p) [FD2598E7].mkv": &animenames.Anime{
		Title:   "GATE - Thus, the Self Defense Force Fought There",
		Episode: 3,
		Group:   "sushit",
		CRC32:   "FD2598E7",
	},
	"[Doki] Kore wa Zombie Desu ka - 01 (1280x720 HEVC BD AAC) [2A6C448F]_Track02.ass": &animenames.Anime{
		Title:   "Kore wa Zombie Desu ka",
		Episode: 1,
		Group:   "Doki",
		CRC32:   "2A6C448F",
		IsBD:    true,
	},
	"Date a Live (2013) [Doki][1920x1080 Hi10P BD FLAC]": &animenames.Anime{
		Title: "Date a Live",
		Year:  2013,
		Group: "Doki",
		IsBD:  true,
	},
	"Toradora!": &animenames.Anime{
		Title: "Toradora!",
	},
	"[FFF] Saenai Heroine no Sodatekata - 00v2 [366ABCCA].mkv": &animenames.Anime{
		Title:   "Saenai Heroine no Sodatekata",
		Episode: 0,
		Group:   "FFF",
		CRC32:   "366ABCCA",
	},
	"[HorribleSubs] Gochuumon wa Usagi Desu ka S2 - 01 [720p].mkv": &animenames.Anime{
		Title:   "Gochuumon wa Usagi Desu ka",
		Episode: 1,
		Season:  2,
		Group:   "HorribleSubs",
	},
	"[DeadFish] Himouto! Umaru-chan S - 04 - Special [720p][AAC].mp4": &animenames.Anime{
		Title:   "Himouto! Umaru-chan S",
		Episode: 4,
		Group:   "DeadFish",
	},
	"[SabiShin] Mondaiji-tachi ga Isekai kara Kuru Sou Desu yo BD -Completa-": &animenames.Anime{
		Title: "Mondaiji-tachi ga Isekai kara Kuru Sou Desu yo",
		Group: "SabiShin",
		IsBD:  true,
	},
	"[PCNet] Hugtto Pretty Cure - 01 [BD 720p] [2D3B6393].mkv": &animenames.Anime{
		Title:   "Hugtto Pretty Cure",
		Group:   "PCNet",
		Episode: 1,
		IsBD:    true,
		CRC32:   "2D3B6393",
	},
}

func TestParse(t *testing.T) {
	for name, expected := range parserTests {
		result, err := animenames.Parse(name)
		if err != nil {
			t.Fatal(err)
		}
		if result == nil {
			t.Fatal("expecting anime")
		}

		if result.Title != expected.Title {
			t.Fatalf("expecting Anime.Title of %#v to be %#v (got %#v)", name, expected.Title, result.Title)
		}

		if result.Year != expected.Year {
			t.Fatalf("expecting Anime.Year of %#v to be %#v (got %#v)", name, expected.Year, result.Year)
		}

		if result.Episode != expected.Episode {
			t.Fatalf("expecting Anime.Episode of %#v to be `%d` (got `%d`)", name, expected.Episode, result.Episode)
		}

		if result.Season != expected.Season {
			t.Fatalf("expecting Anime.Season of %#v to be %d (got %d)", name, expected.Season, result.Season)
		}

		if result.Volume != expected.Volume {
			t.Fatalf("expecting Anime.Volume of %#v to be `%d` (got `%d`)", name, expected.Volume, result.Volume)
		}

		if result.Group != expected.Group {
			t.Fatalf("expecting Anime.Group of %#v to be %#v (got %#v)", name, expected.Group, result.Group)
		}

		if result.CRC32 != expected.CRC32 {
			t.Fatalf("expecting Anime.CRC32 of %#v to be %#v (got %#v)", name, expected.CRC32, result.CRC32)
		}

		if result.IsOVA != expected.IsOVA {
			t.Fatalf("expecting Anime.IsOVA of %#v to be %#v (got %#v)", name, expected.IsOVA, result.IsOVA)
		}

		if result.IsBD != expected.IsBD {
			t.Fatalf("expecting Anime.IsBD of %#v to be %#v (got %#v)", name, expected.IsBD, result.IsBD)
		}

		if expected.Batch != nil {
			if result.Batch.Start != expected.Batch.Start {
				t.Fatalf("expecting Anime.Batch.Start of %#v to be %#v (got %#v)", name, expected.Batch.Start, result.Batch.Start)
			}

			if result.Batch.End != expected.Batch.End {
				t.Fatalf("expecting Anime.Batch.End of %#v to be %#v (got %#v)", name, expected.Batch.End, result.Batch.End)
			}
		}
	}
}
