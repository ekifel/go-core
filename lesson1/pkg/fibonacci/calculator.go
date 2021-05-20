package fibonacci

func Calculate(numb int) int {
	f1, f2 := 0, 1

	if numb == 0 {
		return 0
	}
	if numb == 1 {
		return 1
	}

	for i := 1; i < numb; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}
