export const isInspectableTabUrl = (url: string | undefined): boolean => {
  if (!url) {
    return false
  }

  try {
    const { protocol } = new URL(url)

    return protocol === 'http:' || protocol === 'https:'
  } catch {
    return false
  }
}
