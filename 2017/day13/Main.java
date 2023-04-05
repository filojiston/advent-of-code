import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;
import java.util.stream.IntStream;

// my initial solution was working too slow on part2 because i was simulating the firewall
// so i had to look for a better solution, and i found this one:
// https://www.reddit.com/r/adventofcode/comments/7jgyrt/comment/dr6cuza/?utm_source=share&utm_medium=web2x&context=3
// thx dude
public class Main {
    public static void main(String[] args) {
        String input = "";
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException exception) {
            System.err.println("Error reading input file. Stack trace: " + exception);
        }

        Map<Integer, Integer> layers = parseLayers(input);
        System.out.println("Part 1: " + part1(layers));
        System.out.println("Part 2: " + part2(layers));
    }

    private static int part1(Map<Integer, Integer> layers) {
        return layers.entrySet().stream()
                .filter(entry -> entry.getKey() % (2 * (entry.getValue() - 1)) == 0)
                .mapToInt(entry -> entry.getKey() * entry.getValue()).sum();
    }

    private static int part2(Map<Integer, Integer> layers) {
        return IntStream.iterate(0, i -> i + 1)
                .filter(delay -> layers.entrySet().stream().noneMatch(
                        entry -> (entry.getKey() + delay) % (2 * (entry.getValue() - 1)) == 0))
                .findFirst().getAsInt();
    }

    private static Map<Integer, Integer> parseLayers(String input) {
        Map<Integer, Integer> layers = new HashMap<>();
        Arrays.stream(input.split("\n")).map(line -> line.split(": ")).forEach(
                parts -> layers.put(Integer.parseInt(parts[0]), Integer.parseInt(parts[1])));
        return layers;
    }
}
