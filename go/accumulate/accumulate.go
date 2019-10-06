// Package accumulate provides an acculumator function, which given a collection and an operation to perform on each element of the collection,
// returns a new collection containing the result of applying that operation to each element of the input collection.
package accumulate

// Accumulate applies the callback to every item in the collection
func Accumulate(collection []string, callback func(string) string) []string {
	for key, word := range collection {
		collection[key] = callback(word)
	}
	return collection
}
