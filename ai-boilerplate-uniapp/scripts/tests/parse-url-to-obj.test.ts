import assert from 'node:assert/strict'

import { ensureDecodeURIComponent, parseUrlToObj } from '../../src/utils/url.ts'

const nested = parseUrlToObj(
  '/pages-fg/login/login?redirect=%2Fpages%2Forders%2Fdetail%3ForderId%3Dabc%3D123&token=abc%3Ddef',
)

assert.equal(nested.path, '/pages-fg/login/login')
assert.deepEqual(nested.query, {
  redirect: '/pages/orders/detail?orderId=abc=123',
  token: 'abc=def',
})

const plainEncoded = parseUrlToObj('/pages/profile/edit?name=Foo%20Bar&flag')

assert.deepEqual(plainEncoded.query, {
  name: 'Foo Bar',
  flag: '',
})

assert.equal(ensureDecodeURIComponent('%252Fpages%252Findex%252Findex'), '/pages/index/index')
assert.equal(ensureDecodeURIComponent('%E0%A4%A'), '%E0%A4%A')
