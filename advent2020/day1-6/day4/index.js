// https://adventofcode.com/2020/day/4

module.exports = (scriptArgs, rl) => {
  return new Promise((resolve) => {
    let validPassports = 0

    let currentPassportRuleCount = 0
    rl.on('line', (line) => {
      // Check which fields are in this line
      rules.forEach(rule => {
        if (line.includes(rule)) {

          if (scriptArgs[0] === 'part-2') {
            const fieldBeingIndex = line.indexOf(`${rule}:`) + 4       
            const valueEndIndex = line.replace(/\n/g, ' ').indexOf(' ', fieldBeingIndex)
            const value = line.substring(fieldBeingIndex, valueEndIndex !== -1 ? valueEndIndex : line.length)
            
            if (validateField(rule, value)) currentPassportRuleCount++

          } else {
            currentPassportRuleCount++
          }
        }
      })

      // On blank line Evaluate if passport is valid, then reset passport
      if (line === '') {
        if (currentPassportRuleCount === requiredRules) validPassports++
        currentPassportRuleCount = 0
      }
    })

    rl.on('close', () => {
      // Check last count and then return result
      if (currentPassportRuleCount === requiredRules) validPassports++

      resolve(validPassports)
    })
  })
}

// Passport rules
const rules = ['hgt', 'byr', 'iyr', 'eyr', 'hcl', 'ecl', 'pid']
const requiredRules = rules.length

const validateField = (rule, value) => {
  switch (rule) {
    case 'hgt':
      const units = value.substring(value.length - 2)
      const number = Number(value.substring(0, value.indexOf(units)))
      if (units === 'cm' && number >= 150 && number <= 193) return true
      if (units === 'in' && number >= 59 && number <= 76) return true
      return false

    case 'byr':
      if (isNaN(value)) return false
      if (value.length !== 4) return false
      if (Number(value) < 1920 || Number(value) > 2002) return false
      return true

    case 'iyr':
      if (isNaN(value)) return false
      if (value.length !== 4) return false
      if (Number(value) < 2010 || Number(value) > 2020) return false
      return true

    case 'eyr':
      if (isNaN(value)) return false
      if (value.length !== 4) return false
      if (Number(value) < 2020 || Number(value) > 2030) return false
      return true

    case 'hcl':
      return  /^#([a-fA-F0-9]{6}|[a-fA-F0-9]{3})$/.test(value)

    case 'ecl':
      return ['amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'].includes(value)

    case 'pid':
      return /^\d{9}$/.test(value)
    default:
      return false
  }
}
