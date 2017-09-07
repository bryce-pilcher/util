package slice

//Swap position a and postion b in a slice
func Swap(l []interface{}, a int, b int) {
	tmp := l[a]
	l[a] = l[b]
	l[b] = tmp
}
