package dictionary

type Dictionary map[string]string

const (
	ErrNotFound = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot perform operation on word because it does not exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	if err == ErrNotFound {
		d[word] = definition
	} else if err == nil {
		return ErrWordExists
	} else {
		return err
	}
	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	if err == ErrNotFound {
		return ErrWordDoesNotExist
	} else if err == nil {
		d[word] = definition
	} else {
		return err
	}

	return nil
}

func (d Dictionary) Delete(word string) error {
_, err := d.Search(word)

	if err == ErrNotFound {
		return ErrWordDoesNotExist
	} else if err == nil {
		delete(d, word)
	} else {
		return err
	}

	return nil
}

