import java.io.IOException;
import java.nio.CharBuffer;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;

public class Main {
    private static final String START = ".#./..#/###";
    private static final int PART1_ITERATIONS = 5;
    private static final int PART2_ITERATIONS = 18;

    public static void main(String[] args) {
        String input = "";
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException exception) {
            System.err.println("Error reading input file");
        }

        Map<String, String> enhancements = parseInput(input);
        System.out.println("Part 1: " + solve(enhancements, PART1_ITERATIONS));
        System.out.println("Part 2: " + solve(enhancements, PART2_ITERATIONS));
    }

    private static int solve(Map<String, String> enhancements, int iterations) {
        char[][] grid = getGrid(START);
        for (int i = 0; i < iterations; i++) {
            grid = enhance(grid, enhancements);
        }
        return countOn(grid);
    }

    private static char[][] enhance(char[][] grid, Map<String, String> enhancements) {
        int size = grid.length;
        int subSize = size % 2 == 0 ? 2 : 3;
        int newSize = size / subSize * (subSize + 1);
        char[][] newGrid = new char[newSize][newSize];
        for (int i = 0; i < size; i += subSize) {
            for (int j = 0; j < size; j += subSize) {
                char[][] subGrid = new char[subSize][subSize];
                for (int k = 0; k < subSize; k++) {
                    for (int l = 0; l < subSize; l++) {
                        subGrid[k][l] = grid[i + k][j + l];
                    }
                }
                String key = gridToString(subGrid);
                char[][] enhancedSubGrid = getGrid(enhancements.get(key));
                for (int k = 0; k < subSize + 1; k++) {
                    for (int l = 0; l < subSize + 1; l++) {
                        newGrid[i * (subSize + 1) / subSize + k][j * (subSize + 1) / subSize + l] =
                                enhancedSubGrid[k][l];
                    }
                }
            }
        }
        return newGrid;
    }

    private static int countOn(char[][] grid) {
        return Arrays.stream(grid).mapToInt(row -> {
            return (int) CharBuffer.wrap(row).chars().filter(c -> c == '#').count();
        }).sum();
    }

    private static Map<String, String> parseInput(String input) {
        Map<String, String> enhancements = new HashMap<>();
        input.lines().forEach(line -> {
            String[] parts = line.split(" => ");
            String key = parts[0];
            List<String> keyRotations = getRotations(key);
            keyRotations.forEach(rotation -> {
                enhancements.put(rotation, parts[1]);
            });
        });
        return enhancements;
    }

    private static List<String> getRotations(String key) {
        List<String> rotations = new ArrayList<>();
        char[][] grid = getGrid(key);
        for (int i = 0; i < 4; i++) {
            grid = transpose(grid);
            grid = flip(grid);
            rotations.add(gridToString(grid));
        }
        grid = flip(grid);
        for (int i = 0; i < 4; i++) {
            grid = transpose(grid);
            grid = flip(grid);
            rotations.add(gridToString(grid));
        }
        return rotations;
    }

    private static String gridToString(char[][] grid) {
        return Arrays.stream(grid).map(String::valueOf).collect(Collectors.joining("/"));
    }

    private static char[][] getGrid(String key) {
        String[] rows = key.split("/");
        int size = rows.length;
        char[][] grid = new char[size][size];
        for (int i = 0; i < size; i++) {
            grid[i] = rows[i].toCharArray();
        }
        return grid;
    }

    private static char[][] transpose(char[][] input) {
        int size = input.length;
        char[][] output = new char[size][size];
        for (int i = 0; i < size; i++) {
            char[] row = input[i];
            for (int j = 0; j < size; j++) {
                output[j][i] = row[j];
            }
        }
        return output;
    }

    private static char[][] flip(char[][] input) {
        int size = input.length;
        char[][] output = new char[size][size];
        for (int i = 0; i < size; i++) {
            char[] row = input[i];
            for (int j = 0; j < size; j++) {
                output[i][j] = row[size - j - 1];
            }
        }
        return output;
    }
}
