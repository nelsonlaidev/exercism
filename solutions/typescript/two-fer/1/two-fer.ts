const names = ['Alice', 'Bohdan', 'Zaphod']

export function twoFer(person?: string): string {
  let name = 'you'

  if (person && names.includes(person)) {
    name = person
  }

  return `One for ${name}, one for me.`
}
