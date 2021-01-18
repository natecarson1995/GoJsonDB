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
func (db *JsonDB) Get(key string, item interface{}) error {
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

// Delete deletes a key and its associated data
func (db *JsonDB) Delete(key string) error {
	delete(db.Data, key)
	return db.Save()
}

// ListKeys lists all of the keys in the map of data
func (db *JsonDB) ListKeys() []string {
	var result []string
	for key := range db.Data {
		result = append(result, key)
	}
	return result
}

// SetRaw sets the data associated with a key
func (db *JsonDB) SetRaw(key string, data []byte) error {
	db.Data[key] = data
	return db.Save()
}
// Set unmarshals an interface pointer into the data associated with a key
func (db *JsonDB) Set(key string, item interface{}) error {
	result, err := json.Marshal(item)
	if err != nil {
		return err
	}

	return db.SetRaw(key, result)
}
// SetString sets a string associated with a key
func (db *JsonDB) SetString(key string, item string) error {
	return db.SetRaw(key, []byte(item))
}

// Save saves the contents of the database to its corresponding filepath
func (db *JsonDB) Save() error {
	data, err := json.Marshal(db)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(db.Filename, data, os.ModePerm)
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