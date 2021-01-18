package gojsondb

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

// JsonDB is the main database driver
type JsonDB struct {
	Filename string `json:"filename"`
	Data map[string][]byte `json:"data"`
}

// New creates a new JsonDB at the specified filename, loading the contents if the file exists
func New(filename string) (*JsonDB, error) {
	result := JsonDB{Filename: filename, Data: make(map[string][]byte)}

	if fileExists(filename) {
		dat, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(dat, &result)
		if err != nil {
			return nil, err
		}
	}

	return &result, nil
}

// GetRaw returns the byte array data associated with the key
func (db *JsonDB) GetRaw(key string) ([]byte, error) {
	if val, ok := db.Data[key]; ok {
		return val, nil
	}

	return nil, &DataNotExistsError{}
}
// Get unmarshals the data associated with the key into the interface pointer
func (db *JsonDB) Get(key string, item *interface{}) error {
	result, err := db.GetRaw(key)
	if err != nil {
		return err
	}

	return json.Unmarshal(result, item)
}
// GetString returns the string associated with the key
func (db *JsonDB) GetString(key string) (string, error) {
	result, err := db.GetRaw(key)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

// SetRaw sets the data associated with a key
func (db *JsonDB) SetRaw(key string, data []byte) {
	db.Data[key] = data
}
// Set unmarshals an interface pointer into the data associated with a key
func (db *JsonDB) Set(key string, item *interface{}) error {
	result, err := json.Marshal(item)
	if err != nil {
		return err
	}

	db.SetRaw(key, result)
	return nil
}
// SetString sets a string associated with a key
func (db *JsonDB) SetString(key string, item string) {
	db.SetRaw(key, []byte(item))
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

// DataNotExistsError instructs the user that the key they've requested has no associated value
type DataNotExistsError struct{}
func (m *DataNotExistsError) Error() string {
    return "Dictionary has no value paired with this key"
}