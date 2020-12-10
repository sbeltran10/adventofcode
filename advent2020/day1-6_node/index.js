// using node v12.13.0
// run with `node index.js day<day_number> part-<part_number>`

// Main executor
const fs = require('fs')
const readline = require('readline')

const scripts = {
  day1: require('./day1'),
  day2: require('./day2'),
  day3: require('./day3'),
  day4: require('./day4'),
  day5: require('./day5'),
  day6: require('./day6')
}

const args = process.argv.slice(2)

const scriptName = args[0] // Script to execute
const scriptArgs = args.slice(1) // optional args

const exec = async () => {
  const input = `./${scriptName}/input.txt`
  const stream = fs.createReadStream((input))
  const rl = readline.createInterface({
    input: stream
  })

  const result = await scripts[scriptName](scriptArgs, rl, input)

  console.log("========Result========")
  console.log(result)
}


exec()