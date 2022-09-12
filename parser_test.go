package animenames_test

import (
	"testing"

	"github.com/c032/go-animenames"
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
	"[Anime Time] Kaguya-sama wa Kokurasetai! (Love Is War) S2 - 01 [Dual Audio][HEVC 10bit x265][AAC].mkv": &animenames.Anime{
		Title:   "Kaguya-sama wa Kokurasetai! (Love Is War)",
		Group:   "Anime Time",
		Episode: 1,
		Season:  2,
	},
	"[EMBER] Kanojo mo Kanojo S01E03 [1080p] [HEVC WEBRip] (Girlfriend, Girlfriend)": &animenames.Anime{
		Title:   "Kanojo mo Kanojo",
		Group:   "EMBER",
		Episode: 3,
		Season:  1,
	},
	"[Fix-Fontsizecolor] Strike the Blood IV - 12 (BD 1920x1080 x265 FLAC)": &animenames.Anime{
		Title:   "Strike the Blood IV",
		Group:   "Fix-Fontsizecolor",
		Episode: 12,
		IsBD:    true,
	},
	"[CBM] Uzaki-chan Wants to Hang Out! 1-12 Complete (Dual Audio) [BDRip 1080p x265 10bit]": &animenames.Anime{
		Title: "Uzaki-chan Wants to Hang Out!",
		Group: "CBM",
		IsBD:  true,
		Batch: &animenames.Batch{
			Start: 1,
			End:   12,
		},
	},
	"[ΑΩ] Uzaki-chan wa Asobitai! Vol.3 (BD Remux 1920x1080 AVC PCM)": &animenames.Anime{
		Title:  "Uzaki-chan wa Asobitai!",
		Group:  "ΑΩ",
		Volume: 3,
		IsBD:   true,
	},
	"[Pookie] Flying Witch + SPs [BD 1920x1080 x264 FLAC] [Dual-Audio]": &animenames.Anime{
		Title:       "Flying Witch",
		Group:       "Pookie",
		IsBD:        true,
		HasSpecials: true,
	},
	"[RH] Flying Witch + Specials [Dual Audio] [BDRip] [Hi10] [1080p] [FLAC]": &animenames.Anime{
		Title:       "Flying Witch",
		Group:       "RH",
		IsBD:        true,
		HasSpecials: true,
	},
	"[Golumpa] Flying Witch + Specials v2 [Dual Audio] [BDRip] [1080p] [10-bit] [MKV]": &animenames.Anime{
		Title:       "Flying Witch",
		Group:       "Golumpa",
		IsBD:        true,
		HasSpecials: true,
	},
	"[ASW] 86 - Eighty Six [1080p HEVC x265 10Bit][AAC] (Batch)": &animenames.Anime{
		Title: "86 - Eighty Six",
		Group: "ASW",
	},
	"[SSA] Eighty Six Season 1 (1-11) [1080p][Batch]": &animenames.Anime{
		Title: "Eighty Six Season 1",
		Group: "SSA",
		Batch: &animenames.Batch{
			Start: 1,
			End:   11,
		},
	},
	"86": &animenames.Anime{
		Title: "86",
	},
	"86 - 01": &animenames.Anime{
		Title:   "86",
		Episode: 1,
	},
	"[Kantai] Eighty Six (86) - 23 (1920x1080 AC3) [05BD70FE].mkv": &animenames.Anime{
		Title:   "Eighty Six (86)",
		Group:   "Kantai",
		Episode: 23,
		CRC32:   "05BD70FE",
	},
}

func TestParse(t *testing.T) {
	for name, expectedAnime := range parserTests {
		gotAnime, err := animenames.Parse(name)
		if err != nil {
			t.Fatal(err)
		}

		if gotAnime.Title != expectedAnime.Title {
			t.Errorf("animenames.Parse(%#v).Title = %#v; expected %#v", name, gotAnime.Title, expectedAnime.Title)
		}

		if gotAnime.Year != expectedAnime.Year {
			t.Errorf("animenames.Parse(%#v).Year = %#v; expected %#v", name, gotAnime.Year, expectedAnime.Year)
		}

		if gotAnime.Episode != expectedAnime.Episode {
			t.Errorf("animenames.Parse(%#v).Episode = %#v; expected %#v", name, gotAnime.Episode, expectedAnime.Episode)
		}

		if gotAnime.Season != expectedAnime.Season {
			t.Errorf("animenames.Parse(%#v).Season = %#v; expected %#v", name, gotAnime.Season, expectedAnime.Season)
		}

		if gotAnime.Volume != expectedAnime.Volume {
			t.Errorf("animenames.Parse(%#v).Volume = %#v; expected %#v", name, gotAnime.Volume, expectedAnime.Volume)
		}

		if gotAnime.Group != expectedAnime.Group {
			t.Errorf("animenames.Parse(%#v).Group = %#v; expected %#v", name, gotAnime.Group, expectedAnime.Group)
		}

		if gotAnime.CRC32 != expectedAnime.CRC32 {
			t.Errorf("animenames.Parse(%#v).CRC32 = %#v; expected %#v", name, gotAnime.CRC32, expectedAnime.CRC32)
		}

		if gotAnime.IsOVA != expectedAnime.IsOVA {
			t.Errorf("animenames.Parse(%#v).IsOVA = %#v; expected %#v", name, gotAnime.IsOVA, expectedAnime.IsOVA)
		}

		if gotAnime.IsBD != expectedAnime.IsBD {
			t.Errorf("animenames.Parse(%#v).IsBD = %#v; expected %#v", name, gotAnime.IsBD, expectedAnime.IsBD)
		}

		if gotAnime.HasSpecials != expectedAnime.HasSpecials {
			t.Errorf("animenames.Parse(%#v).HasSpecials = %#v; expected %#v", name, gotAnime.HasSpecials, expectedAnime.HasSpecials)
		}

		if expectedAnime.Batch != nil {
			if gotAnime.Batch.Start != expectedAnime.Batch.Start {
				t.Errorf("expecting Anime.Batch.Start of %#v to be %#v (got %#v)", name, expectedAnime.Batch.Start, gotAnime.Batch.Start)
			}

			if gotAnime.Batch.End != expectedAnime.Batch.End {
				t.Errorf("expecting Anime.Batch.End of %#v to be %#v (got %#v)", name, expectedAnime.Batch.End, gotAnime.Batch.End)
			}
		}
	}
}
