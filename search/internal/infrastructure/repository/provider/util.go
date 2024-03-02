package provider

func isError(status int) bool {
	if status != 200 {
		return true
	}
	return false
}
