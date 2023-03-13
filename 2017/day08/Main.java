import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.HashMap;
import java.util.Map;
import java.util.function.Function;

public class Main {
    private static final Map<String, Function<Integer, Integer>> CONDITIONS_MAP =
            createConditions();

    public static void main(String[] args) {
        String input;
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (Exception e) {
            System.err.println("Error reading input file");
            return;
        }

        solve(input);
    }

    private static void solve(String input) {
        Map<String, Integer> registers = new HashMap<>();
        int max = 0;

        for (String line : input.split("\n")) {
            Instruction inst = Instruction.parse(line);

            registers.putIfAbsent(inst.reg1, 0);
            registers.putIfAbsent(inst.reg2, 0);

            if (CONDITIONS_MAP.get(inst.cond)
                    .apply(registers.get(inst.reg2) - inst.condValue) == 1) {
                if (inst.op.equals("inc")) {
                    registers.put(inst.reg1, registers.get(inst.reg1) + inst.value);
                } else {
                    registers.put(inst.reg1, registers.get(inst.reg1) - inst.value);
                }
                if (registers.get(inst.reg1) > max) {
                    max = registers.get(inst.reg1);
                }
            }
        }

        System.out.println(
                "Part 1: " + registers.values().stream().max(Integer::compareTo).orElse(0));
        System.out.println("Part 2: " + max);
    }

    private static Map<String, Function<Integer, Integer>> createConditions() {
        Map<String, Function<Integer, Integer>> conditions = new HashMap<>();
        conditions.put(">", (x) -> x > 0 ? 1 : 0);
        conditions.put("<", (x) -> x < 0 ? 1 : 0);
        conditions.put(">=", (x) -> x >= 0 ? 1 : 0);
        conditions.put("<=", (x) -> x <= 0 ? 1 : 0);
        conditions.put("==", (x) -> x == 0 ? 1 : 0);
        conditions.put("!=", (x) -> x != 0 ? 1 : 0);
        return conditions;
    }

    private static class Instruction {
        String reg1;
        String op;
        int value;
        String reg2;
        String cond;
        int condValue;

        Instruction(String reg1, String op, int value, String reg2, String cond, int condValue) {
            this.reg1 = reg1;
            this.op = op;
            this.value = value;
            this.reg2 = reg2;
            this.cond = cond;
            this.condValue = condValue;
        }

        static Instruction parse(String line) {
            String[] tokens = line.split(" ");
            String reg1 = tokens[0];
            String op = tokens[1];
            int value = Integer.parseInt(tokens[2]);
            String reg2 = tokens[4];
            String cond = tokens[5];
            int condValue = Integer.parseInt(tokens[6]);
            return new Instruction(reg1, op, value, reg2, cond, condValue);
        }
    }
}
