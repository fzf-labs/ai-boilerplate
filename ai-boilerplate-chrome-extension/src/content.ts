type PageInfoResponse = {
  title: string
  url: string
  readyState: DocumentReadyState
}

chrome.runtime.onMessage.addListener((message, _sender, sendResponse) => {
  if (message?.type !== 'GET_PAGE_INFO') {
    return false
  }

  const response: PageInfoResponse = {
    title: document.title || 'Untitled page',
    url: location.href,
    readyState: document.readyState,
  }

  sendResponse(response)

  return false
})

console.info('[ai-boilerplate-chrome-extension] content script ready')
