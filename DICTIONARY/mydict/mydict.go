package mydict

import (
	"errors"
)

// type define ex) type Money int
// Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errCantUpdate = errors.New("Cant update non-existing word")
	errWordExists = errors.New("That word already exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	// i, ok := m["route"]
	//map에 값이 있는지 boolean 타입인 ok로 확인 가능

	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add a word to dictionary
// map은 리시버에 * 필요 없음, 이미 갖고 있음
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil
}

// Update a word
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete a word
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
