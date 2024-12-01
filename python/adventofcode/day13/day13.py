from ast import literal_eval
from functools import cmp_to_key

day = "day13"

#both are ints
def compare_ints(left: int, right: int) -> int:
    return left - right

# both are lists
def compare_lists(left: list, right: list) -> int:
    len_left, len_right = len(left), len(right)
    min_len = min(len_left, len_right)
    for i in range(min_len):
        comparison = compare(left[i], right[i])
        if comparison == 0:
            # values are the same -- keep going
            continue
        else:
            return comparison
    
    # if all the values up to the minimum length matched, check that they are the same length
    return compare_ints(len_left, len_right)

def compare(left, right) -> int:
    if isinstance(left, int) and isinstance(right, int):
        return compare_ints(left, right)
    elif isinstance(left, int):
        return compare_lists([left], right)
    elif isinstance(right, int):
        return compare_lists(left, [right])
    else:
        return compare_lists(left, right)

def part_one(filename: str) -> int:
    with open(filename, 'r') as file:
        data = []
        for line in file.read().split("\n\n"):
            for split in line.split():
                # the data is formatted the same as a python list, so we can use literal_eval to read it back into the correct data type
                data.append(literal_eval(split))

        sum = 0
        pair_index = 1
        for left, right in zip(*[iter(data)] * 2):
            if compare(left, right) < 0:
                sum += pair_index
            pair_index += 1
        return sum     

def part_two(filename: str) -> int:
    with open(filename, 'r') as file:
        divider_1, divider_2 = [[2]], [[6]]
        data = [divider_1, divider_2]
        for line in file.read().split("\n\n"):
            for split in line.split():
                data.append(literal_eval(split))

        cmp_key = cmp_to_key(compare)
        sorted_data = sorted(data, key=cmp_key)
        
        div_1_inx = sorted_data.index(divider_1) + 1 # add 1 because the packets are not zero indexed in the problem description
        div_2_inx = sorted_data.index(divider_2) + 1

        return div_1_inx * div_2_inx
        

if __name__ == '__main__':
    test_file = f'python/adventofcode/{day}/test.txt'
    input_file = f'python/adventofcode/{day}/input.txt'

    test = part_one(test_file)
    assert test == 13
    p1 = part_one(input_file)
    print(f"Part One: {p1}")

    test = part_two(test_file)
    assert test == 140
    p2 = part_two(input_file)
    print(f"Part Two: {p2}")
    