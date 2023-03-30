import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashSet;
import java.util.List;
import java.util.Set;
import java.util.Stack;
import java.util.stream.Collectors;

public class Main {
    public static void main(String[] args) {
        String input;
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (Exception e) {
            System.err.println("Error reading input file. Stack trace: " + e.getStackTrace());
            return;
        }

        Graph<Integer> graph = initGraph(input);

        System.out.println("Part 1: " + part1(graph));
        System.out.println("Part 2: " + part2(graph));
    }

    private static int part1(Graph<Integer> graph) {
        return dfs(graph, 0).size();
    }

    private static int part2(Graph<Integer> graph) {
        Set<Integer> allVertices = new HashSet<>();
        graph.getAllVertices().forEach(allVertices::add);

        Set<Set<Integer>> groups = new HashSet<>();
        while (!allVertices.isEmpty()) {
            Integer start = allVertices.iterator().next();
            Set<Integer> group = dfs(graph, start);
            groups.add(group);
            allVertices.removeAll(group);
        }

        return groups.size();
    }

    private static List<Integer> parseVertices(String input) {
        return Arrays.stream(input.split("\n")).map(line -> line.split(" <-> ")[0])
                .map(Integer::parseInt).collect(Collectors.toList());
    }

    private static List<Pair<Integer, Integer>> parseEdges(String input) {
        List<Pair<Integer, Integer>> edges = new ArrayList<>();
        String[] lines = input.split("\n");
        for (String line : lines) {
            Integer vertex = Integer.parseInt(line.split(" <-> ")[0]);
            List<Integer> edgesOfVertex = Arrays.stream(line.split(" <-> ")[1].split(", "))
                    .map(Integer::parseInt).collect(Collectors.toList());
            edgesOfVertex.forEach(edge -> edges.add(new Pair<>(vertex, edge)));
        }

        return edges;
    }

    private static Set<Integer> dfs(Graph<Integer> graph, Integer start) {
        Set<Integer> visited = new HashSet<>();
        Stack<Integer> stack = new Stack<>();
        stack.push(start);

        while (!stack.isEmpty()) {
            Integer vertex = stack.pop();
            if (!visited.contains(vertex)) {
                visited.add(vertex);
                graph.getNeighbors(vertex).forEach(stack::push);
            }
        }

        return visited;
    }

    private static Graph<Integer> initGraph(String input) {
        Graph<Integer> graph = new Graph<>();

        List<Integer> vertices = parseVertices(input);
        vertices.forEach(graph::addVertex);

        List<Pair<Integer, Integer>> edges = parseEdges(input);
        edges.forEach(edge -> graph.addEdge(edge.getKey(), edge.getValue()));

        return graph;
    }
}
