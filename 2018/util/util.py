def read_input(file_path):
    with open(file_path, "r") as file:
        return [line.strip() for line in file.readlines()]


def read_input_without_strip(file_path):
    with open(file_path, "r") as file:
        return [line.rstrip("\n") for line in file]
