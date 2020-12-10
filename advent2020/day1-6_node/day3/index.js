// https://adventofcode.com/2020/day/1
// Input movement pairs separated with coma, ex 1,5 3,4 2,3

const readline = require('readline')
const fs = require('fs')

module.exports = async (scriptArgs, rl, input) => {
  let totalMult = 1

  for (const movementPair of scriptArgs) {
    const rl = readline.createInterface({
      input: fs.createReadStream((input))
    })    
    totalMult *= await countPathTrees(movementPair, rl)
  }

  return totalMult
}

const countPathTrees = (movementPair, rl) => {
  return new Promise((resolve) => {
    let totalTrees = 0
    const splitMovement = movementPair.split(',')
    const rightMovement = Number(splitMovement[0])
    const downMovement = Number(splitMovement[1])

    let currentStepRight = 0
    let currentStepDown = 0

    rl.on('line', (line) => {
      if (currentStepDown && (currentStepDown % downMovement === 0)) {
        currentStepRight += rightMovement
        const positionChar = line.charAt(currentStepRight % line.length)
        if (positionChar === '#') totalTrees++
      }
      currentStepDown++
    })

    rl.on('close', () => {
      resolve(totalTrees)
    })
  })
}