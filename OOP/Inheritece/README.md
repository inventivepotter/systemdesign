# Inheritence
Inheriting the properties of the superclass into the base class and is one of the most important concepts in Object-Oriented Programming. 

In Golang inheritence can achieved by using composition where the struct is used to form other objects. 

Given there are no classes, you can say there is No Inheritance Concept in Golang.

In composition, base structs can be embedded into a child struct and the methods of the base struct can be directly called on the child struct as shown in the following example.

Note:
- Interface inheritance: Way to create implementations that adhere to abstractions.
- Class inheritance violates encapsulation - The below code is an example to justify the statement (Java).
  ```
    class SpecialSet extends HashSet {
        int count = 0
        add(int x){
            this.count++;
            return super.add(x);
        }
        addAll(int args[]){
            this.count++;
            return super.addAll();
        }    
    }
  ```
  but HashSet implemented `addAll()` method internally by call `add()` method for each input, we wouldn't know it until we actually execute above code.
    ```
    Set set = new SpecialSet()
    set.add(1) ⇒ size=1, addCount = 1
    set.addAll([1, 2, 3]) ⇒ size=3, addCount = 1 + 3 = 4
    ```
- Class Inheritance pollutes the public interface because some times unncessary methods apply to .
- Compositions are better to achieve inheritence than class inheritence.