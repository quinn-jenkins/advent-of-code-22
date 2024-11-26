import println
import readInput

fun main() {
    fun part1(input: List<String>): Int {
        return input.size
    }

    fun part2(input: List<String>): Int {
        return input.size
    }

    val testInput = readInput("day01/test")
    // update with the correct answer from the sample input in the problem description
    check(part1(testInput) == 1)

    val input = readInput("day01/input")
    part1(input).println()
    part2(input).println()
}
