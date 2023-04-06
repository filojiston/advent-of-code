public class Main {
    private static final int GENERATOR_A_FACTOR = 16807;
    private static final int GENERATOR_B_FACTOR = 48271;

    private static final long GENERATOR_A_START = 116;
    private static final long GENERATOR_B_START = 299;

    private static final long PART1_LOOPS = 40000000;
    private static final long PART2_LOOPS = 5000000;

    public static void main(String[] args) {
        System.out.println("Part 1: " + part1(GENERATOR_A_START, GENERATOR_B_START));
        System.out.println("Part 2: " + part2(GENERATOR_A_START, GENERATOR_B_START));
    }

    public static int part1(long valueGeneratorA, long valueGeneratorB) {
        int count = 0;
        for (int i = 0; i < PART1_LOOPS; i++) {
            valueGeneratorA =
                    calculateValueForGenerator(valueGeneratorA, GENERATOR_A_FACTOR, false, null);
            valueGeneratorB =
                    calculateValueForGenerator(valueGeneratorB, GENERATOR_B_FACTOR, false, null);
            if (isLast16BitsEqual(valueGeneratorA, valueGeneratorB)) {
                count++;
            }
        }
        return count;
    }

    public static int part2(long valueGeneratorA, long valueGeneratorB) {
        int count = 0;
        for (int i = 0; i < PART2_LOOPS; i++) {
            valueGeneratorA =
                    calculateValueForGenerator(valueGeneratorA, GENERATOR_A_FACTOR, true, 4);
            valueGeneratorB =
                    calculateValueForGenerator(valueGeneratorB, GENERATOR_B_FACTOR, true, 8);
            if (isLast16BitsEqual(valueGeneratorA, valueGeneratorB)) {
                count++;
            }
        }
        return count;
    }

    private static long calculateValueForGenerator(long startValue, int factor,
            boolean shouldCheckMultiple, Integer multiple) {
        long value = (startValue * factor) % 2147483647;
        if (shouldCheckMultiple) {
            assert multiple != null;
            while (value % multiple != 0) {
                value = (value * factor) % 2147483647;
            }
        }
        return value;
    }

    private static boolean isLast16BitsEqual(long a, long b) {
        return (a & 0xFFFF) == (b & 0xFFFF);
    }
}
