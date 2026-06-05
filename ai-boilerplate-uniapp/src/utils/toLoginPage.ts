import { LOGIN_PAGE } from '@/router/config'
import { currRoute, getLastPage } from '@/utils'
import { debounce } from '@/utils/debounce'

interface ToLoginPageOptions {
  /**
   * 跳转模式, uni.navigateTo | uni.reLaunch
   * @default 'navigateTo'
   */
  mode?: 'navigateTo' | 'reLaunch'
  /**
   * 查询参数
   * @example '?redirect=/pages/home/index'
   */
  queryString?: string
}

function buildRedirectQueryString() {
  const { path, query } = currRoute()
  if (!path || path === LOGIN_PAGE) {
    return ''
  }

  const queryString = Object.entries(query)
    .map(([key, value]) => `${key}=${encodeURIComponent(value)}`)
    .join('&')
  const currentUrl = queryString ? `${path}?${queryString}` : path
  return `?redirect=${encodeURIComponent(currentUrl)}`
}

/**
 * 跳转到登录页, 带防抖处理
 *
 * 如果要立即跳转，不做延时，可以使用 `toLoginPage.flush()` 方法
 */
export const toLoginPage = debounce((options: ToLoginPageOptions = {}) => {
  const { mode = 'navigateTo', queryString = '' } = options

  // 获取当前页面路径
  const currentPage = getLastPage()
  const currentPath = currentPage?.route ? `/${currentPage.route}` : ''
  // 如果已经在登录页，则不跳转
  if (currentPath === LOGIN_PAGE) {
    return
  }

  const url = `${LOGIN_PAGE}${queryString || buildRedirectQueryString()}`

  if (mode === 'navigateTo') {
    uni.navigateTo({ url })
  }
  else {
    uni.reLaunch({ url })
  }
}, 500)
