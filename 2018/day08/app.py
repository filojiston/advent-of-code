from util.util import read_input

inp = list(map(lambda x: int(x), read_input("day08/input.txt")[0].split(" ")))


class Node:
    def __init__(self, children=None, metadata=None):
        self.children = children if children is not None else []
        self.metadata = metadata if metadata is not None else []

    def get_metadata_sum(self):
        metadata_sum = sum(self.metadata)

        for child in self.children:
            metadata_sum += child.get_metadata_sum()

        return metadata_sum

    def get_value(self):
        if len(self.children) == 0:
            return sum(self.metadata)
        else:
            value = 0
            for index in self.metadata:
                if 1 <= index <= len(self.children):
                    value += self.children[index - 1].get_value()
            return value


def build_tree(data, start_index=0):
    node_children = []
    num_children = data[start_index]
    num_metadata_entries = data[start_index + 1]

    current_index = start_index + 2
    for _ in range(num_children):
        child_node, next_index = build_tree(data, current_index)
        node_children.append(child_node)
        current_index = next_index

    metadata = data[current_index : current_index + num_metadata_entries]
    node = Node(children=node_children, metadata=metadata)

    return node, current_index + num_metadata_entries


def part1(inp):
    root, _ = build_tree(inp)
    return root.get_metadata_sum()


def part2(inp):
    root, _ = build_tree(inp)
    return root.get_value()


print("Part 1: ", part1(inp))
print("Part 2: ", part2(inp))
