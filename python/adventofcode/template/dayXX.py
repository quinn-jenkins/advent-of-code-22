day = "day01"

def part_one(filename: str) -> int:
    with open(filename, 'r') as file:
        return len(file.readlines)

def part_two(filename: str) -> int:
    with open(filename, 'r') as file:
        return len(file.readlines)


if __name__ == '__main__':
    test_file = f'python/adventofcode/{day}/test.txt'
    input_file = f'python/adventofcode/{day}/input.txt'

    test = part_one(test_file)
    assert test == 123
    p1 = part_one(input_file)
    print(f"Part One: {p1}")

    test = part_two(test_file)
    assert test == 123
    p2 = part_two(input_file)
    print(f"Part Two: {p2}")
    