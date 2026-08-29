export function decodedResistorValue(colors: string[]): string {
  const zeros = getCode(colors[2])
  const total = (getCode(colors[0]) * 10 + getCode(colors[1])) * getByZeros(zeros)

  if (total >= 1e9) {
    return (total / 1e9).toFixed(1).replace(/\.0$/, '') + ' gigaohms'
  }

  if (total >= 1e6) {
    return (total / 1e6).toFixed(1).replace(/\.0$/, '') + ' megaohms'
  }

  if (total >= 1e3) {
    return (total / 1e3).toFixed(1).replace(/\.0$/, '') + ' kiloohms'
  }

  return total.toString() + ' ohms'
}

function getCode(color: string): number {
  return COLORS.findIndex((c) => c === color)
}

function getByZeros(zeros: number): number {
  return 10 ** zeros
}

const COLORS = ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white']
