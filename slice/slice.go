package slice

// Tar lengde og kapasitet som b og lager en ny tom []byte slice
func AllocateVar(b int) []byte {
	var s []byte
	s = make([]byte, b, b)
	return s
}

// Tar lengde og kapasitet som b og lager en ny tom []byte slice
func AllocateMake(b int) []byte {
	s := make([]byte, b, b)
	return s
}

// Reslice tar en bit av slicen og returner en ny slice
func Reslice(slc []byte, lidx int, uidx int) []byte {
	s := CopySlice(slc)
	s = s[lidx:uidx]
	return s
}

func CopySlice(slc []byte) []byte {
	s := AllocateMake(len(slc))
	copy(s, slc)
	return s
}