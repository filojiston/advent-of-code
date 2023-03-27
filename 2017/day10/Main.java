import java.util.Arrays;
import java.util.List;
import java.util.stream.IntStream;

public class Main {

    private static final List<Integer> INPUT =
            Arrays.asList(199, 0, 255, 136, 174, 254, 227, 16, 51, 85, 1, 2, 22, 17, 7, 192);

    public static void main(String[] args) {
        System.out.println("Part 1: " + part1());
        System.out.println("Part 2: " + part2());
    }

    private static int part1() {
        int[] list = IntStream.range(0, 256).toArray();
        for (int i = 0, skip = 0, pos = 0; i < INPUT.size(); i++, skip++) {
            int length = INPUT.get(i);
            reverse(list, pos, length);
            pos = (pos + length + skip) % list.length;
        }

        return list[0] * list[1];
    }

    private static String part2() {
        int[] list = IntStream.range(0, 256).toArray();
        int[] lengths = getLengths();

        for (int i = 0, skip = 0, pos = 0; i < 64; i++) {
            for (int length : lengths) {
                reverse(list, pos, length);
                pos = (pos + length + skip) % list.length;
                skip++;
            }
        }

        return calculateHash(list);
    }

    private static void reverse(int[] list, int pos, int length) {
        for (int i = 0; i < length / 2; i++) {
            int a = (pos + i) % list.length;
            int b = (pos + length - 1 - i) % list.length;
            int tmp = list[a];
            list[a] = list[b];
            list[b] = tmp;
        }
    }

    private static String calculateHash(int[] list) {
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < 16; i++) {
            int xor = 0;
            for (int j = 0; j < 16; j++) {
                xor ^= list[i * 16 + j];
            }
            sb.append(String.format("%02x", xor));
        }
        return sb.toString();
    }

    private static int[] getLengths() {
        String input = String.join(",",
                INPUT.toString().substring(1, INPUT.toString().length() - 1).split(", "));
        int[] lengths = input.chars().toArray();
        lengths =
                IntStream.concat(IntStream.of(lengths), IntStream.of(17, 31, 73, 47, 23)).toArray();
        return lengths;
    }
}
