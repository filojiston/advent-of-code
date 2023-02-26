import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Main {
    public static void main(String[] args) {
        String input = "";
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException exception) {
            System.err.println("Error reading input file");
        }

        System.out.println("Part 1: " + part1(input));
        System.out.println("Part 2: " + part2(input));
    }

    // Find the sum of all digits that match the next digit in the list. The list is
    // circular, so the digit after the last digit is the first digit in the list.
    private static int part1(String input) {
        int sum = 0;
        for (int i = 0; i < input.length(); i++) {
            if (input.charAt(i) == input.charAt((i + 1) % input.length())) {
                sum += Character.getNumericValue(input.charAt(i));
            }
        }
        return sum;
    }

    // Now, instead of considering the next digit, it wants you to consider
    // the digit halfway around the circular list.
    // That is, if your list contains 10 items, only include a digit in your sum
    // if the digit 10/2 = 5 steps forward matches it.
    private static int part2(String input) {
        int sum = 0;
        for (int i = 0; i < input.length(); i++) {
            if (input.charAt(i) == input.charAt((i + input.length() / 2) % input.length())) {
                sum += Character.getNumericValue(input.charAt(i));
            }
        }
        return sum;
    }
}
