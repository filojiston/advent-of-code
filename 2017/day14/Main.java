import java.util.stream.IntStream;

public class Main {
    private static final String INPUT = "oundnydw";

    public static void main(String[] args) {
        boolean[][] grid = new boolean[128][128];
        initializeGrid(grid);
        System.out.println("Part 1: " + part1(grid));
        System.out.println("Part 2: " + part2(grid));
    }

    private static int part1(final boolean[][] grid) {
        int count = 0;
        for (int i = 0; i < 128; i++) {
            for (int j = 0; j < 128; j++) {
                count += grid[i][j] ? 1 : 0;
            }
        }
        return count;
    }

    private static int part2(boolean[][] grid) {
        int regions = 0;
        for (int i = 0; i < 128; i++) {
            for (int j = 0; j < 128; j++) {
                if (grid[i][j]) {
                    regions++;
                    removeRegion(grid, i, j);
                }
            }
        }
        return regions;
    }

    private static void removeRegion(boolean[][] grid, int i, int j) {
        if (i < 0 || i >= 128 || j < 0 || j >= 128 || !grid[i][j]) {
            return;
        }
        grid[i][j] = false;
        removeRegion(grid, i - 1, j);
        removeRegion(grid, i + 1, j);
        removeRegion(grid, i, j - 1);
        removeRegion(grid, i, j + 1);
    }

    private static void initializeGrid(boolean[][] grid) {
        for (int i = 0; i < 128; i++) {
            String hash = calculateHash(INPUT + "-" + i);
            boolean[] binary = hashToBinary(hash);
            for (int j = 0; j < 128; j++) {
                grid[i][j] = binary[j];
            }
        }
    }

    private static String calculateHash(String input) {
        int[] list = IntStream.range(0, 256).toArray();
        int[] lengths = stringToIntArray(input);
        lengths = IntStream.concat(IntStream.of(lengths), IntStream.of(17, 31, 73, 47, 23))
                .toArray();;

        // calculate sparse hash
        for (int i = 0, skip = 0, pos = 0; i < 64; i++) {
            for (int length : lengths) {
                reverse(list, pos, length);
                pos = (pos + length + skip) % list.length;
                skip++;
            }
        }

        return denseHash(list);
    }

    private static String denseHash(int[] list) {
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

    private static boolean[] hashToBinary(String hash) {
        boolean[] result = new boolean[hash.length() * 4];
        for (int i = 0; i < hash.length(); i++) {
            int value = Integer.parseInt(hash.substring(i, i + 1), 16);
            for (int j = 0; j < 4; j++) {
                result[i * 4 + j] = (value & (1 << (3 - j))) != 0;
            }
        }

        return result;
    }

    private static int[] stringToIntArray(String input) {
        int[] list = new int[input.length()];
        for (int i = 0; i < input.length(); i++) {
            list[i] = input.charAt(i);
        }
        return list;
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
}
