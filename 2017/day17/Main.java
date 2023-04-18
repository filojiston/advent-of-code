public class Main {
    private static final int INPUT = 345;

    public static void main(String[] args) {
        System.out.println("Part 1: " + part1(INPUT));
        System.out.println("Part 2: " + part2(INPUT));
    }

    private static int part1(int input) {
        int[] buffer = new int[2018];
        buffer[0] = 0;
        int pos = 0;
        for (int i = 1; i < buffer.length; i++) {
            pos = (pos + input) % i + 1;
            System.arraycopy(buffer, pos, buffer, pos + 1, i - pos);
            buffer[pos] = i;
        }
        return buffer[pos + 1];
    }

    private static int part2(int input) {
        int pos = 0;
        int result = -1;
        for (int i = 1; i <= 50_000_000; i++) {
            pos = (pos + input) % i + 1;
            if (pos == 1) {
                result = i;
            }
        }
        return result;
    }
}
