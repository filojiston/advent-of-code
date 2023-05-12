import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;

public class Main {
    private static final int PART1_ITERATIONS = 10000;
    private static final int PART2_ITERATIONS = 10000000;

    public static void main(String[] args) {
        String input = "";
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException exception) {
            System.err.println("Error reading input file");
        }

        Grid grid = parseInput(input);
        System.out.println("Part 1: " + part1(grid));
        grid = parseInput(input);
        System.out.println("Part 2: " + part2(grid));
    }

    private static int part1(Grid grid) {
        int count = 0;
        for (int i = 0; i < PART1_ITERATIONS; i++) {
            if (grid.getCurrentNode().isInfected()) {
                grid.turnVirus("right");
                grid.getCurrentNode().clean();
            } else {
                grid.turnVirus("left");
                grid.getCurrentNode().infect();
                count++;
            }
            grid.moveVirus();
        }
        return count;
    }

    private static int part2(Grid grid) {
        int count = 0;
        for (int i = 0; i < PART2_ITERATIONS; i++) {
            switch (grid.getCurrentNode().state) {
                case CLEAN:
                    grid.turnVirus("left");
                    grid.getCurrentNode().weaken();
                    break;
                case WEAKENED:
                    grid.getCurrentNode().infect();
                    count++;
                    break;
                case INFECTED:
                    grid.turnVirus("right");
                    grid.getCurrentNode().flag();
                    break;
                case FLAGGED:
                    grid.turnVirus("reverse");
                    grid.getCurrentNode().clean();
                    break;
            }
            grid.moveVirus();
        }
        return count;
    }

    private static Grid parseInput(String input) {
        Node[][] grid = input.lines().map(line -> line.chars().mapToObj(c -> {
            switch (c) {
                case '#':
                    return new Node(State.INFECTED);
                case '.':
                    return new Node(State.CLEAN);
                default:
                    throw new IllegalArgumentException("Invalid character in input");
            }
        }).toArray(Node[]::new)).toArray(Node[][]::new);
        return new Grid(grid);
    }

    private static class Grid {
        Node[][] grid;
        Point virusLocation;
        char virusDirection;

        public Grid(Node[][] grid) {
            this.grid = grid;
            virusLocation = new Point(grid.length / 2, grid.length / 2);
            virusDirection = 'U';
        }

        public Node getCurrentNode() {
            return grid[virusLocation.y][virusLocation.x];
        }

        public void turnVirus(String direction) {
            switch (virusDirection) {
                case 'U':
                    virusDirection =
                            direction.equals("right") ? 'R' : direction.equals("left") ? 'L' : 'D';
                    break;
                case 'D':
                    virusDirection =
                            direction.equals("right") ? 'L' : direction.equals("left") ? 'R' : 'U';
                    break;
                case 'L':
                    virusDirection =
                            direction.equals("right") ? 'U' : direction.equals("left") ? 'D' : 'R';
                    break;
                case 'R':
                    virusDirection =
                            direction.equals("right") ? 'D' : direction.equals("left") ? 'U' : 'L';
                    break;
            }
        }

        public void moveVirus() {
            switch (virusDirection) {
                case 'U':
                    virusLocation.y--;
                    break;
                case 'D':
                    virusLocation.y++;
                    break;
                case 'L':
                    virusLocation.x--;
                    break;
                case 'R':
                    virusLocation.x++;
                    break;
            }
            if (virusLocation.x < 0 || virusLocation.x >= grid.length || virusLocation.y < 0
                    || virusLocation.y >= grid.length) {
                expandGrid();
            }
        }

        private void expandGrid() {
            Node[][] newGrid = new Node[grid.length + 2][grid.length + 2];
            for (int i = 0; i < newGrid.length; i++) {
                for (int j = 0; j < newGrid.length; j++) {
                    if (i == 0 || i == newGrid.length - 1 || j == 0 || j == newGrid.length - 1) {
                        newGrid[i][j] = new Node(State.CLEAN);
                    } else {
                        newGrid[i][j] = grid[i - 1][j - 1];
                    }
                }
            }
            grid = newGrid;
            virusLocation.x++;
            virusLocation.y++;
        }
    }

    private static class Node {
        State state;

        public Node(State state) {
            this.state = state;
        }

        public boolean isInfected() {
            return state == State.INFECTED;
        }

        public void infect() {
            state = State.INFECTED;
        }

        public void clean() {
            state = State.CLEAN;
        }

        public void weaken() {
            state = State.WEAKENED;
        }

        public void flag() {
            state = State.FLAGGED;
        }
    }

    private static enum State {
        CLEAN, WEAKENED, INFECTED, FLAGGED
    }

    private static class Point {
        int x;
        int y;

        public Point(int x, int y) {
            this.x = x;
            this.y = y;
        }
    }
}
