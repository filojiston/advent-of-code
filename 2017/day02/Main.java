import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
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

        List<List<Integer>> spreadsheet = parseSpreadsheet(input);
        System.out.println("Part 1: " + part1(spreadsheet));
        System.out.println("Part 2: " + part2(spreadsheet));
    }

    // For each row, determine the difference between the largest value and
    // the smallest value; the checksum is the sum of all of these differences.
    private static int part1(List<List<Integer>> spreadsheet) {
        return spreadsheet.stream().map(row -> {
            int max = row.stream().max(Integer::compareTo).get();
            int min = row.stream().min(Integer::compareTo).get();
            return max - min;
        }).reduce(0, Integer::sum);
    }

    // For each row, find the only two numbers where one evenly divides the other
    // the checksum is the sum of all of these quotients.
    private static int part2(List<List<Integer>> spreadsheet) {
        return spreadsheet.stream().map(row -> {
            for (int i = 0; i < row.size(); i++) {
                for (int j = 0; j < row.size(); j++) {
                    if (i != j && row.get(i) % row.get(j) == 0) {
                        return row.get(i) / row.get(j);
                    }
                }
            }
            return 0;
        }).reduce(0, Integer::sum);
    }

    private static List<List<Integer>> parseSpreadsheet(String input) {
        return Arrays
                .stream(input.split("\n")).map(line -> Arrays.stream(line.split("\\s+"))
                        .map(Integer::parseInt).collect(Collectors.toList()))
                .collect(Collectors.toList());
    }
}
