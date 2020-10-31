package proverb

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {

	size := len(rhyme)
	proverb := make([]string, size)
	if size == 0 {
		return proverb
	}

	for i := 0; i < size; i++ {
		if i == size-1 {
			proverb[i] = "And all for the want of a " + rhyme[0] + "."
		} else {
			proverb[i] = "For want of a " + rhyme[i] + " the " + rhyme[i+1] + " was lost."
		}
	}

	return proverb
}
