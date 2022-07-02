package src

func isLetter(r rune) bool {
	return (r >= 65 && r <= 90) || (r >= 97 && r <= 122)
}
