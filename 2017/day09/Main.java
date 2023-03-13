import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Main {
    private static final Pattern IGNORE_PATTERN = Pattern.compile("!.");
    private static final Pattern GARBAGE_PATTERN = Pattern.compile("<[^>]*>");

    public static void main(String[] args) {
        String input;
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException exception) {
            System.err.println("Error reading input file");
            return;
        }

        System.out.println("Part 1: " + part1(input));
        System.out.println("Part 2: " + part2(input));
    }

    private static int part1(String input) {
        input = IGNORE_PATTERN.matcher(input).replaceAll("");
        input = GARBAGE_PATTERN.matcher(input).replaceAll("");

        int sum = 0;
        int depth = 1;

        for (char c : input.toCharArray()) {
            if (c == '{') {
                sum += depth;
                depth++;
            } else if (c == '}') {
                depth--;
            }
        }

        return sum;
    }

    private static int part2(String input) {
        input = IGNORE_PATTERN.matcher(input).replaceAll("");
        Matcher matcher = GARBAGE_PATTERN.matcher(input);

        int sum = 0;
        while (matcher.find()) {
            sum += matcher.end() - matcher.start() - 2;
        }

        return sum;
    }
}
