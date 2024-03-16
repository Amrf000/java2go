package test5;


class Animal {
    private String name;
    private int id;
    public Animal(){}
    public Animal(String myName, int myid) {
        //初始化属性值
    }
    public void eat() {  //吃东西方法的具体实现
    }
    public void sleep() { //睡觉方法的具体实现
    }
    public void move(){
          System.out.println("动物可以移动");
    }
}
class Dog extends Animal{
   public void move(){
        super.move();
        System.out.println("狗可以跑和走");
   }
}
public class Penguin  extends  Animal{
    public Penguin(String myName, int myid) {
        super(myName, myid);
    }
    public static void main(String args[]){
       Animal a = new Animal(); // Animal 对象
       Animal b = new Dog(); // Dog 对象

       a.move();// 执行 Animal 类的方法

       b.move();//执行 Dog 类的方法
    }
}

class Mouse extends Animal {
    public Mouse(String myName, int myid) {
        super(myName, myid);
    }
}
