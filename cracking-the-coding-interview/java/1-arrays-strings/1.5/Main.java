import java.lang.StringBuilder;

public class Main {
    public static void main(String[] args) {
        String test = compress("aabcccccaaa");
        System.out.println("result : " + test);
    }

    public static String compress(String str) {
        char[] chars = str.toCharArray();
        StringBuilder compressed = new StringBuilder();

        int count = 0;
        char last = chars[0];

        for (int i = 0; i < chars.length; i++) {
            if (i == 0) {
                last = chars[i];
                count++;
            } else if (last == chars[i]) {
                count++;
            } else {
                compressed.append(last);
                compressed.append(count);
                last = chars[i];
                count = 1;
            }
        }
        compressed.append(last);
        compressed.append(count);
        if (compressed.toString().length() < str.length() ) {
            return compressed.toString();
        } else {
            return str;
        }
    }
}
