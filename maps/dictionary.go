package main

type Dictionary map[string]string

const (
	ErrWordNotFound     = DictErr("Word not found")
	ErrWordExists       = DictErr("Word already defined")
	ErrWordDoesNotExist = DictErr("Word not defined")
)

type DictErr string

func (e DictErr) Error() string {
	return string(e)
}

// Search a definition in the dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, found := d[word]

	if !found {
		return "", ErrWordNotFound
	}

	return definition, nil
}

// Add a definition to the dictionary
func (d Dictionary) Add(word string, def string) error {
	_, exists := d[word]
	if exists {
		return ErrWordExists
	}

	d[word] = def
	return nil
}

// Update a definition in the dictionary
func (d Dictionary) Update(word, def string) error {
	_, exists := d[word]

	if !exists {
		return ErrWordDoesNotExist
	}

	d[word] = def
	return nil
}

// Delete a definition on the dictionary
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
