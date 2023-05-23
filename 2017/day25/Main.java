public class Main {
    private static final int INITIAL_TAPE_SIZE = 1000;
    private static final int INITIAL_TAPE_INDEX = INITIAL_TAPE_SIZE / 2;
    private static final int STEPS = 12172063;

    public static void main(String[] args) {
        boolean[] tape = new boolean[INITIAL_TAPE_SIZE];
        State state = State.A;
        int index = INITIAL_TAPE_INDEX;
        for (int i = 0; i < STEPS; i++) {
            try {
                Result result = step(tape, index, state);
                tape = result.getTape();
                index = result.getIndex();
                state = result.getState();
            } catch (IndexOutOfBoundsException exception) {
                System.out.println("Index out of bounds, expanding tape");
                boolean[] newTape = new boolean[tape.length * 2];
                System.arraycopy(tape, 0, newTape, tape.length / 2, tape.length);
                tape = newTape;
                index += tape.length / 4;
                i--;
            }
        }

        int checksum = 0;
        for (boolean value : tape) {
            if (value) {
                checksum++;
            }
        }
        System.out.println("Checksum: " + checksum);
    }

    private static Result step(boolean[] tape, int index, State state) {
        switch (state) {
            case A:
                if (!tape[index]) {
                    tape[index] = true;
                    index++;
                    state = State.B;
                } else {
                    tape[index] = false;
                    index--;
                    state = State.C;
                }
                break;
            case B:
                if (!tape[index]) {
                    tape[index] = true;
                    index--;
                    state = State.A;
                } else {
                    tape[index] = true;
                    index--;
                    state = State.D;
                }
                break;
            case C:
                if (!tape[index]) {
                    tape[index] = true;
                    index++;
                    state = State.D;
                } else {
                    tape[index] = false;
                    index++;
                    state = State.C;
                }
                break;
            case D:
                if (!tape[index]) {
                    tape[index] = false;
                    index--;
                    state = State.B;
                } else {
                    tape[index] = false;
                    index++;
                    state = State.E;
                }
                break;
            case E:
                if (!tape[index]) {
                    tape[index] = true;
                    index++;
                    state = State.C;
                } else {
                    tape[index] = true;
                    index--;
                    state = State.F;
                }
                break;
            case F:
                if (!tape[index]) {
                    tape[index] = true;
                    index--;
                    state = State.E;
                } else {
                    tape[index] = true;
                    index++;
                    state = State.A;
                }
                break;
            default:
                throw new IllegalArgumentException("Unknown state: " + state);
        }
        return new Result(tape, index, state);
    }

    private static class Result {
        private final boolean[] tape;
        private final int index;
        private final State state;

        public Result(boolean[] tape, int index, State state) {
            this.tape = tape;
            this.index = index;
            this.state = state;
        }

        public boolean[] getTape() {
            return tape;
        }

        public int getIndex() {
            return index;
        }

        public State getState() {
            return state;
        }
    }

    private static enum State {
        A, B, C, D, E, F
    }
}
