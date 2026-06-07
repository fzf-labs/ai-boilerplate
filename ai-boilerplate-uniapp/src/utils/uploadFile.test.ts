import assert from 'node:assert/strict'

import {
  isFileWithinMaxSize,
  parseUploadResponseData,
} from './uploadFile'

assert.equal(isFileWithinMaxSize(4 * 1024 * 1024, 5), true)
assert.equal(isFileWithinMaxSize(6 * 1024 * 1024, 5), false)
assert.equal(isFileWithinMaxSize(undefined, 5), true)

assert.deepEqual(
  parseUploadResponseData<{ url: string }>(
    JSON.stringify({ code: 0, data: { url: 'https://cdn.example/avatar.png' } }),
  ),
  { url: 'https://cdn.example/avatar.png' },
)

assert.deepEqual(
  parseUploadResponseData<{ url: string }>(
    JSON.stringify({ url: 'https://cdn.example/direct.png' }),
  ),
  { url: 'https://cdn.example/direct.png' },
)

assert.throws(
  () => parseUploadResponseData(JSON.stringify({ code: 500, message: 'upload failed' })),
  /upload failed/,
)

assert.throws(
  () => parseUploadResponseData('{bad json'),
  /上传响应解析失败/,
)
