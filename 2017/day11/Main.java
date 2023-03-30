import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Map;

public class Main {

    private static final Map<String, Point> DIRECTIONS =
            Map.ofEntries(Map.entry("n", new Point(0, -1)), Map.entry("ne", new Point(1, -1)),
                    Map.entry("se", new Point(1, 0)), Map.entry("s", new Point(0, 1)),
                    Map.entry("sw", new Point(-1, 1)), Map.entry("nw", new Point(-1, 0)));

    public static void main(String[] args) {
        String input;
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException e) {
            System.err.println("Error reading input file. Stack trace: " + e.getStackTrace());
            return;
        }

        System.out.println("Part 1: " + part1(input));
        System.out.println("Part 2: " + part2(input));
    }

    private static int part1(String input) {
        Point current = new Point(0, 0);
        for (String direction : input.split(",")) {
            Point step = DIRECTIONS.get(direction);
            current.x += step.x;
            current.y += step.y;
        }
        return distance(current);
    }

    private static int part2(String input) {
        Point current = new Point(0, 0);
        int maxDistance = 0;
        for (String direction : input.split(",")) {
            Point step = DIRECTIONS.get(direction);
            current.x += step.x;
            current.y += step.y;
            maxDistance = Math.max(maxDistance, distance(current));
        }
        return maxDistance;
    }

    private static int distance(Point point) {
        return (Math.abs(point.x) + Math.abs(point.y) + Math.abs(point.x + point.y)) / 2;
    }

    private static class Point {
        int x;
        int y;

        public Point(int x, int y) {
            this.x = x;
            this.y = y;
        }
    }
}
