export function decodedValue(colors: string[]) {
  return getCode(colors[0]) * 10 + getCode(colors[1])
}

function getCode(color: string): number {
  return COLORS.findIndex((c) => c === color)
}

const COLORS = ['black', 'brown', 'red', 'orange', 'yellow', 'green', 'blue', 'violet', 'grey', 'white']
