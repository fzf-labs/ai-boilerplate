import {
  isExtensionStatusRequest,
  type ExtensionStatusResponse,
} from './extension-messages'

chrome.runtime.onInstalled.addListener(async ({ reason }) => {
  if (reason === chrome.runtime.OnInstalledReason.INSTALL) {
    await chrome.storage.local.set({
      installedAt: new Date().toISOString(),
    })
  }
})

chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  if (!isExtensionStatusRequest(message)) {
    return false
  }

  const response: ExtensionStatusResponse = {
    ok: true,
    tabId: sender.tab?.id,
  }

  sendResponse(response)

  return false
})
