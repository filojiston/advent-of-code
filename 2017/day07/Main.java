import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.stream.Collectors;

public class Main {
    public static void main(String[] args) {
        String input;
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException exception) {
            System.out.println("Error reading input file");
            return;
        }

        System.out.println("Part 1: " + part1(input));
        System.out.println("Part 2: ");
        part2(input);
    }

    private static String part1(String input) {
        Set<String> discs = new HashSet<>();
        Set<String> leafs = new HashSet<>();

        for (String line : input.split("\n")) {
            String[] parts = line.split(" ");
            String disc = parts[0];
            discs.add(disc);

            if (parts.length > 2) {
                for (int i = 3; i < parts.length; i++) {
                    String leaf = parts[i].replace(",", "");
                    leafs.add(leaf);
                }
            }
        }

        discs.removeAll(leafs);
        return discs.iterator().next();
    }

    private static void part2(String input) {
        Set<Disc> discs = new HashSet<>();

        for (String line : input.split("\n")) {
            Disc disc = new Disc();
            String[] parts = line.split(" ");
            disc.name = parts[0];
            disc.weight = Integer.parseInt(parts[1].replace("(", "").replace(")", ""));
            discs.add(disc);
        }

        for (String line : input.split("\n")) {
            String[] parts = line.split(" ");
            if (parts.length > 2) {
                Disc parent = discs.stream().filter(d -> d.name.equals(parts[0])).findFirst().get();
                parent.children = new ArrayList<>();
                for (int i = 3; i < parts.length; i++) {
                    String childName = parts[i].replace(",", "");
                    Disc child =
                            discs.stream().filter(d -> d.name.equals(childName)).findFirst().get();
                    parent.children.add(child);
                }
            }
        }
        Disc root = discs.stream().filter(disc -> disc.name.equals(part1(input))).findFirst().get();
        dfs(root);
    }

    private static int dfs(Disc disc) {
        if (disc.children == null) {
            return disc.weight;
        }

        int[] weights = new int[disc.children.size()];
        for (int i = 0; i < disc.children.size(); i++) {
            weights[i] = dfs(disc.children.get(i));
        }

        int sum = disc.weight;
        for (int weight : weights) {
            sum += weight;
        }

        if (weights.length > 0) {
            int[] sorted = Arrays.copyOf(weights, weights.length);
            Arrays.sort(sorted);

            if (sorted[0] != sorted[sorted.length - 1] && sorted[0] != sorted[1]) {
                int diff = sorted[1] - sorted[0];
                int index = Arrays.stream(weights).boxed().collect(Collectors.toList())
                        .indexOf(sorted[sorted.length - 1]);
                Disc child = disc.children.get(index);
                System.out.println("Disc " + child.name + "has weight: " + child.weight + " and should have weight: " + (child.weight + diff));
            } else if (sorted[0] != sorted[sorted.length - 1]){
                int diff = sorted[sorted.length - 1] - sorted[0];
                int index = Arrays.stream(weights).boxed().collect(Collectors.toList())
                        .indexOf(sorted[sorted.length - 1]);
                Disc child = disc.children.get(index);
                System.out.println("Disc " + child.name + "has weight: " + child.weight + " and should have weight: " + (child.weight - diff));
            }
        }

        return sum;
    }

    private static class Disc {
        private String name;
        private int weight;
        private List<Disc> children;
    }
}
