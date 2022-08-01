package flags

var LETTERS_EMOJIS = map[string]string{
	"A": "\U0001F1E6",
	"B": "\U0001F1E7",
	"C": "\U0001F1E8",
	"D": "\U0001F1E9",
	"E": "\U0001F1EA",
	"F": "\U0001F1EB",
	"G": "\U0001F1EC",
	"H": "\U0001F1ED",
	"I": "\U0001F1EE",
	"J": "\U0001F1EF",
	"K": "\U0001F1F0",
	"L": "\U0001F1F1",
	"M": "\U0001F1F2",
	"N": "\U0001F1F3",
	"O": "\U0001F1F4",
	"P": "\U0001F1F5",
	"Q": "\U0001F1F6",
	"R": "\U0001F1F7",
	"S": "\U0001F1F8",
	"T": "\U0001F1F9",
	"U": "\U0001F1FA",
	"V": "\U0001F1FB",
	"W": "\U0001F1FC",
	"X": "\U0001F1FD",
	"Y": "\U0001F1FE",
	"Z": "\U0001F1FF",
}

func GetFlag(code string) (flag string) {
	for _, letter := range code {
		flag += LETTERS_EMOJIS[string(letter)]
	}

	return flag
}
