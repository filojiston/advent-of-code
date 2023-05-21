import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.HashSet;
import java.util.Set;

public class Main {
    public static void main(String[] args) {
        String input = "";
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException e) {
            System.err.println("Error reading input file");
        }

        Set<Pair> ports = parseInput(input);
        Set<Set<Pair>> bridges = constructAllPossibleBridges(ports, 0, new HashSet<>());
        System.out.println("Part 1: " + part1(bridges));
        System.out.println("Part 1 (with dfs): " + part1WithDFS(ports));
        System.out.println("Part 2: " + part2(bridges));
        System.out.println("Part 2 (with dfs): " + part2WithDFS(ports));
    }

    private static int part1(Set<Set<Pair>> bridges) {
        return bridges.stream().mapToInt(bridge -> bridge.stream().mapToInt(p -> p.x + p.y).sum())
                .max().getAsInt();
    }

    private static int part2(Set<Set<Pair>> bridges) {
        int maxLength = bridges.stream().mapToInt(Set::size).max().getAsInt();
        return bridges.stream().filter(bridge -> bridge.size() == maxLength)
                .mapToInt(bridge -> bridge.stream().mapToInt(p -> p.x + p.y).sum()).max()
                .getAsInt();
    }

    private static int part1WithDFS(Set<Pair> ports) {
        return part1DFS(ports, 0, 0);
    }

    private static int part2WithDFS(Set<Pair> ports) {
        int[] result = part2DFS(ports, 0, 0, 0);
        return result[1];
    }

    private static int part1DFS(Set<Pair> ports, int port, int strength) {
        int maxStrength = strength;
        for (Pair p : ports) {
            if (p.x == port || p.y == port) {
                Set<Pair> newPorts = new HashSet<>(ports);
                newPorts.remove(p);
                maxStrength = Math.max(maxStrength,
                        part1DFS(newPorts, p.x == port ? p.y : p.x, strength + p.x + p.y));
            }
        }
        return maxStrength;
    }

    private static int[] part2DFS(Set<Pair> ports, int port, int strength, int length) {
        int maxLength = length;
        int maxStrength = strength;
        for (Pair p : ports) {
            if (p.x == port || p.y == port) {
                Set<Pair> newPorts = new HashSet<>(ports);
                newPorts.remove(p);
                int newLength = length + 1;
                int newStrength = strength + p.x + p.y;
                int[] result = part2DFS(newPorts, p.x == port ? p.y : p.x, newStrength, newLength);
                if (result[0] > maxLength) {
                    maxLength = result[0];
                    maxStrength = result[1];
                } else if (result[0] == maxLength) {
                    maxStrength = Math.max(maxStrength, result[1]);
                }
            }
        }
        return new int[] {maxLength, maxStrength};
    }

    private static Set<Set<Pair>> constructAllPossibleBridges(Set<Pair> ports, int port,
            Set<Pair> bridge) {
        Set<Set<Pair>> bridges = new HashSet<>();
        for (Pair p : ports) {
            if (p.x == port || p.y == port) {
                Set<Pair> newPorts = new HashSet<>(ports);
                newPorts.remove(p);
                Set<Pair> newBridge = new HashSet<>(bridge);
                newBridge.add(p);
                bridges.addAll(
                        constructAllPossibleBridges(newPorts, p.x == port ? p.y : p.x, newBridge));
            }
        }
        bridges.add(bridge);
        return bridges;
    }

    private static Set<Pair> parseInput(String input) {
        Set<Pair> ports = new HashSet<>();
        input.lines().forEach(line -> {
            String[] parts = line.split("/");
            ports.add(new Pair(Integer.parseInt(parts[0]), Integer.parseInt(parts[1])));
        });
        return ports;
    }


    private static class Pair {
        int x;
        int y;

        public Pair(int x, int y) {
            this.x = x;
            this.y = y;
        }

        @Override
        public String toString() {
            return "(" + x + ", " + y + ")";
        }
    }
}

