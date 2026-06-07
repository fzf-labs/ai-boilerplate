const ENCODED_COMPONENT_PATTERN = /%[0-9a-f]{2}/i
const MAX_DECODE_DEPTH = 5

export interface ParsedUrl {
  path: string
  query: Record<string, string>
}

export function ensureDecodeURIComponent(value: string) {
  let decoded = value

  for (let depth = 0; depth < MAX_DECODE_DEPTH && ENCODED_COMPONENT_PATTERN.test(decoded); depth += 1) {
    try {
      const next = decodeURIComponent(decoded)
      if (next === decoded)
        return decoded
      decoded = next
    }
    catch {
      return decoded
    }
  }

  return decoded
}

/**
 * 解析 url 得到 path 和 query
 * 比如输入 url: /pages/login/login?redirect=%2Fpages%2Forders%2Fdetail%3Fid%3D1
 * 输出: {path: /pages/login/login, query: {redirect: /pages/orders/detail?id=1}}
 */
export function parseUrlToObj(url: string): ParsedUrl {
  const queryStartIndex = url.indexOf('?')
  const path = queryStartIndex === -1 ? url : url.slice(0, queryStartIndex)
  const queryStr = queryStartIndex === -1 ? '' : url.slice(queryStartIndex + 1)

  if (!queryStr) {
    return {
      path,
      query: {},
    }
  }

  const query: Record<string, string> = {}
  queryStr.split('&').forEach((item) => {
    if (!item)
      return

    const separatorIndex = item.indexOf('=')
    const rawKey = separatorIndex === -1 ? item : item.slice(0, separatorIndex)
    if (!rawKey)
      return

    const rawValue = separatorIndex === -1 ? '' : item.slice(separatorIndex + 1)
    query[ensureDecodeURIComponent(rawKey)] = ensureDecodeURIComponent(rawValue)
  })

  return { path, query }
}
