## Usage

#### type DataNotExistsError

```go
type DataNotExistsError struct{}
```

DataNotExistsError instructs the user that the key they've requested has no
associated value

#### func (*DataNotExistsError) Error

```go
func (m *DataNotExistsError) Error() string
```

#### type JsonDB

```go
type JsonDB struct {
	Filename string            `json:"filename"`
	Data     map[string][]byte `json:"data"`
}
```

JsonDB is the main database driver

#### func  New

```go
func New(filename string) (*JsonDB, error)
```
New creates a new JsonDB at the specified filename, loading the contents if the
file exists

#### func (*JsonDB) Get

```go
func (db *JsonDB) Get(key string, item interface{}) error
```
Get unmarshals the data associated with the key into the interface pointer

#### func (*JsonDB) GetRaw

```go
func (db *JsonDB) GetRaw(key string) ([]byte, error)
```
GetRaw returns the byte array data associated with the key

#### func (*JsonDB) GetString

```go
func (db *JsonDB) GetString(key string) (string, error)
```
GetString returns the string associated with the key

#### func (*JsonDB) Save

```go
func (db *JsonDB) Save() error
```
Save save's the contents of the database to its corresponding filepath

#### func (*JsonDB) Set

```go
func (db *JsonDB) Set(key string, item interface{}) error
```
Set unmarshals an interface pointer into the data associated with a key

#### func (*JsonDB) SetRaw

```go
func (db *JsonDB) SetRaw(key string, data []byte) error
```
SetRaw sets the data associated with a key

#### func (*JsonDB) SetString

```go
func (db *JsonDB) SetString(key string, item string) error
```
SetString sets a string associated with a key
