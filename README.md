# Template literals in Go

```go
	templ := NewTempl()
	name := "bob"
	templ.Printf("this is {name}\n", name) // this is bob

	age := 6
	msg := templ.Sprintf("this is {name} and he is {age} years old\n", name, age) // this is bob and he is 6 years old
	templ.Print(msg)
```
