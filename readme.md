# n36
--
    import "github.com/gianebao/groxy_app/n36"


## Usage

```go
const (
	CharRange62 = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	CharRange36 = "0123456789abcdefghijklmnopqrstuvwxyz"
)
```

#### type N36

```go
type N36 struct {
}
```

N36 represents a numeric map

#### func  New

```go
func New(charset string) *N36
```
New creates a new n36 numeric map

#### func (*N36) Iton

```go
func (n *N36) Iton(i uint64) string
```
Iton converts a uint64 value to string

#### func (*N36) Ntoi

```go
func (n *N36) Ntoi(s string) (uint64, error)
```
Ntoi converts a string to uint64

#### func (*N36) Random

```go
func (n *N36) Random(l int) string
```
Random creates an (l)-long random string based on the character set
