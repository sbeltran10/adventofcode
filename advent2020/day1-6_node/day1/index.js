// https://adventofcode.com/2020/day/1

module.exports = (scriptArgs, rl) => {
  return new Promise((resolve) => {
    let multiplyResult = 0
    let numbers = []

    rl.on('line', (line) => {
      const currentNumber = Number(line)
      numbers.forEach((storedNumber1, index1) => {

        if (scriptArgs[0] === 'part-2') {
          numbers.forEach((storedNumber2, index2) =>{

            if (index1 !== index2 && (currentNumber + storedNumber1 + storedNumber2 === 2020)) {
              multiplyResult = currentNumber * storedNumber1 * storedNumber2
              rl.close()
            }

          })
        } else {
          if (currentNumber + storedNumber1 === 2020) {
            multiplyResult = currentNumber * storedNumber1
            rl.close()
          }
        }
      })
      numbers.push(currentNumber)
    })

    rl.on('close', () => {
      resolve(multiplyResult)
    })
  })
}