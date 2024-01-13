package mapsdict

type Dictionary map[string]string

var (
	ErrNotFound       = DictionaryErr("could not find the word you were looking for")
	ErrWordExists     = DictionaryErr("dictionary key already exists")
	ErrWordDontExists = DictionaryErr("dictionary key doe not exists")
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

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, newValue string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		return ErrWordDontExists
	case nil:
		d[key] = newValue
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
