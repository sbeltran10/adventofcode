// https://adventofcode.com/2020/day/1

module.exports = (scriptArgs, rl) => {
  return new Promise((resolve) => {
    let amountValid = 0
    rl.on('line', (line) => {
      const splitLine = line.split(' ')
      const numPolicies = splitLine[0].split('-')

      const firstVal = Number(numPolicies[0])
      const secondVal = Number(numPolicies[1])

      const policyChar = splitLine[1].charAt(0)

      if(scriptArgs[0] === 'part-2'){
        const pos1Valid = splitLine[2].charAt(firstVal - 1) === policyChar
        const pos2Valid = splitLine[2].charAt(secondVal - 1) === policyChar
        
        if((pos1Valid && !pos2Valid) || (!pos1Valid && pos2Valid)) amountValid++

      } else{
        let charAmount = 0
        for (const char of splitLine[2]) {
          if(char === policyChar) charAmount++
        }
  
        if(charAmount >= firstVal && charAmount <= secondVal) amountValid++
      }
    })

    rl.on('close', () => {
      resolve(amountValid)
    })
  })
}