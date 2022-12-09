def get_opp_choice_value(opponent_choice):
    return ord(opponent_choice) - 64


def get_self_choice_value(self_choice):
    return ord(self_choice) - 87


def is_win(opponent_choice, self_choice):
    if opponent_choice == "A" and self_choice == "Y":
        return True
    if opponent_choice == "B" and self_choice == "Z":
        return True
    if opponent_choice == "C" and self_choice == "X":
        return True

    return False


def is_draw(opponent_choice, self_choice):
    if get_opp_choice_value(opponent_choice) == get_self_choice_value(self_choice):
        return True
    return False


def part_one():
    total_score = 0
    with open("input.txt", 'r') as file:
        for line in file:
            line = line.strip()
            choices = line.split(" ")

            if is_win(choices[0], choices[1]):
                total_score += (get_self_choice_value(choices[1]) + 6)
                print(f"Win -- score is now {total_score}")
            elif is_draw(choices[0], choices[1]):
                total_score += (get_opp_choice_value(choices[0]) + 3)
                print(f"Tie -- score is now {total_score}")
            else:
                total_score += (ord(choices[1][0]) - 87)
                print(f"Loss -- score is now {total_score}")


def get_losing_score(opponent_choice):
    if opponent_choice == "A":
        return 3
    elif opponent_choice == "B":
        return 1
    else:
        return 2


def get_winning_score(opponent_choice):
    if opponent_choice == "A":
        return 2
    elif opponent_choice == "B":
        return 3
    else:
        return 1


def part_two():
    total_score = 0
    with open("input.txt", 'r') as file:
        for line in file:
            line = line.strip()
            choices = line.split(" ")

            if choices[1] == "X":
                # lose
                total_score += get_losing_score(choices[0])
                print(f"Lose -- score is now {total_score}")
            elif choices[1] == "Y":
                # tie
                total_score += get_opp_choice_value(choices[0]) + 3
                print(f"Tie -- score is now {total_score}")
            elif choices[1] == "Z":
                # win
                total_score += get_winning_score(choices[0]) + 6
                print(f"Win -- score is now {total_score}")


if __name__ == '__main__':
    # part_one()
    part_two()
