// https://adventofcode.com/2020/day/6

module.exports = (scriptArgs, rl) => {
  return new Promise((resolve) => {
    const part2 = scriptArgs[0] === 'part-2'
    let totalCount = 0

    let currentGroupSet = new Set()
    let currentGroupSetCount = {}
    let amountOfMembers = 0
    rl.on('line', (line) => {
      if (line === '') {
        if (part2) {
          for (const answerCount of Object.values(currentGroupSetCount)) {
            if (answerCount === amountOfMembers) totalCount++
          }
          currentGroupSetCount = {}
          amountOfMembers = 0
        } else {
          // On blank line restart group set and count answers
          totalCount += currentGroupSet.size
          currentGroupSet = new Set()
        }
      } else {
        amountOfMembers++
        for (const char of line) {
          if (part2) {
            if (!currentGroupSetCount[char]) currentGroupSetCount[char] = 1
            else currentGroupSetCount[char]++
          } else {
            currentGroupSet.add(char)
          }
        }
      }
    })

    rl.on('close', () => {
      // Add sum of last group
      if (part2) {
        for (const answerCount of Object.values(currentGroupSetCount)) {
          if (answerCount === amountOfMembers) totalCount++
        }
      } else {
        totalCount += currentGroupSet.size
      }

      resolve(totalCount)
    })
  })
}