def get_opp_choice_value(opponentChoice):
    return ord(opponentChoice) - 64

def get_self_choice_value(selfChoice):
    return ord(selfChoice) - 87

def is_win(opponentChoice, selfChoice):
    if opponentChoice == "A" and selfChoice == "Y":
        return True
    if opponentChoice == "B" and selfChoice == "Z":
        return True
    if opponentChoice == "C" and selfChoice == "X":
        return True

    return False

def is_draw(opponentChoice, selfChoice):
    if get_opp_choice_value(opponentChoice) == get_self_choice_value(selfChoice):
        return True
    return False

def part_one():
    totalScore = 0
    with open("input.txt", 'r') as file:
        for line in file:
            line = line.strip()
            choices = line.split(" ")

            if is_win(choices[0], choices[1]):
                totalScore += (get_self_choice_value(choices[1]) + 6)
                print(f"Win -- score is now {totalScore}")
            elif is_draw(choices[0], choices[1]):
                totalScore += (get_opp_choice_value(choices[0]) + 3)
                print(f"Tie -- score is now {totalScore}")
            else:
                totalScore += (ord(choices[1][0]) - 87)
                print(f"Loss -- score is now {totalScore}")

def get_losing_score(opponentChoice):
    if opponentChoice == "A":
        return 3
    elif opponentChoice == "B":
        return 1
    else:
        return 2

def get_winning_score(opponentChoice):
    if opponentChoice == "A":
        return 2
    elif opponentChoice == "B":
        return 3
    else:
        return 1

def part_two():
    totalScore = 0
    with open("input.txt", 'r') as file:
        for line in file:
            line = line.strip()
            choices = line.split(" ")

            if choices[1] == "X":
                # lose
                totalScore += get_losing_score(choices[0])
                print(f"Lose -- score is now {totalScore}")
            elif choices[1] == "Y":
                # tie
                totalScore += get_opp_choice_value(choices[0]) + 3
                print(f"Tie -- score is now {totalScore}")
            elif choices[1] == "Z":
                # win
                totalScore += get_winning_score(choices[0]) + 6
                print(f"Win -- score is now {totalScore}")



if __name__ == '__main__':
    # part_one()
    part_two()




