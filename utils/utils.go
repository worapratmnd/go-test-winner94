package utils

func ReturnStringPointer(s string) *string {
	v := s
	return &v
}

func ReturnIntPointer(s int) *int {
	v := s
	return &v
}
