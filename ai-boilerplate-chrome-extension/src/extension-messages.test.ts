import assert from 'node:assert/strict'
import test from 'node:test'

import { isExtensionStatusRequest } from './extension-messages.ts'

test('recognizes extension status requests', () => {
  assert.equal(isExtensionStatusRequest({ type: 'GET_EXTENSION_STATUS' }), true)
})

test('rejects unknown or malformed extension messages', () => {
  assert.equal(isExtensionStatusRequest({ type: 'OTHER_MESSAGE' }), false)
  assert.equal(isExtensionStatusRequest({}), false)
  assert.equal(isExtensionStatusRequest(null), false)
})
