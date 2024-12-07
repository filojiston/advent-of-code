from util.util import read_input
from collections import deque
import heapq

requirements = read_input("day07/input.txt")

def parse_requirements(requirements):
    steps = {}

    for requirement in requirements:
        prerequisite, dependent = requirement.split()[1], requirement.split()[7]

        if prerequisite not in steps:
            steps[prerequisite] = set()
        if dependent not in steps:
            steps[dependent] = set()

        steps[dependent].add(prerequisite)

    return steps

steps = parse_requirements(requirements)

def part1(steps):
    available_steps = [step for step in steps if not steps[step]]
    heapq.heapify(available_steps)

    result = []

    while available_steps:
        next_step = heapq.heappop(available_steps)
        result.append(next_step)

        for dependent, dependencies in steps.items():
            if next_step in dependencies:
                dependencies.remove(next_step)
                if not dependencies:
                    heapq.heappush(available_steps, dependent)

    return ''.join(result)

def part2(steps, num_workers=5, base_time=60):
    available_steps = [step for step in steps if not steps[step]]
    heapq.heapify(available_steps)

    worker_queues = [deque() for _ in range(num_workers)]

    total_time = 0

    while available_steps or any(worker_queues):
        for worker_queue in worker_queues:
            if not worker_queue and available_steps:
                step = heapq.heappop(available_steps)
                completion_time = base_time + ord(step) - ord('A') + 1
                worker_queue.append((step, total_time + completion_time))

        if any(worker_queues):
            next_completion_time = min(queue[0][1] for queue in worker_queues if queue)
        else:
            break

        total_time = next_completion_time

        for worker_queue in worker_queues:
            while worker_queue and worker_queue[0][1] <= next_completion_time:
                step, _ = worker_queue.popleft()
                for dependent, dependencies in steps.items():
                    if step in dependencies:
                        dependencies.remove(step)
                        if not dependencies:
                            heapq.heappush(available_steps, dependent)
    return total_time

print("Part 1: ", part1(steps))
print("Part 2: ", part2(steps))
