package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrorInvalid        = DictionaryErr("invalid")
	ErrorKeyExists      = DictionaryErr("key already exists")
	ErrorKeyDoesntExist = DictionaryErr("key does not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	if value, ok := d[key]; ok {
		return value, nil
	}
	return "", ErrorInvalid

}

func (d Dictionary) Add(key, value string) error {
	if _, ok := d[key]; ok {
		return ErrorKeyExists
	}
	d[key] = value
	return nil
}

func (d Dictionary) Update(key, value string) error {
	if _, ok := d[key]; ok {
		d[key] = value
		return nil
	}
	return ErrorKeyDoesntExist
}

func (d Dictionary) Delete(key string) error {
	if _, ok := d[key]; ok {
		delete(d, key)
		return nil
	}
	return ErrorKeyDoesntExist
}
