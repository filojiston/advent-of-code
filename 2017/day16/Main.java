import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.Arrays;
import java.util.List;
import java.util.stream.Collectors;
import java.util.stream.IntStream;

public class Main {

    private static final int DANCE_REPEAT_COUNT = 1000000000;

    public static void main(String[] args) {
        String input = "";
        try {
            input = new String(Files.readAllBytes(Paths.get("input.txt")));
        } catch (IOException e) {
            System.err.println("Error reading input file");
        }

        char[] programs = initPrograms();
        List<Move> moves = initMoves(input);
        System.out.println("Part 1: " + applyMoves(programs, moves, 1));
        programs = initPrograms();
        System.out.println("Part 2: " + applyMoves(programs, moves, DANCE_REPEAT_COUNT));
    }

    private static char[] initPrograms() {
        return IntStream.range(0, 16).mapToObj(i -> (char) ('a' + i))
                .collect(StringBuilder::new, StringBuilder::append, StringBuilder::append)
                .toString().toCharArray();
    }

    private static List<Move> initMoves(String input) {
        return Arrays.stream(input.split(",")).map(Move::parse).collect(Collectors.toList());
    }

    private static String applyMoves(char[] programs, List<Move> moves, int count) {
        String initial = new String(programs);
        for (int i = 0; i < count; i++) {
            moves.forEach(move -> move.execute(programs));
            if (new String(programs).equals(initial)) {
                i += (Math.floor(count / (i + 1)) - 1) * (i + 1);
            }
        }
        return new String(programs);
    }

    private static class Move {
        private static final String SPIN = "s";
        private static final String EXCHANGE = "x";
        private static final String PARTNER = "p";

        private String type;
        private String args;

        private Move(String type, String args) {
            this.type = type;
            this.args = args;
        }

        public static Move parse(String s) {
            if (s.startsWith(SPIN)) {
                return new Move(SPIN, s.substring(1));
            } else if (s.startsWith(EXCHANGE)) {
                return new Move(EXCHANGE, s.substring(1));
            } else if (s.startsWith(PARTNER)) {
                return new Move(PARTNER, s.substring(1));
            }
            return null;
        }

        public void execute(char[] programs) {
            switch (type) {
                case SPIN:
                    spin(programs);
                    break;
                case EXCHANGE:
                    exchange(programs);
                    break;
                case PARTNER:
                    partner(programs);
                    break;
            }
        }

        private void spin(char[] programs) {
            int n = Integer.parseInt(args);
            char[] newPrograms = new char[programs.length];
            for (int i = 0; i < programs.length; i++) {
                newPrograms[(i + n) % programs.length] = programs[i];
            }
            System.arraycopy(newPrograms, 0, programs, 0, programs.length);
        }

        private void exchange(char[] programs) {
            String[] args = this.args.split("/");
            int a = Integer.parseInt(args[0]);
            int b = Integer.parseInt(args[1]);
            char temp = programs[a];
            programs[a] = programs[b];
            programs[b] = temp;
        }

        private void partner(char[] programs) {
            String[] args = this.args.split("/");
            char a = args[0].charAt(0);
            char b = args[1].charAt(0);
            int aIndex = -1;
            int bIndex = -1;
            for (int i = 0; i < programs.length; i++) {
                if (programs[i] == a) {
                    aIndex = i;
                } else if (programs[i] == b) {
                    bIndex = i;
                }
            }
            programs[aIndex] = b;
            programs[bIndex] = a;
        }
    }
}
