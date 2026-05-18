chrome.runtime.onInstalled.addListener(async ({ reason }) => {
  if (reason === chrome.runtime.OnInstalledReason.INSTALL) {
    await chrome.storage.local.set({
      installedAt: new Date().toISOString(),
    })
  }
})

chrome.runtime.onMessage.addListener((message, sender, sendResponse) => {
  if (message?.type !== 'GET_EXTENSION_STATUS') {
    return false
  }

  sendResponse({
    ok: true,
    tabId: sender.tab?.id,
  })

  return false
})
