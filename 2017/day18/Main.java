import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Queue;

public class Main {
    public static void main(String[] args) {
        String input = "";
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException e) {
            System.err.println("Error reading input file");
        }
        part1(input);
        part2(input);
    }

    private static void part1(String input) {
        Map<String, Long> registers = parseRegisters(input);
        List<Instruction> instructions = parseInstructions(input);
        Program program = new Program(registers, instructions, null);
        while (!program.waiting) {
            program.applyNextInstruction();
        }
        System.out.println("Part 1: " + program.lastSound);
    }

    private static void part2(String input) {
        Map<String, Long> registers0 = parseRegisters(input);
        List<Instruction> instructions = parseInstructions(input);
        Program program0 = new Program(registers0, instructions, null);

        Map<String, Long> registers1 = parseRegisters(input);
        registers1.put("p", 1L);
        Program program1 = new Program(registers1, instructions, program0);
        program0.other = program1;

        while (!program0.waiting || !program1.waiting) {
            program0.applyNextInstruction();
            program1.applyNextInstruction();
        }

        System.out.println("Part 2: " + program1.sent);
    }

    private static Map<String, Long> parseRegisters(String input) {
        Map<String, Long> registers = new HashMap<>();
        for (String line : input.split("\n")) {
            String[] parts = line.split(" ");
            if (!parts[1].matches("-?\\d+")) {
                registers.put(parts[1], 0L);
            }
            if (parts.length == 3) {
                if (!parts[2].matches("-?\\d+")) {
                    registers.put(parts[2], 0L);
                }
            }
        }
        return registers;
    }

    private static List<Instruction> parseInstructions(String input) {
        List<Instruction> instructions = new ArrayList<>();
        for (String line : input.split("\n")) {
            String[] parts = line.split(" ");
            List<String> args = new ArrayList<>();
            for (int i = 1; i < parts.length; i++) {
                args.add(parts[i]);
            }
            instructions.add(new Instruction(parts[0], args));
        }
        return instructions;
    }

    private static class Program {
        Map<String, Long> registers;
        List<Instruction> instructions;
        Queue<Long> queue;
        boolean waiting;
        int current;
        int sent;
        long lastSound;
        Program other;

        public Program(Map<String, Long> registers, List<Instruction> instructions, Program other) {
            this.registers = registers;
            this.instructions = instructions;
            this.queue = new LinkedList<>();
            this.waiting = false;
            this.current = 0;
            this.sent = 0;
            this.lastSound = 0L;
            this.other = other;
        }

        private void applyNextInstruction() {
            Instruction instruction = instructions.get(current);
            switch (instruction.name) {
                case "snd":
                    snd(registers, instruction.args);
                    current++;
                    break;
                case "set":
                    set(registers, instruction.args);
                    current++;
                    break;
                case "add":
                    add(registers, instruction.args);
                    current++;
                    break;
                case "mul":
                    mul(registers, instruction.args);
                    current++;
                    break;
                case "mod":
                    mod(registers, instruction.args);
                    current++;
                    break;
                case "rcv":
                    rcv(registers, instruction.args);
                    current = waiting ? current : current + 1;
                    break;
                case "jgz":
                    current += jgz(registers, instruction.args);
                    break;
                default:
                    System.err.println("Unknown instruction: " + instruction.name);
            }
        }

        void snd(Map<String, Long> registers, List<String> args) {
            long value = args.get(0).matches("-?\\d+") ? Long.parseLong(args.get(0))
                    : registers.get(args.get(0));
            if (other != null) {
                other.queue.add(value);
            } else {
                lastSound = value;
            }
            sent++;
        }

        void rcv(Map<String, Long> registers, List<String> args) {
            if (queue.isEmpty()) {
                waiting = true;
                return;
            }
            long value = queue.poll();
            waiting = false;
            registers.put(args.get(0), value);
        }

        void set(Map<String, Long> registers, List<String> args) {
            long newValue = args.get(1).matches("-?\\d+") ? Long.parseLong(args.get(1))
                    : registers.get(args.get(1));
            registers.put(args.get(0), newValue);
        }

        void add(Map<String, Long> registers, List<String> args) {
            long current = registers.get(args.get(0));
            long newValue = args.get(1).matches("-?\\d+") ? Long.parseLong(args.get(1))
                    : registers.get(args.get(1));
            registers.put(args.get(0), current + newValue);
        }

        void mul(Map<String, Long> registers, List<String> args) {
            long current = registers.get(args.get(0));
            long newValue = args.get(1).matches("-?\\d+") ? Long.parseLong(args.get(1))
                    : registers.get(args.get(1));
            registers.put(args.get(0), current * newValue);
        }

        void mod(Map<String, Long> registers, List<String> args) {
            long current = registers.get(args.get(0));
            long newValue = args.get(1).matches("-?\\d+") ? Long.parseLong(args.get(1))
                    : registers.get(args.get(1));
            registers.put(args.get(0), current % newValue);
        }

        long jgz(Map<String, Long> registers, List<String> args) {
            long value = args.get(0).matches("-?\\d+") ? Long.parseLong(args.get(0))
                    : registers.get(args.get(0));
            long offset = args.get(1).matches("-?\\d+") ? Long.parseLong(args.get(1))
                    : registers.get(args.get(1));
            return value > 0 ? offset : 1;
        }
    }

    private static class Instruction {
        String name;
        List<String> args;

        public Instruction(String name, List<String> args) {
            this.name = name;
            this.args = args;
        }

        @Override
        public String toString() {
            return "Instruction{" + "name='" + name + '\'' + ", args=" + args + '}';
        }
    }
}
