def is_win(opponentChoice, selfChoice):
    if opponentChoice == "A" and selfChoice == "Y":
        return True
    if opponentChoice == "B" and selfChoice == "Z":
        return True
    if opponentChoice == "C" and selfChoice == "X":
        return True

    return False

def is_draw(opponentChoice, selfChoice):
    if (ord(opponentChoice) - 64) == (ord(selfChoice) - 87):
        return True
    return False

if __name__ == '__main__':
    totalScore = 0
    with open("input.txt", 'r') as file:
        for line in file:
            line = line.strip()
            choices = line.split(" ")

            if is_win(choices[0], choices[1]):
                totalScore += ((ord(choices[1][0]) - 87) + 6)
                print(f"Win -- score is now {totalScore}")
            elif is_draw(choices[0], choices[1]):
                totalScore += (ord(choices[0][0]) - 64 + 3)
                print(f"Tie -- score is now {totalScore}")
            else:
                totalScore += (ord(choices[1][0]) - 87)
                print(f"Loss -- score is now {totalScore}")




