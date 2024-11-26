package day11

import println
import readInput

class Monkey(var monkeyId: Int, var currentItems : MutableList<Long>, var operation: String, var test: Int, var ifTrue: Int, var ifFalse: Int, val reduceWorry: Boolean = true) {
    private var itemsInspected = 0L

    fun getItemsInspected() : Long
    {
        return itemsInspected
    }

    fun catchItem(itemNum: Long)
    {
        currentItems.add(itemNum)
    }

    fun throwItem(lcm : Int = 1) : Pair<Int, Long>? {
        if (currentItems.isEmpty()) {
            return null
        }

        itemsInspected++

        val startingWorry = currentItems.removeFirst()
        val operationSplit = operation.split(' ')
        val firstParam = operationSplit[0].trim()
        val symbol = operationSplit[1].trim()
        val secondParam = operationSplit[2].trim()
        var worry = 0L
        println("Monkey $monkeyId inspects item with worry level of $startingWorry")

        val firstVal : Long = if (firstParam == "old") {
            startingWorry
        } else {
            firstParam.toLong()
        }

        val secondVal : Long = if (secondParam == "old") {
            startingWorry
        } else {
            secondParam.toLong()
        }

        when (symbol) {
            "+" -> {
                worry = firstVal + secondVal
            }
            "*" -> {
                worry = firstVal * secondVal
            }
            else -> {
                println("Something wrong -- symbol isn't + or *")
            }
        }

        println("\tWorry level grows to $worry ($firstVal $symbol $secondVal)")
        if (reduceWorry) {
            worry /= 3
            println("\tWorry divided by 3 is: $worry")
        }

        worry %= lcm
        println("\tReducing by LCM results in: $worry")

        if (worry % test == 0L) {
            println("\t$worry is divisible by $test, throwing to $ifTrue")
            return Pair(ifTrue, worry)
        } else {
            println("\t$worry is not divisible by $test, throwing to $ifFalse")
            return Pair(ifFalse, worry)
        }
    }
}

