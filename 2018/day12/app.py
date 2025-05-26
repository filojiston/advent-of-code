from util.util import read_input

inp = list(filter(lambda line: line, read_input("day12/input.txt")))


def parse_rules(rule_lines):
    """Parses rules into a dictionary."""
    return dict(line.split(" => ") for line in rule_lines)


def get_window(state, idx):
    """Gets a 5-character window around the given index with padding."""
    window = state[idx - 2 : idx + 3]
    return window.rjust(5, ".") if len(window) < 5 else window.ljust(5, ".")


def pad_state(state, pad_size=4):
    """Pads the state with dots on both sides and returns updated state and offset shift."""
    return "...." + state + "....", -pad_size


def trim_state(state):
    """Trims leading and trailing dots and returns trimmed state and first plant index."""
    first = state.find("#")
    last = state.rfind("#")
    return state[first : last + 1], first


def calculate_plant_sum(state, offset):
    """Calculates the sum of pot numbers with plants."""
    return sum(i + offset for i, c in enumerate(state) if c == "#")


def simulate(rules, initial_state, max_generations, detect_stabilization=False):
    """Simulates the plant growth with optional stabilization detection."""
    state = initial_state
    offset = 0
    seen_patterns = {}

    for generation in range(max_generations):
        state, offset_shift = pad_state(state)
        offset += offset_shift

        new_state = "".join(
            rules.get(get_window(state, i), ".") for i in range(len(state))
        )

        state, trim_shift = trim_state(new_state)
        offset += trim_shift

        plant_sum = calculate_plant_sum(state, offset)

        if detect_stabilization:
            pattern = state
            if pattern in seen_patterns:
                _, prev_sum = seen_patterns[pattern]
                delta_sum = plant_sum - prev_sum
                remaining = max_generations - generation - 1
                return plant_sum + (remaining * delta_sum)
            seen_patterns[pattern] = (generation, plant_sum)

    return calculate_plant_sum(state, offset)


def part1(rules, initial_state, generations=20):
    return simulate(rules, initial_state, generations)


def part2(rules, initial_state, generations=50000000000):
    return simulate(rules, initial_state, generations, detect_stabilization=True)


initial_state = inp.pop(0).lstrip("initial state: ")
rules = parse_rules(inp)

print(f"Part 1: {part1(rules, initial_state)}")
print(f"Part 2: {part2(rules, initial_state)}")
