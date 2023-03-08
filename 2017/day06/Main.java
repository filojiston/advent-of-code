import java.util.ArrayList;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Set;

public class Main {

    private static final List<Integer> INPUT = List.of(11, 11, 13, 7, 0, 15, 5, 5, 4, 4, 1, 1, 7, 1, 15, 11);

    public static void main(String[] args) {
        System.out.println("Part 1: " + part1());
        System.out.println("Part 2: " + part2());
    }

    private static int part1() {
        List<Integer> banks = new ArrayList<>(INPUT);
        Set<String> seen = new HashSet<>();

        int cycles = 0;
        while (seen.add(banks.toString())) {
            redistribute(banks);
            cycles++;
        }

        return cycles;
    }

    private static int part2() {
        List<Integer> banks = new ArrayList<>(INPUT);
        HashMap<String, Integer> seen = new HashMap<>();

        int cycles = 0;
        while (!seen.containsKey(banks.toString())) {
            seen.put(banks.toString(), cycles);
            redistribute(banks);
            cycles++;
        }

        return seen.size() - seen.get(banks.toString());
    }

    private static void redistribute(List<Integer> banks) {
        int max = banks.stream().mapToInt(Integer::intValue).max().getAsInt();
        int index = banks.indexOf(max);
        banks.set(index, 0);

        while (max > 0) {
            index = (index + 1) % banks.size();
            banks.set(index, banks.get(index) + 1);
            max--;
        }
    }
}
