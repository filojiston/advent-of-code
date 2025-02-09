from collections import deque


def solve(player_count=9, last_marble=25):
    circle = deque([0])
    scores = [0] * player_count

    for marble in range(1, last_marble + 1):
        if marble % 23 == 0:
            circle.rotate(7)
            scores[marble % player_count] += marble + circle.pop()
            circle.rotate(-1)
        else:
            circle.rotate(-1)
            circle.append(marble)

    return max(scores)


print(f"Part 1: {solve(465, 71498)}")
print(f"Part 2: {solve(465, 71498 * 100)}")