fun main() {
    fun separateMonkies(input: List<String>): List<List<String>> {
        var startingIndex = 0
        var monkeyBlocks = mutableListOf<List<String>>()
        for ((currentIndex, line) in input.withIndex()) {
            if (line.isEmpty()) {
                monkeyBlocks.add(input.subList(startingIndex, currentIndex))
                startingIndex = currentIndex+1
            }
        }
        monkeyBlocks.add(input.subList(startingIndex, input.size))
        return monkeyBlocks
    }

    fun createMonkey(monkeyNumber: Int, monkeyText: List<String>, reduceWorry: Boolean = false) : Monkey {
        var startingItemsText = monkeyText[1]
            .substringAfter("Starting items: ")
            .split(",")
            .map { it.trim().toLong() }
        var operationText = monkeyText[2].substringAfter("new = ")
        var testValue = monkeyText[3].substringAfter("divisible by ").toInt()
        var ifTrueValue = monkeyText[4].substringAfter("monkey ").toInt()
        var ifFalseValue = monkeyText[5].substringAfter("monkey ").toInt()

        return Monkey(monkeyNumber, startingItemsText.toMutableList(), operationText, testValue, ifTrueValue, ifFalseValue, reduceWorry)
    }

    fun printMonkies(input: List<Monkey>) {
        for (monkey in input) {
            println("Monkey: ${monkey.monkeyId}: ${monkey.currentItems}")
        }
    }

    fun part1(input: List<String>): Long {
        var monkeyBlocks = separateMonkies(input)
        var monkeyList : MutableList<Monkey> = mutableListOf()
        for ((currentIndex, monkeyBlock) in monkeyBlocks.withIndex()) {
            monkeyList.add(createMonkey(currentIndex, monkeyBlock))
        }

        val numRounds = 20
        var currentRound = 1
        while(currentRound <= numRounds) {
            println("Starting round $currentRound of $numRounds")
            val iterator = monkeyList.listIterator()
            while (iterator.hasNext()) {
                var monkey = iterator.next()
                val numItemsHeld = monkey.currentItems.size;
                for (itemNumber in 0..<numItemsHeld)
                {
                    val throwPair = monkey.throwItem()
                    val receivingMonkey = throwPair?.first
                    val itemValue = throwPair?.second

                    if (receivingMonkey != null && itemValue != null) {
                        monkeyList.find { m -> m.monkeyId == receivingMonkey }?.catchItem(itemValue)
                        println("\tMonkey ${monkey.monkeyId} throws item $itemValue to Monkey $receivingMonkey")
                    }
                }
            }

            println("End of round $currentRound")
            printMonkies(monkeyList)
            currentRound++
        }

        var mostInspected = 0L
        var secondMostInspected = 0L
        for (monkey in monkeyList) {
            println("Monkey ${monkey.monkeyId} inspected ${monkey.getItemsInspected()} items")
            if (monkey.getItemsInspected() >= mostInspected) {
                secondMostInspected = mostInspected
                mostInspected = monkey.getItemsInspected()
            } else if (monkey.getItemsInspected() > secondMostInspected) {
                secondMostInspected = monkey.getItemsInspected()
            }
        }

        val monkeyBusiness = mostInspected * secondMostInspected
        println("Most inspected items: $mostInspected")
        println("Second most inspected: $secondMostInspected")
        println("Monkey Business: $monkeyBusiness")

        return monkeyBusiness
    }

    fun part2(input: List<String>): Long {
        var monkeyBlocks = separateMonkies(input)
        var monkeyList : MutableList<Monkey> = mutableListOf()
        for ((currentIndex, monkeyBlock) in monkeyBlocks.withIndex()) {
            monkeyList.add(createMonkey(currentIndex, monkeyBlock, false))
        }

        var lcm = 1
        for (monkey in monkeyList) {
            lcm *= monkey.test
        }

        val numRounds = 10000
        var currentRound = 1
        while(currentRound <= numRounds) {
            println("Starting round $currentRound of $numRounds")
            val iterator = monkeyList.listIterator()
            while (iterator.hasNext()) {
                var monkey = iterator.next()
                val numItemsHeld = monkey.currentItems.size;
                for (itemNumber in 0..<numItemsHeld)
                {
                    val throwPair = monkey.throwItem(lcm)
                    val receivingMonkey = throwPair?.first
                    val itemValue = throwPair?.second

                    if (receivingMonkey != null && itemValue != null) {
                        monkeyList.find { m -> m.monkeyId == receivingMonkey }?.catchItem(itemValue)
                        println("\tMonkey ${monkey.monkeyId} throws item $itemValue to Monkey $receivingMonkey")
                    }
                }
            }

            println("End of round $currentRound")
            printMonkies(monkeyList)
            currentRound++
        }

        var mostInspected = 0L
        var secondMostInspected = 0L
        for (monkey in monkeyList) {
            println("Monkey ${monkey.monkeyId} inspected ${monkey.getItemsInspected()} items")
            if (monkey.getItemsInspected() >= mostInspected) {
                secondMostInspected = mostInspected
                mostInspected = monkey.getItemsInspected()
            } else if (monkey.getItemsInspected() > secondMostInspected) {
                secondMostInspected = monkey.getItemsInspected()
            }
        }

        val monkeyBusiness = mostInspected * secondMostInspected
//        println("Most inspected items: $mostInspected")
//        println("Second most inspected: $secondMostInspected")
        println("Monkey Business: $monkeyBusiness")
        println("LCM is : $lcm")

        return monkeyBusiness
    }

    val testInput = readInput("day11/test")
    // update with the correct answer from the sample input in the problem description
    check(part2(testInput) == 2713310158L)

    val input = readInput("day11/input")
    part1(input).println()
    part2(input).println()
}