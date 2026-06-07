import assert from 'node:assert/strict'
import test from 'node:test'

import { isInspectableTabUrl } from './tabs.ts'

test('allows only http and https tab URLs for inspection', () => {
  assert.equal(isInspectableTabUrl('http://example.test'), true)
  assert.equal(isInspectableTabUrl('https://example.test/path'), true)
})

test('rejects extension, browser, file, malformed, and prefix-only URLs', () => {
  assert.equal(isInspectableTabUrl('chrome://extensions'), false)
  assert.equal(isInspectableTabUrl('chrome-extension://abc/popup.html'), false)
  assert.equal(isInspectableTabUrl('file:///tmp/page.html'), false)
  assert.equal(isInspectableTabUrl('httpx://example.test'), false)
  assert.equal(isInspectableTabUrl(undefined), false)
})
