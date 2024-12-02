package day12

import println
import readInput
import java.util.LinkedList
import java.util.Queue

class Node(var location: Pair<Int, Int>, var parent: Node?) {
    fun getValidNeighborLocations(maxX: Int, maxY: Int): MutableSet<Pair<Int, Int>> {
        val neighbors = listOf(
            Pair(location.first, location.second - 1), Pair(location.first, location.second + 1),
            Pair(location.first - 1, location.second), Pair(location.first + 1, location.second)
        )

        val validNeighbors = mutableSetOf<Pair<Int, Int>>()
        for (neighbor in neighbors) {
            if (neighbor.first >= 0 && neighbor.second >= 0 && neighbor.first < maxX && neighbor.second < maxY) {
                validNeighbors.add(neighbor)
            }
        }
        return validNeighbors
    }

    fun isLocationAccessible(input: List<String>, otherLocation: Pair<Int, Int>, partTwo: Boolean = false): Boolean {
        if (partTwo) {
            return isLocationAccessiblePt2(input, otherLocation)
        }
        return isLocationAccessiblePt1(input, otherLocation)
    }

    private fun isLocationAccessiblePt1(input: List<String>, otherLocation: Pair<Int, Int>): Boolean {
        val thisNodeVal = input[location.first][location.second]
        val otherLocationVal = input[otherLocation.first][otherLocation.second]
        return (otherLocationVal.code - thisNodeVal.code) <= 1
    }

    private fun isLocationAccessiblePt2(input: List<String>, otherLocation: Pair<Int, Int>): Boolean {
        val thisNodeVal = input[location.first][location.second]
        val otherNodeVal = input[otherLocation.first][otherLocation.second]
        return (thisNodeVal.code - otherNodeVal.code) <= 1
    }

    override fun toString(): String {
        return "${location.first},${location.second}"
    }
}

fun main() {
    fun findCharacter(input: List<String>, endCharacter: Char): Pair<Int, Int>? {
        for (i in input.indices) {
            for (j in input[i].indices) {
                if (input[i][j] == endCharacter) {
                    return Pair(i, j)
                }
            }
        }

        return null
    }

    fun findAllPt2EndPoints(input: List<String>, endChar: Char): List<Pair<Int, Int>> {
        val potentialEndPoints = mutableListOf<Pair<Int, Int>>()
        for (i in input.indices) {
            for (j in input[i].indices) {
                if (input[i][j] == endChar) {
                    potentialEndPoints.add(Pair(i, j))
                }
            }
        }
        return potentialEndPoints
    }

    fun getPathFromEndNode(endNode: Node): List<Pair<Int, Int>>? {
        val path = mutableListOf<Pair<Int, Int>>()
        var currentNode = endNode
        while (currentNode.parent != null) {
            path.add(currentNode.location)
            currentNode = currentNode.parent!!
        }
        return path.reversed()
    }

    fun pathSearch(
        input: List<String>, endPoints: List<Pair<Int, Int>>, queue: Queue<Node>, partTwo: Boolean = false
    ): List<Pair<Int, Int>>? {
        val maxX = input.size
        val maxY = input[0].length
        val visitedLocations = mutableSetOf<Pair<Int, Int>>()
        while (queue.isNotEmpty()) {
            val currentNode = queue.poll()
            if (visitedLocations.contains(currentNode.location)) {
                continue
            }
            visitedLocations.add(currentNode.location)
//            println("Visiting ${currentNode.location} with height ${input[currentNode.location.first][currentNode.location.second]}")
            if (endPoints.contains(currentNode.location)) {
                return getPathFromEndNode(currentNode)
            }

            val neighborLocations = currentNode.getValidNeighborLocations(maxX, maxY)
            for (neighbor in neighborLocations) {
                if (!visitedLocations.contains(neighbor) && currentNode.isLocationAccessible(input, neighbor, partTwo)) {
                    queue.add(Node(neighbor, currentNode))
                }
            }
        }

        println("Failed to find the end node ($endPoints)")
        return null
    }

    fun part1(input: MutableList<String>): Int? {
        var endPoint = findCharacter(input, 'E')
        input[0] = input[0].replace('S', 'a')
        if (endPoint != null) {
            input[endPoint.first] = input[endPoint.first].replace('E', 'z')
            val queue: Queue<Node> = LinkedList()
            queue.add(Node(Pair(0, 0), null))
            val path = pathSearch(input, listOf(endPoint), queue)
            if (path != null) {
                for (i in 1..<path.size) {
                    val last = path[i - 1]
                    val current = path[i]
                    if (last.first != current.first && last.second != current.second) {
                        println("ERROR! Diagonal move: $last to $current")
                    }
                    if (input[current.first][current.second] - input[last.first][last.second] > 1) {
                        println("ERROR! Moved more than 1 up: $last to $current")
                    }
                }
            }
            println("Num steps: ${path?.size} to end location $endPoint")
            return path?.size
        }

        return null
    }

    fun part2(input: MutableList<String>): Int? {
        val startPoint = findCharacter(input, 'E')
        input[0] = input[0].replace('S', 'a')
        if (startPoint != null) {
            input[startPoint.first] = input[startPoint.first].replace('E', 'z')
            val queue: Queue<Node> = LinkedList()
            queue.add(Node(Pair(startPoint.first, startPoint.second), null))
            val potentialEndPoints = findAllPt2EndPoints(input, 'a')
            println("Potential end points: $potentialEndPoints")
            val path = pathSearch(input, potentialEndPoints, queue, true)
            if (path != null) {
                println(path)
                for (i in 1..<path.size) {
                    val last = path[i - 1]
                    val current = path[i]
                    if (last.first != current.first && last.second != current.second) {
                        println("ERROR! Diagonal move: $last to $current")
                    }
                    if (input[current.first][current.second] - input[last.first][last.second] > 1) {
                        println("ERROR! Moved more than 1 up: $last to $current")
                    }
                }
            }
            println("Num steps: ${path?.size}")
            return path?.size
        }

        return null
    }

    val testInput = readInput("day12/test")
    // update with the correct answer from the sample input in the problem description
    check(part1(testInput.toMutableList()) == 31)

    check(part2(testInput.toMutableList()) == 29)

    val input = readInput("day12/input")
    part1(input.toMutableList()).println()
    part2(input.toMutableList()).println()
}
