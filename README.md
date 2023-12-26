# JavaVsGoOO
## Basic Inheritance

### Source Files
* [Java](./java_src/BasicInheritance.java)
* [Go](./go_src/basic_inheritance.go)

### Java

Java offers full support for inheritance. The `extends` keyword is used to indicate that a class inherits from another class. In the linked example we create a car class with two methods `verhicleType` and `brand`.
Based on this class we created a `porsche`` and a audi` `class` wich `extends` the car class. Both override the `brand` method and both implement the `speed` method.

### Go

Go does not support `classic` inheritance. Instead we make use of Struct embedding and Interfaces. In the linked example we create a `Verhicle` interface with two methods `VerhicleType` and `Brand`. Based on this interface we create a `Car` struct which acts as a generic car. This struct implements the `Verhicle` interface by implementing the `VerhicleType` and `Brand` methods In Go every struct that implements all methods of an interface is automatically an implementation of that interface.
In Go, this is known as structural typing, as opposed to Java's nominal typing. The key difference is that in Go, there is no need to explicitly declare that a struct implements an interface. If the methods match, the implementation is implicit.
To make use of struct embedding we create a `Porsche` and a `Audi` struct which embeds the `Car` struct. This means that the `Porsche` and `Audi` structs inherit all methods of the `Car` struct.
This is not a classic inheritance but rather a composition of structs.
Both now implement the `Speed` method.
By implementing the Brand method for both structs we override the Brand method of the `Car` struct.
