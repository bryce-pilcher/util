package slice

/*
   This function deep copies a single dimensional slice from the given slice to a new one.
*/
func DeepCopy(l []interface{}) (n []interface{}) {
	n = make([]interface{}, len(l))
	for i := range l {
		n[i] = l[i]
	}
	return
}
