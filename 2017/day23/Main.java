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
            System.err.println("Error reading input file.");
        }

        Register[] registers = createRegisters();
        List<Instruction> instructions = parseInstructions(input);
        System.out.println("Part 1: " + part1(instructions, registers));
        registers = createRegisters();
        System.out.println("Part 2: " + part2(instructions, registers));
    }

    private static int part1(List<Instruction> instructions, Register[] registers) {
        int count = 0;
        int idx = 0;
        while (idx >= 0 && idx < instructions.size()) {
            Instruction instruction = instructions.get(idx);
            switch (instruction.command) {
                case SET:
                    set(registers, instruction.x, instruction.y);
                    break;
                case SUB:
                    sub(registers, instruction.x, instruction.y);
                    break;
                case MUL:
                    mul(registers, instruction.x, instruction.y);
                    count++;
                    break;
                case JNZ:
                    idx += jnz(registers, instruction.x, instruction.y);
                    continue;
            }
            idx++;
        }

        return count;
    }

    // https://www.reddit.com/r/adventofcode/comments/7lms6p/2017_day_23_solutions/
    // https://www.reddit.com/r/adventofcode/comments/7lms6p/comment/drnhcx8/?utm_source=share&utm_medium=web2x&context=3
    // spesifically the comment above explains the logic. I just translated it to java.
    // once again, thanks reddit!
    private static int part2(List<Instruction> instructions, Register[] registers) {
        int b, c, d, f, g, h = 0;
        b = 67 * 100 + 100000;
        c = b + 17000;
        g = 0;
        h = 0;
        do {
            d = 2;
            f = 1;
            for (d = 2; d * d <= b; d++) {
                if (b % d == 0) {
                    f = 0;
                    break;
                }
            }
            if (f == 0) {
                h++;
            }
            g = b - c;
            b += 17;
        } while (g != 0);

        return h;
    }

    private static void set(Register[] registers, String x, String y) {
        int value = getValue(registers, y);
        registers[x.charAt(0) - 'a'].value = value;
    }

    private static void sub(Register[] registers, String x, String y) {
        int value = getValue(registers, y);
        registers[x.charAt(0) - 'a'].value -= value;
    }

    private static void mul(Register[] registers, String x, String y) {
        int value = getValue(registers, y);
        registers[x.charAt(0) - 'a'].value *= value;
    }

    private static int jnz(Register[] registers, String x, String y) {
        int value = getValue(registers, x);
        if (value != 0) {
            return getValue(registers, y);
        }
        return 1;
    }

    private static int getValue(Register[] registers, String x) {
        if (x.charAt(0) >= 'a' && x.charAt(0) <= 'z') {
            return registers[x.charAt(0) - 'a'].value;
        }
        return Integer.parseInt(x);
    }

    private static Register[] createRegisters() {
        Register[] registers = new Register[8];
        for (int i = 0; i < registers.length; i++) {
            registers[i] = new Register((char) ('a' + i));
        }
        return registers;
    }

    private static List<Instruction> parseInstructions(String input) {
        return input.lines().map(line -> {
            String[] parts = line.split(" ");
            Command command = Command.valueOf(parts[0].toUpperCase());
            String x = parts[1];
            String y = parts[2];
            return new Instruction(command, x, y);
        }).collect(Collectors.toList());
    }

    private static class Register {
        char name;
        int value;

        public Register(char name) {
            this.name = name;
            this.value = 0;
        }

        @Override
        public String toString() {
            return "Register{" + "name=" + name + ", value=" + value + '}';
        }
    }

    private static class Instruction {
        Command command;
        String x;
        String y;

        public Instruction(Command command, String x, String y) {
            this.command = command;
            this.x = x;
            this.y = y;
        }
    }

    private static enum Command {
        SET, SUB, MUL, JNZ
    }
}
