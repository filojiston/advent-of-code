package day03;

import java.util.List;

// i couldn't be able to solve it programatically. so i cheated :( i don't like math heavy problems.
// thx for the solution:
// https://github.com/if0nz/advent-of-code-2017/blob/master/src/main/java/it/ifonz/puzzle/Day03.java
public class Main {

    private static final int INPUT = 325489;

    public static void main(String[] args) {
        System.out.println("Part 1: " + part1());
        System.out.println("Part 2: " + part2());
    }

    private static int part1() {
        int center = (int) Math.ceil((Math.sqrt(INPUT) - 1) / 2);
        int distanceP1 = center;
        long startingPoint = (long) Math.pow(2 * center - 1, 2);
        int k = (int) Math.ceil((INPUT - startingPoint) / (2 * center));
        startingPoint += 2 * k * center;
        distanceP1 += Math.abs(INPUT - (startingPoint + center));
        return distanceP1;
    }

    private static int part2() {
        int center = (int) Math.ceil((Math.sqrt(INPUT) - 1) / 2);

        int size = 2 * center + 1;
        int[][] grid = new int[size][size];

        grid[center][center] = 1;

        List<Direction> directions =
                List.of(Direction.RIGHT, Direction.UP, Direction.LEFT, Direction.DOWN);
        Direction direction = Direction.RIGHT;
        int vertical = center, horizontal = center;
        int steps = 1;
        while (grid[horizontal][vertical] <= INPUT) {
            for (int i = 0; i < 2 && grid[horizontal][vertical] <= INPUT; i++) {
                for (int j = 0; j < steps && grid[horizontal][vertical] <= INPUT; j++) {
                    switch (direction) {
                        case RIGHT: {
                            ++horizontal;
                            break;
                        }
                        case UP: {
                            --vertical;
                            break;
                        }
                        case LEFT: {
                            --horizontal;
                            break;
                        }
                        case DOWN: {
                            ++vertical;
                            break;
                        }
                    }
                    for (int x = -1; x <= 1; x++) {
                        for (int y = -1; y <= 1; y++) {
                            grid[horizontal][vertical] +=
                                    (boundaryTest(size, horizontal, vertical, x, y)
                                            ? grid[horizontal + x][vertical + y]
                                            : 0);
                        }
                    }
                }
                direction = directions.get((directions.indexOf(direction) + 1) % directions.size());
            }
            steps++;
        }

        return grid[horizontal][vertical];
    }

    private static boolean boundaryTest(int size, int horizontal, int vertical, int x, int y) {
        return !((x == 0 && y == 0) || horizontal + x < 0 || horizontal + x >= size
                || vertical + y < 0 || vertical + y > size - 1);
    }

    private static enum Direction {
        UP, DOWN, LEFT, RIGHT
    }

}
