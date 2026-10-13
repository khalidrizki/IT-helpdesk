export function levenshteinDistance(a: string, b: string): number {
  const matrix: number[][] = []

  for (let i = 0; i <= b.length; i++) {
    matrix[i] = [i]
  }

  for (let j = 0; j <= a.length; j++) {
    matrix[0]![j] = j
  }

  for (let i = 1; i <= b.length; i++) {
    for (let j = 1; j <= a.length; j++) {
      if (b[i - 1] === a[j - 1]) {
        matrix[i]![j] = matrix[i - 1]![j - 1]!
      } else {
        matrix[i]![j] = Math.min(
          matrix[i - 1]![j]! + 1,
          matrix[i]![j - 1]! + 1,
          matrix[i - 1]![j - 1]! + 1,
        )
      }
    }
  }

  return matrix[b.length]![a.length]!
}

export function matchesSearch(title: string, query: string): boolean {
  const words = title.toLowerCase().split(/\s+/)
  const search = query.toLowerCase().trim()

  if (search === '') return true

  // Tahap 1: Cari kata yang berawalan sesuai query
  const hasPrefixMatch = words.some((word) => word.startsWith(search))

  if (hasPrefixMatch) return true

  // Tahap 2: Fallback ke pencarian fuzzy
  return words.some((word) => levenshteinDistance(word, search) <= 2)
}
