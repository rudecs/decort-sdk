package validators

func StringInSlice(str string, target []string) bool {
	for _, v := range target {
		if v == str {
			return true
		}
	}
	return false
}
