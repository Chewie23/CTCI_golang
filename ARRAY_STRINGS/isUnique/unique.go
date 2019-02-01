package unique

func IsUnique(my_str string) bool {
	char_count := make(map[rune]int)
	for _, c := range my_str {

		//This fancy smancy one liner is doing two things.
		//1) Initializing the two variables of "_" and "ok"
		//		"_" == value of map[key], ok == bool
		//2) Evaluating if the value exists, via "ok"
		//		It's a little weird, but essentially it is abbreviating the following:
		//		val, ok := map[key]
		//		if ok {
		//			do cool stuff
		//		}
		//	The "ok" captures the bool if it got something or not
		if _, ok := char_count[c]; ok {
			return false
		} //if

		char_count[c] = 1

	} //for
	return true
} //IsUnique
