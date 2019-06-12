package main

type Dictionary map[string]string

type ErrorString string

func (e ErrorString) Error() string {
	return string(e)
}

const (
	ErrKeyNotExists = ErrorString("key do not exists")
	ErrKeyExists = ErrorString("key already exists")
)

func (d Dictionary) Search(key string) (string, error) {
	v, ok := d[key]
	if !ok {
		return "", ErrKeyNotExists
	}
	return v, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrKeyNotExists:
		d[key] = value
	case nil:
		return ErrKeyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrKeyNotExists:
		return ErrKeyNotExists
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}

func main() {

}
