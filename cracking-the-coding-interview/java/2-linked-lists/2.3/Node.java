import java.lang.StringBuilder;

public class Node {
    Node next = null;
    int data;

    public Node(int d) {
        data = d;
    }

    void appendToTail(int d) {
        Node end = new Node(d);
        Node n = this;
        while (n.next != null) {
            n = n.next;
        }

        n.next = end;
    }

    Node deleteNode(int d) {
        Node head = this;
        Node n = this;

        if (n.data == d) {
            return head.next;
        }

        while (n.next != null) {
            if (n.next.data == d) {
                n.next = n.next.next;
                return head;
            }
            n = n.next;
        }
        return head;
    }

    @Override public String toString() {
        StringBuilder sb = new StringBuilder();
        Node n = this;
        sb.append(n.data);
        while (n.next != null) {
            sb.append(n.next.data);
            n = n.next;
        }
        return sb.toString();
    }
}
