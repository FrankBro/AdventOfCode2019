package day4

const min = 125730
const max = 579381

func split(number int, count int) []int {
	ds := make([]int, count)
	div := 1
	for i := count - 1; i >= 0; i-- {
		d := number / div % 10
		ds[i] = d
		div *= 10
	}
	return ds
}

func valid(number int) bool {
	ds := split(number, 6)
	var double bool
	var goesDown bool
	for i := range ds {
		if i != len(ds)-1 {
			if ds[i] > ds[i+1] {
				goesDown = true
			}
			if ds[i] == ds[i+1] {
				double = true
			}
		}
	}
	return double && !goesDown
}

func part1() int {
	var count int
	for i := min; i < max; i++ {
		if valid(i) {
			count++
		}
	}
	return count
}

func checkDouble(ds []int, i int) bool {
	switch i {
	case 0:
		return ds[0] == ds[1] && ds[0] != ds[2]
	case 1:
		return ds[1] == ds[2] && ds[1] != ds[0] && ds[1] != ds[3]
	case 2:
		return ds[2] == ds[3] && ds[2] != ds[1] && ds[2] != ds[4]
	case 3:
		return ds[3] == ds[4] && ds[3] != ds[2] && ds[3] != ds[5]
	case 4:
		return ds[4] == ds[5] && ds[4] != ds[3]
	}
	return false
}

func valid2(number int) bool {
	ds := split(number, 6)
	var double bool
	var goesDown bool
	for i := range ds {
		if i != len(ds)-1 {
			if ds[i] > ds[i+1] {
				goesDown = true
			}
		}
		if checkDouble(ds, i) {
			double = true
		}
	}
	return double && !goesDown
}

func part2() int {
	var count int
	for i := min; i < max; i++ {
		if valid2(i) {
			count++
		}
	}
	return count
}
