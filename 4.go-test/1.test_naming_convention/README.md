## Basic Golang Test

Run  `go test` to test golang tests.
Run `go test -v` to turn on the verbose mode.

### 1.Test File Naming Conventions
The test file name is within the same folder as the go code, and usually with underscore plus test. Go build will automatically ignore those test files. Ex:
`main.go, main_test.go`

### 2. Test Function Naming Conventions
The function is name as below. Put the type or Function name after the word Test

```go
func TestDog (t *testing.T) {

}

func TestDog_Bark (t *testing.T) {

}

func TestBark (t *testing.T) {

}
```

### 3.Test Variables Naming Conventions
Use `got`, `want`, `arg` for test variables' names

```
func TestColor (t *testing.T) {
  arg := "blue"
  want := "#0000FE"
  got := Color(arg)
  if got != want {
    t.Errorf("color(%q)= %q, want %q", arg, got, want)
  }
}
```
