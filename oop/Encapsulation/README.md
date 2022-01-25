# Encapsulation
Binding of code and data together in a capsule to hide the complexity of a class. Encapsulation deals with hiding the complexity of the program. 

For Example, following code still follows Encapusulation irrespecive of whether it hides data or not.

```
type Counter struct {
    Value integer
}
func (c *Counter) add() {
    c.Value = c.Value + 1
}
func (c *Counter) get() int {
    return c.Value
}
```

Data hiding deals with security of the data but not a responsiblity of Encapsulation.

Althogh the code doesn't look like it's actually wrapping data and methods (implentation) together within a block of code, the methods are applied on the pointer of the struct which makes it a capsule.

Usually, in the industry, we hide data to make it secure.
```
type Counter struct {
	value int
}

func (c *Counter) add() {
	c.value = c.value + 1
}

func (c *Counter) get() int {
	return c.value
}
```
Note : In go, lowercase variable names are not exported and Lettercase variable names are always exported. 