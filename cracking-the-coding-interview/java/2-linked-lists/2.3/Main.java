
public class Main {
    public static void main(String[] args) {
        Node test = new Node(9);
        test.appendToTail(8);
        test.appendToTail(7);
        test.appendToTail(6);
        test.appendToTail(5);
        test.appendToTail(4);
        test.appendToTail(3);
        System.out.println(test.toString());
    }

}
