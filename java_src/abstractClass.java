package java_src;

// Java allows abstract classes to be created
// Abstract classes cannot be instantiated and require subclasses to provide implementations for the abstract methods
abstract class car {
    public void vehicleType() {
        System.out.println("Vehicle Type: Car");
    }
    public abstract void brand();
}

class porsche extends car {
    @Override
    public void brand() {
        System.out.println("Brand: Porsche");
    }
    public void speed() {
        System.out.println("Max: 200kmph");
    }
}

class audi extends car {
    @Override
    public void brand() {
        System.out.println("Brand: Audi");
    }
    public void speed() {
        System.out.println("Max: 180kmph");
    }
}

public class abstractClass {
    public static void main(String[] args) {
        porsche p1 = new porsche();
        p1.vehicleType();
        p1.brand();
        p1.speed();

        audi a1 = new audi();
        a1.vehicleType();
        a1.brand();
        a1.speed();
    }
    
}
