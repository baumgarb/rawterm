package rawterm

func removeCharAt(xr []rune, pos int) []rune {
	xr2 := make([]rune, len(xr)-1)
	for i, j := 0, 0; i < len(xr); i++ {
		if i == pos {
			continue
		}
		xr2[j] = xr[i]
		j++
	}
	return xr2
}

func insertCharAt(xr []rune, pos int, r rune) []rune {
	xr2 := make([]rune, len(xr)+1)
	for i, j := 0, 0; i < len(xr2); i++ {
		if i == pos {
			xr2[i] = r
			continue
		}
		xr2[i] = xr[j]
		j++
	}
	return xr2
}
