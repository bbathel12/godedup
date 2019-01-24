package main

type stringFilter func(str string) bool

func main() {

}

func dedup(names []string) (uniq []string) {

	if null(names) {
		uniq = []string{}
	} else if dupe(
		first(names),
		rest(names),
	) {
		filterFunc := func(str string) bool {
			if str == first(names) {
				return false
			}
			return true
		}
		uniq = dedup(filter(filterFunc, rest(names)))
	} else {
		uniq = append(uniq, first(names))
		uniq = append(uniq, dedup(rest(names))...)
	}
	return
}

func dupe(name string, names []string) bool {
	if len(names) <= 0 {
		return false
	} else if name == names[0] {
		return true
	} else {
		return (dupe(name, names[1:len(names)]))
	}

}

func null(hasLength []string) bool {
	if len(hasLength) <= 0 {
		return true
	} else {
		return false
	}
}

func first(firstable []string) string {
	return firstable[0]
}

func rest(restable []string) []string {
	if len(restable) > 1 {
		return restable[1:len(restable)]
	} else {
		return []string{}
	}
}

func filter(fun stringFilter, list []string) []string {
	var carry []string
	for _, v := range list {
		if fun(v) {
			carry = append(carry, v)
		}
	}

	return carry
}
