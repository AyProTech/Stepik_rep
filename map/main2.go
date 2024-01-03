for i, _ := range cityPopulation {
	islnMap := false
	for _, el := range groupCity[100] {
		if i == el {
			islnMap = true
		}

	}
	if islnMap != true {
		delete(cityPopulation, i)
	}
}