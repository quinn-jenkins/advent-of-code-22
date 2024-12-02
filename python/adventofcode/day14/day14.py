day = "day14"

SAND_START_COORD = (500, 0)

# expects format x,y as a string, and returns (x, y) as a tuple of ints
def get_coord(token: str) -> tuple:
    splits = token.split(",")
    return (int(splits[0]), int(splits[1]))

# creates a dictionary representing all coordinates that are currently filled in the cave
# The key would be the height or row, and the value is the column
def create_cave(filename: str) -> dict:
    with open(filename, 'r') as file:
        cave = {}
        for line in file.readlines():
            cave_surface = line.split(" -> ")
            start_point = get_coord(cave_surface[0])
            for line in cave_surface[1:]:
                end_point = get_coord(line)
                if start_point[0] != end_point[0]:
                    row = start_point[1]
                    wall_lower = min(start_point[0], end_point[0])
                    wall_upper = max(start_point[0], end_point[0])
                    for col in range(wall_lower, wall_upper+1):
                        if row not in cave:
                            cave[row] = set()
                        cave[row].add(col)
                elif start_point[1] != end_point[1]:
                    col = start_point[0]
                    wall_lower = min(start_point[1], end_point[1])
                    wall_upper = max(start_point[1], end_point[1])
                    for row in range(wall_lower, wall_upper+1):
                        if row not in cave:
                            cave[row] = set()
                        cave[row].add(col)
                start_point = end_point

    return cave

# checks if the sand either hits a solid defined in our solids dictionary, or for part 2 if it hits the floor
def hit_something(solids: dict, coor: tuple, floor: int=None) -> bool:
    if floor is not None:
        if coor[1] == floor:
            return True
    return (coor[1] in solids and coor[0] in solids[coor[1]])

# tries to move the coordinate down, down-left, and down-right in order. If the sand is able to move in any of those directions,
# that coordinate is returned. However, if all 3 directions are blocked, then we return the existing coordinate to indicate that 
# the sand didn't have anywhere to go
def move_sand_down(solids: dict, current_coord: tuple, floor: int=None) -> tuple:
    down_coor = (current_coord[0], current_coord[1]+1)
    if hit_something(solids, down_coor, floor):
        # there's a solid below, so we can't do that
        left_coor = (down_coor[0] - 1, down_coor[1])
        if hit_something(solids, left_coor, floor):
            # there's a solid diagonal left, so we can't do that
            right_coor = (down_coor[0] + 1, down_coor[1])
            if hit_something(solids, right_coor, floor):
                # there's a solid diagonal left, so we're out of options
                # return the current_coor to signal that we didn't move
                return current_coord
            else:
                return right_coor
        else:
            return left_coor
    else:
        return down_coor

# for part one, anything that falls lower than the lowest level is going into the abyss, 
# so our end condition is when the first piece of sand reaches that lowest level
def get_sand_resting_coord(solids: dict, lowest_level: int) -> tuple:
    current_sand_position = SAND_START_COORD
    last_sand_position = current_sand_position
    current_sand_position = move_sand_down(solids, current_sand_position)
    while (last_sand_position != current_sand_position) and current_sand_position[1] <= lowest_level:
        last_sand_position = current_sand_position
        current_sand_position = move_sand_down(solids, current_sand_position)
    
    if current_sand_position[1] > lowest_level:
        return None

    return current_sand_position

# for part two, instead of a lowest level, we have a floor that sand can continue piling up on, 
# so our end condition is when the pile of sand reaches the start point
def get_sand_resting_coord_pt2(solids: dict, floor: int) -> tuple:
    current_sand_position = SAND_START_COORD
    last_sand_position = current_sand_position
    current_sand_position = move_sand_down(solids, current_sand_position, floor)
    while (last_sand_position != current_sand_position):
        last_sand_position = current_sand_position
        current_sand_position = move_sand_down(solids, current_sand_position, floor)
    
    if current_sand_position == SAND_START_COORD:
        return None

    return current_sand_position

def part_one(filename: str) -> int:
    cave = create_cave(filename)
    tallest_rock = max(cave)
    # print(f"Tallest rock is {tallest_rock}")
    count = 0
    resting_coord = get_sand_resting_coord(cave, tallest_rock)
    while resting_coord is not None:
        count += 1
        if resting_coord[1] not in cave:
            cave[resting_coord[1]] = set()
        cave[resting_coord[1]].add(resting_coord[0])
        # print(f"Sand unit {count} comes to rest at {resting_coord[0]}, {resting_coord[1]}")
        resting_coord = get_sand_resting_coord(cave, tallest_rock)

    return count

def part_two(filename: str) -> int:
    cave = create_cave(filename)
    tallest_rock = max(cave)
    floor = tallest_rock + 2
    # print(f"Floor is {floor}")
    count = 1 # starting the count at 1, because the example also counts the last piece of sand where it doesn't leave the origin
    resting_coord = get_sand_resting_coord_pt2(cave, floor)
    while resting_coord is not None:
        count += 1
        if resting_coord[1] not in cave:
            cave[resting_coord[1]] = set()
        cave[resting_coord[1]].add(resting_coord[0])
        # print(f"Sand unit {count} comes to rest at {resting_coord[0]}, {resting_coord[1]}")
        resting_coord = get_sand_resting_coord_pt2(cave, floor)

    return count


if __name__ == '__main__':
    test_file = f'python/adventofcode/{day}/test.txt'
    input_file = f'python/adventofcode/{day}/input.txt'

    test = part_one(test_file)
    assert test == 24
    p1 = part_one(input_file)
    print(f"Part One: {p1}")
    assert p1 == 1061

    test = part_two(test_file)
    assert test == 93
    p2 = part_two(input_file)
    print(f"Part Two: {p2}")
    