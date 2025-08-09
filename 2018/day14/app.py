import unittest

INPUT = 147061


def build_recipes(after=None, pattern=None):
    recipes = [3, 7]
    elf1 = 0
    elf2 = 1

    # Part 2, find how many recipes appear before the pattern
    if pattern is not None:
        pattern = str(pattern)
        pattern_len = len(pattern)
        pattern_ints = [int(c) for c in pattern]

        while True:
            new_sum = recipes[elf1] + recipes[elf2]

            if new_sum >= 10:
                recipes.append(1)
                if (
                    len(recipes) >= pattern_len
                    and recipes[-pattern_len:] == pattern_ints
                ):
                    return len(recipes) - pattern_len

                recipes.append(new_sum - 10)
                if (
                    len(recipes) >= pattern_len
                    and recipes[-pattern_len:] == pattern_ints
                ):
                    return len(recipes) - pattern_len
            else:
                recipes.append(new_sum)
                if (
                    len(recipes) >= pattern_len
                    and recipes[-pattern_len:] == pattern_ints
                ):
                    return len(recipes) - pattern_len

            recipes_len = len(recipes)
            elf1 = (elf1 + recipes[elf1] + 1) % recipes_len
            elf2 = (elf2 + recipes[elf2] + 1) % recipes_len

    # Part 1, find the next 10 recipes after given number of recipes
    else:
        while len(recipes) < after + 10:
            new_sum = recipes[elf1] + recipes[elf2]
            new_digits = list(map(int, str(new_sum)))
            recipes.extend(new_digits)

            elf1 = (elf1 + recipes[elf1] + 1) % len(recipes)
            elf2 = (elf2 + recipes[elf2] + 1) % len(recipes)

        return recipes


def part1(input_number):
    recipes = build_recipes(after=input_number)
    return "".join(map(str, recipes[input_number : input_number + 10]))


def part2(pattern):
    return build_recipes(pattern=pattern)


# Test example cases
class TestDay14(unittest.TestCase):
    def test_part1_input_9(self):
        self.assertEqual("5158916779", part1(9))

    def test_part1_input_5(self):
        self.assertEqual("0124515891", part1(5))

    def test_part1_input_18(self):
        self.assertEqual("9251071085", part1(18))

    def test_part1_input_2018(self):
        self.assertEqual("5941429882", part1(2018))

    def test_part2_input_51589(self):
        self.assertEqual(9, part2("51589"))

    def test_part2_input_01245(self):
        self.assertEqual(5, part2("01245"))

    def test_part2_input_92510(self):
        self.assertEqual(18, part2("92510"))

    def test_part2_input_59414(self):
        self.assertEqual(2018, part2("59414"))


# To run with the test inputs, run those two lines
# if __name__ == "__main__":
#     unittest.main()


# To run with the actual input, run those two lines
print("Part 1: ", part1(INPUT))
print("Part 2: ", part2(INPUT))
