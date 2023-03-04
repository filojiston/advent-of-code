import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;

public class Main {
    public static void main(String[] args) {
        String input;
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException e) {
            System.err.println("Could not read input file");
            return;
        }

        String[] passphrases = input.split("\n");
        System.out.println("Part 1: " + part1(passphrases));
        System.out.println("Part 2: " + part2(passphrases));
    }

    // To ensure security, a valid passphrase must contain no duplicate words.
    // How many passphrases are valid?
    private static int part1(String[] passphrases) {
        return (int) Arrays.stream(passphrases).map(passphrase -> passphrase.split(" "))
                .filter(words -> Arrays.stream(words).distinct().count() == words.length).count();
    }

    // Now, a valid passphrase must contain no two words that are anagrams of each other - that is,
    // a passphrase is invalid if any word's letters can be rearranged to form any other word in the
    // passphrase.
    // How many passphrases are valid?
    private static int part2(String[] passphrases) {
        return (int) Arrays.stream(passphrases).map(passphrase -> passphrase.split(" "))
                .filter(words -> !containsAnagram(words)).count();
    }

    private static boolean containsAnagram(String[] words) {
        for (int i = 0; i < words.length; i++) {
            for (int j = i + 1; j < words.length; j++) {
                if (isAnagram(words[i], words[j])) {
                    return true;
                }
            }
        }
        return false;
    }

    private static boolean isAnagram(String a, String b) {
        char[] aChars = a.toCharArray();
        char[] bChars = b.toCharArray();

        Arrays.sort(aChars);
        Arrays.sort(bChars);

        return Arrays.equals(aChars, bChars);
    }
}
