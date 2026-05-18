import type { CustomRequestOptions } from '@/http/types'
import { http } from './http'

/*
 * openapi-ts-request 工具的 request 跨客户端适配方法
 */
export default function request<T>(
  url: string,
  options: Omit<CustomRequestOptions, 'url'> & {
    params?: Record<string, unknown>
    headers?: Record<string, unknown>
  },
) {
  const { params, headers, ...restOptions } = options
  const requestOptions: CustomRequestOptions = {
    url,
    ...restOptions,
  }

  if (params) {
    requestOptions.query = params
  }

  if (headers) {
    requestOptions.header = headers
  }

  return http<T>(requestOptions)
}
