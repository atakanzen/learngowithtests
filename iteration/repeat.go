package iteration

func Repeat(txt string, count int) string {
	if count == 0 {
		count = 5
	}
	var repeated string
	for i := 0; i < count; i++ {
		repeated += txt
	}
	return repeated
}
