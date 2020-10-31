package strand

// ToRNA returns a RNA complement of a given dna
func ToRNA(dna string) string {

	size := len(dna)

	if size == 0 {
		return ""
	}

	transcription := ""

	for i := 0; i < size; i++ {
		switch dna[i] {
		case 'A':
			transcription = transcription + "U"
		case 'T':
			transcription = transcription + "A"
		case 'C':
			transcription = transcription + "G"
		case 'G':
			transcription = transcription + "C"
		}
	}
	return transcription
}
