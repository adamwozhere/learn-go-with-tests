package main

// a `map` is just a key and a value, similar to an array but indexed by a key.
// declared as: `map[key type]value type`.

type Dictionary map[string]string

// create a global error variable we can use in the test,
// means that we can also change the error text and the test will still work
// var (
// 	ErrNotFound   = errors.New("could not find the word you were looking for")
// 	ErrWordExists = errors.New("cannot add word because it already exists")
// )

// notice we don't need a pointer to dictionary here,
// because it is a value itself and not a struct?
func (d Dictionary) Search(word string) (string, error) {
	// a map lookup in Go can return two values -
	// the second value is a boolean which indicates if the key was found successfully.
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	// check if key exists, as adding an existing key to a map in go will overwrite the value.
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		// extra safety: return err if it's not ErrNotFound or ErrWordExists
		return err
	}

	// if sucessfully added then return error nil
	return nil
}

// Refactoring
// make the errors constant; which requirs making our own `DictionaryErr“ type which implements the `error` interface.
// This makes the errors more reusable and immutable
const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exits")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

// creating an error wrapper
type DictionaryErr string

// method to return the string of the error -
// any type with an `Error() string` method fulfills the `error` interface
func (e DictionaryErr) Error() string {
	// not sure why I can't just return e.
	return string(e)
}

// adding an update function to update the definition of a word
func (d Dictionary) Update(word, definition string) error {
	// we need to make sure the key exists before updating it,
	// otherwise it would simply add it to the map and behave like an `Add` function.

	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		// we could just retrun `ErrNotFound` but it's better to return a more specific error.
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

// adding a delete function to remove a word and it's definitions
func (d Dictionary) Delete(word string) {
	// maps have a built-in function `delete`.
	// `delete` doesn't return anything, and so our Delete function won't either.
	// deleting a value that doesn't exist has no effect so this is fine.
	delete(d, word)
}

// Maps are a bit confusing in the way they work / pointers...
// Maps can be modified without passing as an address to it (e.g. `&myMap`),
// this *feels* like a "reference type" but they are not.
// A map value is a pointer to a `runtime.hmap` structure -
// so when you pass a map to a function/method, you are indieed copying it,
// but just the pointer part, not the underlying data structure that contains the data.
// A gotcha with maps is that they can be a `nil` value.
// A `nil` map behaves like an empty map when reading, but attempts to write to a `nil` map
// will cause a runtime panic. Therefor never initialise a nil map variable (`var m map[string]string`)
// instead, initialise an empty map (`var m map[string]string{}`)
// or use the `make` keyword to create one: (`var m = make(map[string]string)`)
// bot create an empty `hash map` and point the variable to it, which means you will never get a runtime panic.
