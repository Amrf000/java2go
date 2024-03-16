package test;

public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello World");
    }
}

private class inner {
    public inner(){}
    public int test(){
            System.out.println("test1");
            return 1;
    }

    public void test(int a){
            System.out.println("test2");
    }
    public String test(int a,String s){
            System.out.println("test3");
            return "returntest3";
    }
    public String test(String s,int a){
            System.out.println("test4");
            return "returntest4";
    }
    public static void main(String[] args){
        inner o = new inner();
        System.out.println(o.test());
        o.test(1);
        System.out.println(o.test(1,"test3"));
        System.out.println(o.test("test4",1));
    }
}