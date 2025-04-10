## Structs are never nil unless you use a pointer to struct.

```
type MyStruct struct {
    Name string
}

var s MyStruct
fmt.Println(s == nil) // ❌ compile error

```

You can’t compare a struct to nil — it’s a value type. <br>

```
var ps *MyStruct = nil
fmt.Println(ps == nil) // ✅ true

```

So for structs: <br>

- You only get nil when you're dealing with pointers
- A value struct (even with zero fields) is never nil
