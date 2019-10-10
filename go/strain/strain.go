package strain

// Ints is an array of ints
type Ints []int

// Lists is a two dimensional array of ints
type Lists [][]int

// Strings is an array of strings
type Strings []string

// Keep retains items in a collection of ints based on a rule defined by the callback function
func (collection Ints) Keep(callback func(int) bool) Ints {
	if collection == nil {
		return nil
	}
	filtered := Ints{}
	for _, item := range collection {
		if callback(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// Discard filters out items from a collection of ints based on a rule defined by the callback function
func (collection Ints) Discard(callback func(int) bool) Ints {
	if collection == nil {
		return nil
	}
	filtered := Ints{}
	for _, item := range collection {
		if !callback(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// Keep retains items in a list based on a rule defined by the callback function
func (collection Lists) Keep(callback func([]int) bool) Lists {
	if collection == nil {
		return nil
	}
	filtered := Lists{}
	for _, item := range collection {
		if callback(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

// Keep retains items in a collection of strings based on a rule defined by the callback function
func (collection Strings) Keep(callback func(string) bool) Strings {
	if collection == nil {
		return nil
	}
	filtered := Strings{}
	for _, item := range collection {
		if callback(item) {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
