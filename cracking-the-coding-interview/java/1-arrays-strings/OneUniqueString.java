import java.util.HashSet;
import java.lang.Character;


public class OneUniqueString {
    public static void main(String []args) {
        System.out.println(uniqueCharsString("aoeua"));
    }

    public static boolean uniqueCharsString(String str) {
        HashSet<Character> chars = new HashSet<Character>();

        for (char a : str.toCharArray()) {
            if (chars.contains(Character.valueOf(a))) {
                return false;
            }
            chars.add(Character.valueOf(a));
        }
        return true;
    }
}
