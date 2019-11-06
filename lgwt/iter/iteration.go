package iteration

func Repeat(character string, n int) string {
	repeated := ""
	
	for i := 0; i < n; i++ {
		repeated += character
	}

	return repeated
}