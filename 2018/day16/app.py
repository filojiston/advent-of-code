from util.util import read_input_without_strip
from itertools import islice, chain, repeat
from copy import deepcopy


# Helper method to split a list into chunks with padding support
def chunk_pad(it, size, padval=None):
    it = chain(iter(it), repeat(padval))
    return iter(lambda: tuple(islice(it, size)), (padval,) * size)


INPUT = list(filter(lambda line: line, read_input_without_strip("day16/input.txt")))
INPUT_SPLIT = "-" * 30


class Sample:
    def __init__(self, before, instruction, after):
        self.before = before
        self.instruction = instruction
        self.after = after


def parse_samples(inp):
    data = list(chunk_pad(inp[: inp.index(INPUT_SPLIT)], 3))
    samples = []

    for chunk in data:
        before = [int(x) for x in chunk[0].lstrip("Before: [").rstrip("]").split(", ")]
        instruction = [int(x) for x in chunk[1].split(" ")]
        after = [int(x) for x in chunk[2].lstrip("After:  [").rstrip("]").split(", ")]
        samples.append(Sample(before, instruction, after))

    return samples


def parse_test_program(inp):
    data = list(inp[inp.index(INPUT_SPLIT) + 1 :])
    return list(map(lambda line: [int(x) for x in line.split(" ")], data))


# Map of the opcodes. It works as follows:
# - r is the register list
# - i is the instruction
# - rr is the register list after the instruction
# - returns the new register list after applying the instruction (block inside the lambda)
instructions = {
    "addr": lambda r, i: (lambda rr: rr.__setitem__(i[3], r[i[1]] + r[i[2]]) or rr)(
        deepcopy(r)
    ),
    "addi": lambda r, i: (lambda rr: rr.__setitem__(i[3], r[i[1]] + i[2]) or rr)(
        deepcopy(r)
    ),
    "mulr": lambda r, i: (lambda rr: rr.__setitem__(i[3], r[i[1]] * r[i[2]]) or rr)(
        deepcopy(r)
    ),
    "muli": lambda r, i: (lambda rr: rr.__setitem__(i[3], r[i[1]] * i[2]) or rr)(
        deepcopy(r)
    ),
    "banr": lambda r, i: (lambda rr: rr.__setitem__(i[3], r[i[1]] & r[i[2]]) or rr)(
        deepcopy(r)
    ),
    "bani": lambda r, i: (lambda rr: rr.__setitem__(i[3], r[i[1]] & i[2]) or rr)(
        deepcopy(r)
    ),
    "borr": lambda r, i: (lambda rr: rr.__setitem__(i[3], r[i[1]] | r[i[2]]) or rr)(
        deepcopy(r)
    ),
    "bori": lambda r, i: (lambda rr: rr.__setitem__(i[3], r[i[1]] | i[2]) or rr)(
        deepcopy(r)
    ),
    "setr": lambda r, i: (lambda rr: rr.__setitem__(i[3], r[i[1]]) or rr)(deepcopy(r)),
    "seti": lambda r, i: (lambda rr: rr.__setitem__(i[3], i[1]) or rr)(deepcopy(r)),
    "gtir": lambda r, i: (
        lambda rr: rr.__setitem__(i[3], 1 if i[1] > r[i[2]] else 0) or rr
    )(deepcopy(r)),
    "gtri": lambda r, i: (
        lambda rr: rr.__setitem__(i[3], 1 if r[i[1]] > i[2] else 0) or rr
    )(deepcopy(r)),
    "gtrr": lambda r, i: (
        lambda rr: rr.__setitem__(i[3], 1 if r[i[1]] > r[i[2]] else 0) or rr
    )(deepcopy(r)),
    "eqir": lambda r, i: (
        lambda rr: rr.__setitem__(i[3], 1 if i[1] == r[i[2]] else 0) or rr
    )(deepcopy(r)),
    "eqri": lambda r, i: (
        lambda rr: rr.__setitem__(i[3], 1 if r[i[1]] == i[2] else 0) or rr
    )(deepcopy(r)),
    "eqrr": lambda r, i: (
        lambda rr: rr.__setitem__(i[3], 1 if r[i[1]] == r[i[2]] else 0) or rr
    )(deepcopy(r)),
}

opcode_possibilities = {}
opcode_mapping = {}


def run_sample(sample):
    global opcode_possibilities

    matching = [
        name
        for name, instr_func in instructions.items()
        if instr_func(sample.before, sample.instruction) == sample.after
    ]

    # Add the opcode to the possibilities with intersection_update so we can narrow down
    opcode = sample.instruction[0]
    if opcode in opcode_possibilities:
        opcode_possibilities[opcode].intersection_update(matching)
    else:
        opcode_possibilities[opcode] = set(matching)

    return matching


def resolve_opcodes():
    global opcode_possibilities, opcode_mapping
    opcode_mapping = {}

    working = {k: set(v) for k, v in opcode_possibilities.items()}

    # Work out which opcode maps to which instruction by finding the only one left, and removing it from the others
    # Repeat until all opcodes are resolved
    while len(opcode_mapping) < len(working):
        for opcode, options in working.items():
            if len(options) == 1:
                instr = next(iter(options))
                opcode_mapping[opcode] = instr

                for other_opcode in working:
                    if other_opcode != opcode:
                        working[other_opcode].discard(instr)

    return opcode_mapping


def run_test_program(program_lines):
    global opcode_mapping

    registers = [0, 0, 0, 0]

    for opcode, A, B, C in program_lines:
        func_name = opcode_mapping[opcode]
        registers = instructions[func_name](registers, [opcode, A, B, C])

    return registers


def part1(inp):
    # Return the number of samples that have at least 3 possible opcodes
    samples = parse_samples(inp)
    return sum(1 for sample in samples if len(run_sample(sample)) >= 3)


def part2(inp):
    # Run the test program and return the value in register 0
    resolve_opcodes()

    test_program = parse_test_program(inp)
    return run_test_program(test_program)[0]


print(f"Part 1: {part1(INPUT)}")
print(f"Part 2: {part2(INPUT)}")
