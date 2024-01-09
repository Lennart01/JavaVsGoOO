# Basic Inheritance

## Source Files
* [Java](./java_src/BasicInheritance.java)
* [Go](./go_src/basicinheritance.go)

## Java

Java fully supports inheritance, which combines subtyping and virtual methods, using the `extends` keyword. Inheritance in Java allows for defining subtyping relationships, known as nominal subtyping.
In our example, we create a `Car` class with two methods: `vehicleType` and `brand`. We then define `Porsche` and `Audi` classes that extend `Car`. Both classes override the `brand` method and introduce a new method, `speed`.

## Go

Go does not support classic inheritance as seen in languages like Java. Instead, it uses a combination of struct embedding and interfaces to achieve polymorphic behavior, which is commonly achieved through inheritance in other languages.

Go employs structural subtyping. In Go:

- There are no classes, but rather structs and interfaces.
- Methods can be defined "freely" in the sense that they are not bound to a class hierarchy.
- Methods can be overloaded, meaning that the receiver type and the method name must uniquely determine the instance.
- Interfaces describe receivers that satisfy the methods contained within the interface.

In our example:

- We define a `Vehicle` interface with two methods: `VehicleType` and `Brand`.
- A `Car` struct is created as a generic car, implementing the `Vehicle` interface by providing `VehicleType` and `Brand` methods. In Go, interface implementation is implicit: any struct that provides the required methods is considered to implement the interface.
- `Porsche` and `Audi` structs are then defined, each embedding `Car`. This embedding allows them to inherit `Car`'s methods and properties.
- Both `Porsche` and `Audi` introduce a `Speed` method and provide their own implementations of the `Brand` method, effectively overriding the `Car`'s `Brand` method. This showcases how Go uses composition and method redefinition to achieve a similar outcome to what is typically achieved with classical inheritance.

## Abstract Classes

### Source Files
* [Java](./java_src/AbstractClass.java)
* [Go](./go_src/abstractClass.go)
