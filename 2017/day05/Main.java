import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Main {
    public static void main(String[] args) {
        String input;
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException exception) {
            System.err.println("Error reading file: " + exception.getMessage());
            return;
        }

        int[] instructions = input.lines().mapToInt(Integer::parseInt).toArray();
        System.out.println("Part 1: " + applyInstructions(instructions.clone(), 1));
        System.out.println("Part 2: " + applyInstructions(instructions.clone(), 2));
    }

    private static int applyInstructions(int[] instructions, int part) {
        int steps = 0;
        int index = 0;
        while (index >= 0 && index < instructions.length) {
            int offset = instructions[index];
            instructions[index] = part == 1 ? offset + 1 : offset >= 3 ? offset - 1 : offset + 1;
            index += offset;
            steps++;
        }

        return steps;
    }
}
