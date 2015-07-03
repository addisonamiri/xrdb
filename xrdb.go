package xrdb

import (
	"bytes"
	"os/exec"
	"strings"
)

var parseDict = map[string]string{
	"bg":            "background",
	"fg":            "foreground",
	"0":             "color0",
	"1":             "color1",
	"2":             "color2",
	"3":             "color3",
	"4":             "color4",
	"5":             "color5",
	"6":             "color6",
	"7":             "color7",
	"8":             "color8",
	"9":             "color9",
	"10":            "color10",
	"11":            "color11",
	"12":            "color12",
	"13":            "color13",
	"14":            "color14",
	"15":            "color15",
	"black":         "color0",
	"red":           "color1",
	"green":         "color2",
	"yellow":        "color3",
	"blue":          "color4",
	"magenta":       "color5",
	"cyan":          "color6",
	"white":         "color7",
	"brightblack":   "color8",
	"brightred":     "color9",
	"brightgreen":   "color10",
	"brightyellow":  "color11",
	"brightblue":    "color12",
	"brightmagenta": "color13",
	"brightcyan":    "color14",
	"brightwhite":   "color15",

	"darkblack":   "color0",
	"darkred":     "color1",
	"darkgreen":   "color2",
	"darkyellow":  "color3",
	"darkblue":    "color4",
	"darkmagenta": "color5",
	"darkcyan":    "color6",
	"darkwhite":   "color7",

	"lightblack":   "color8",
	"lightred":     "color9",
	"lightgreen":   "color10",
	"lightyellow":  "color11",
	"lightblue":    "color12",
	"lightmagenta": "color13",
	"lightcyan":    "color14",
	"lightwhite":   "color15",

	"grey":       "color8",
	"darkgrey":   "color8",
	"brightgrey": "color7",
	"lightgrey":  "color7",

	"bold":						"boldfont",
	"italic":					"italicfont",
	"bolditalic":			"bolditalicfont",
	"italicbold":     "bolditalicfont",

	"color8":  "color0",
	"color9":  "color1",
	"color10": "color2",
	"color11": "color3",
	"color12": "color4",
	"color13": "color5",
	"color14": "color6",
	"color15": "color7",

	"boldfont":       "font",
	"italicfont":     "boldfont",
	"bolditalicfont": "boldfont",
}

func Get(query string) string {
	db := GetAll()

	for ; db[query] == "" && parseDict[query] != ""; {
		query = parseDict[query]
	}

	return db[query]
}

func GetAll() map[string]string {
	out := make(map[string]string)
	buffer := new(bytes.Buffer)
	xrdb := exec.Command("xrdb", "-query")
	stdout, err := xrdb.StdoutPipe()

	if err != nil {
		return make(map[string]string)
	}

	if err := xrdb.Start(); err != nil {
		return make(map[string]string)
	}

	buffer.ReadFrom(stdout)
	raw := buffer.String()

	list := strings.Split(raw, "\n")

	for _, line := range list {
		tmp := strings.Split(line, "\t")
		if len(tmp) > 1 && strings.HasPrefix(tmp[0], "*.") {
			out[strings.ToLower(tmp[0][2:len(tmp[0])-1])] = tmp[1]
		}
	}

	return out
}

func Query(query string) string {
	return GetAll()[query]
}
