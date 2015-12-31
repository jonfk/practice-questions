import java.util.Arrays;
import java.util.Random;


public class Main {

    static Random rnd = new Random(234);

    public static void main(String[] args) {
        System.out.println("result " + perm("abcdz", "dcba"));
    }

    public static boolean perm(String one, String two) {
        String oneSorted = sort(one);
        String twoSorted = sort(two);
        return oneSorted.equals(twoSorted);
    }

    public static String sort(String str) {
        char[] chars = str.toCharArray();
        sort(chars, 0, str.length() - 1);
        return new String(chars);
    }

    public static void sort(char[] arr, int left, int right) {
        if(left < right){
            int pivot = Partition(arr, left, right);

            if(pivot > 1)
                sort(arr, left, pivot - 1);

            if(pivot + 1 < right)
                sort(arr, pivot + 1, right);
        }
    }

    public static int Partition(char[] numbers, int left, int right) {
        char pivot = numbers[left];
        while (true){
            while (numbers[left] < pivot)
                left++;

            while (numbers[right] > pivot)
                right--;

            if (left < right){
                char temp = numbers[right];
                numbers[right] = numbers[left];
                numbers[left] = temp;
            }
            else{
                return right;
            }
        }
    }

}
