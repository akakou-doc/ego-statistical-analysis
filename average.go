package main

func average() int {
	result := 0

	for _, v := range users {
		if v.weight == 0 {
			return 0
		}

		result += v.weight
	}

	return result / len(users)
}
