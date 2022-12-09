if __name__ == '__main__':
    calorieList = []
    calories = 0
    elfCount = 0
    with open("input.txt", 'r') as file:
        for line in file:
            if line == '\n':
                calorieList.append(calories)
                calories = 0
            else:
                calories += int(line)
    calorieList.sort()
    calorieList.reverse()

    print(calorieList[0])
    print(calorieList[1])
    print(calorieList[2])
    top3 = calorieList[0] + calorieList[1] + calorieList[2]
    print(f"Sum: {top3}")
