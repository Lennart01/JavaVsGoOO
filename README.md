## Basic Inheritance

### Source Files
* [Java](./java_src/BasicInheritance.java)
* [Go](./go_src/basicinheritance.go)

### Java

Java fully supports inheritance using the `extends` keyword. In our example, we create a `Car` class with two methods: `vehicleType` and `brand`. We then define `Porsche` and `Audi` classes that extend `Car`. Both classes override the `brand` method and introduce a new method, `speed`.

### Go

Go does not support classic inheritance. Instead, it uses struct embedding and interfaces for similar outcomes. In our example:

- We define a `Vehicle` interface with two methods: `VehicleType` and `Brand`.
- A `Car` struct is created as a generic car, implementing the `Vehicle` interface by providing `VehicleType` and `Brand` methods. In Go, interface implementation is implicit: any struct that has the required methods is considered to implement the interface.
- `Porsche` and `Audi` structs are then defined, each embedding `Car`. This embedding allows them to inherit `Car`'s methods and properties.
- Both `Porsche` and `Audi` introduce a `Speed` method and provide their own implementations of the `Brand` method, effectively overriding the `Car`'s `Brand` method. This is an example of how Go utilizes composition and method redefinition in place of classic inheritance.

## Abstract Classes

### Source Files
* [Java](./java_src/AbstractClass.java)
* [Go](./go_src/abstractClass.go)

### Java
