import './style.css'

type PageInfoResponse = {
  title: string
  url: string
  readyState: DocumentReadyState
}

const pageSummary = document.querySelector<HTMLParagraphElement>('#page-summary')
const tabTitle = document.querySelector<HTMLElement>('#tab-title')
const contentStatus = document.querySelector<HTMLElement>('#content-status')
const refreshButton = document.querySelector<HTMLButtonElement>('#refresh-button')

const setText = (element: Element | null, value: string) => {
  if (element) {
    element.textContent = value
  }
}

const getActiveTab = async () => {
  const [tab] = await chrome.tabs.query({
    active: true,
    currentWindow: true,
  })

  return tab
}

const getPageInfo = async (tabId: number): Promise<PageInfoResponse> => {
  return chrome.tabs.sendMessage(tabId, {
    type: 'GET_PAGE_INFO',
  })
}

const refreshTabInfo = async () => {
  setText(contentStatus, 'Checking')

  const tab = await getActiveTab()
  const tabId = tab.id

  setText(tabTitle, tab.title || 'Untitled tab')

  if (!tabId || !tab.url?.startsWith('http')) {
    setText(pageSummary, 'Open an http or https page to use the content script.')
    setText(contentStatus, 'Not available on this page')
    return
  }

  try {
    const pageInfo = await getPageInfo(tabId)
    setText(pageSummary, pageInfo.url)
    setText(tabTitle, pageInfo.title)
    setText(contentStatus, `Ready (${pageInfo.readyState})`)
  } catch {
    setText(pageSummary, 'Refresh the target page, then open the popup again.')
    setText(contentStatus, 'Content script not connected')
  }
}

refreshButton?.addEventListener('click', () => {
  void refreshTabInfo()
})

void refreshTabInfo()
