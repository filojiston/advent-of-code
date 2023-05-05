import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;
import java.util.stream.Collectors;

public class Main {
    public static void main(String[] args) {
        String input = "";
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException exception) {
            System.err.println("Error reading input file");
        }

        Point start = findStart(input);
        Grid grid = parseGrid(input, start);
        solve(grid);
    }

    private static void solve(Grid grid) {
        StringBuilder path = new StringBuilder();
        Point current = grid.start;
        char direction = 'd';
        int steps = 0;
        while (true) {
            char currentChar = grid.get(current);
            if (currentChar == ' ') {
                break;
            } else if (currentChar == '+') {
                direction = findDirection(grid, current, direction);
            } else if (currentChar != '|' && currentChar != '-') {
                path.append(currentChar);
            }
            current = move(current, direction);
            steps++;
        }
        System.out.println("Part 1: " + path.toString());
        System.out.println("Part 2: " + steps);
    }

    private static char findDirection(Grid grid, Point current, char direction) {
        if (direction == 'd' || direction == 'u') {
            if (grid.get(move(current, 'l')) != ' ') {
                return 'l';
            } else {
                return 'r';
            }
        } else {
            if (grid.get(move(current, 'u')) != ' ') {
                return 'u';
            } else {
                return 'd';
            }
        }
    }

    private static Point move(Point p, char direction) {
        switch (direction) {
            case 'u':
                return new Point(p.x, p.y - 1);
            case 'd':
                return new Point(p.x, p.y + 1);
            case 'l':
                return new Point(p.x - 1, p.y);
            case 'r':
                return new Point(p.x + 1, p.y);
            default:
                throw new IllegalArgumentException("Invalid direction");
        }
    }

    private static Grid parseGrid(String input, Point start) {
        List<List<Character>> grid = input.lines().map(line -> {
            return line.chars().mapToObj(c -> (char) c).collect(Collectors.toList());
        }).collect(Collectors.toList());

        return new Grid(grid, start);
    }

    private static Point findStart(String input) {
        String firstLine = input.split("\n")[0];
        return new Point(firstLine.indexOf('|'), 0);
    }

    private static class Grid {
        List<List<Character>> grid;
        Point start;

        Grid(List<List<Character>> grid, Point start) {
            this.grid = grid;
            this.start = start;
        }

        char get(Point p) {
            return grid.get(p.y).get(p.x);
        }
    }

    private static class Point {
        int x;
        int y;

        Point(int x, int y) {
            this.x = x;
            this.y = y;
        }
    }
}
