package strand

import "strings"

const (
	DNA_SEQUENCE string = "GCTA"
	RNA_SEQUENCE string = "CGAU" // Complement to DNA
)

func ToRNA(dna string) (result string) {
	for _, v := range dna {
		result += string(RNA_SEQUENCE[strings.Index(DNA_SEQUENCE, string(v))])
	}
	return result
}
