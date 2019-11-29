/**
 * @param {number} numRows
 * @param {string} s
 */
function convert(s, numRows) {
  const result = Array(numRows).fill('');
  let currentIndex = 0;
  let forward = true;
  if (numRows < 2) return s;
  [...s].forEach(char => {
    if (currentIndex === numRows - 1) {
      forward = false
    } else if (currentIndex === 0) {
      forward = true
    }
    result[currentIndex]+=char
    forward ? currentIndex++ : currentIndex--
  })
  return result.join("")
}