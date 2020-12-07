// https://adventofcode.com/2020/day/5S

module.exports = (scriptArgs, rl) => {
  return new Promise((resolve) => {

    let highestId = 0
    let seatIds = []
    let candidates = []

    rl.on('line', (line) => {
      const { id } = calculateSeat(line)
      if (id > highestId) highestId = id
      if (scriptArgs[0] === 'part-2') {
        seatIds.forEach(seatId => {
          if (Math.abs(seatId - id) === 2) {   
            let candidate         
            if (seatId > id) candidate = id + 1
            else candidate = seatId + 1
            candidates.push(candidate)
          }
        })
        seatIds.push(id)
      }
    })

    rl.on('close', () => {
      if (scriptArgs[0] === 'part-2') {
        resolve(candidates.filter(candidate => !seatIds.includes(candidate))[0])
      } else{
        resolve(highestId)
      }      
    })
  })
}

const calculateSeat = (seat) => {
  let row
  let column

  const rowCode = seat.substring(0, 6)
  let currentRowLower = 0
  let currentRowUpper = 127
  for (const char of rowCode) {
    const middleNumber = (currentRowUpper + currentRowLower + 1) / 2
    if (char === 'F') currentRowUpper = middleNumber - 1
    else if (char === 'B') currentRowLower = middleNumber
  }

  row = seat.charAt(6) === 'F' ? currentRowLower : currentRowUpper

  const colCode = seat.substring(7, 9)
  let currentColLower = 0
  let currentColUpper = 7
  for (const char of colCode) {
    const middleNumber = (currentColUpper + currentColLower + 1) / 2
    if (char === 'L') currentColUpper = middleNumber - 1
    else if (char === 'R') currentColLower = middleNumber
  }

  column = seat.charAt(9) === 'L' ? currentColLower : currentColUpper

  return {
    row,
    column,
    id: row * 8 + column
  }
}