package main

func square(x int) int{
	return x * x
}

func concat(s1, s2 string) string{
	var s []rune
	for _, c := range s1{
		s = append(s, c)
	}
	for _, c := range s2{
		s = append(s, c)
	}
	return string(s)
}