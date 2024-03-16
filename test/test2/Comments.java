package test2;


/**
 * 这是一个文档注释示例
 * 它通常包含有关类、方法或字段的详细信息
 */
public class Comments {
	// 类的成员和方法
	public int x = 10; // 初始化一个变量x为10

	public Comments() {
		// 这是一个单行注释
		int y = 10; // 初始化一个变量y为20

/*
这是一个多行注释
可以用来注释多行代码
*/
		int z = 10; // 初始化一个变量y为20
	}
    public Comments(String name){
        // 这个构造器仅有一个参数：name
        x = 2;
    }
    public Comments(Comments that){
            // 这个构造器仅有一个参数：name
            x = 2;
    }
    public String Test(int a[]) {}
	public String Test() {
// 这是一个单行注释
		int y = 10; // 初始化一个变量y为20

/*
   这是一个多行注释
   可以用来注释多行代码
*/
		int z = 10; // 初始化一个变量y为20
		x = 3;
		this.x++;
		int w = ++this.x;
		int z = (w == 2) ? 1 : 5;
		return "";
	}
//     public static Object ternary(boolean condition ,Object a,Object b)  {
// 	    if (condition) {
// 		    return a;
// 	    }
// 	    return b;
//     }
	public static void main(String[] args){
	    Comments o = new Comments();
	    Comments b = new Comments("test");
	    Comments c = new Comments(o);
	    o.Test();
	    b.Test();
	    int aa[] = {0};
	    o.Test(aa);
	}
}


