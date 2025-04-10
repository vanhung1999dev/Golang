## 🎭 Fat Pointer in Interfaces

An interface value is also a fat pointer. <br>

It’s implemented like this: <br>

```
type InterfaceHeader struct {
    Type  *rtype  // Pointer to type information (for method dispatch)
    Value uintptr // Pointer to the actual value
}

```

So an interface is a fat pointer: it carries both the type info and the data pointer. <br>

This is what allows Go’s interfaces to do dynamic dispatch and store values of different types. <br>
