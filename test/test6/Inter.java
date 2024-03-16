package Inter;

interface A {
    public void eat();
    public void sleep();
}

interface B {
    public void show();
}

public class Inter implements A,B {
    public void eat() {}
    public void sleep() {}
    public void show() {}
}
