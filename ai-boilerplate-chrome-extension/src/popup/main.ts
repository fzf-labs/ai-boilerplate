import './style.css'

import { isInspectableTabUrl } from './tabs'

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
  const [result] = await chrome.scripting.executeScript({
    target: { tabId },
    func: (): PageInfoResponse => ({
      title: document.title || 'Untitled page',
      url: location.href,
      readyState: document.readyState,
    }),
  })

  if (!result?.result) {
    throw new Error('Page inspection did not return data')
  }

  return result.result
}

const refreshTabInfo = async () => {
  setText(contentStatus, 'Checking')

  const tab = await getActiveTab()
  const tabId = tab.id

  setText(tabTitle, tab.title || 'Untitled tab')

  if (!tabId || !isInspectableTabUrl(tab.url)) {
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
    setText(pageSummary, 'This page cannot be inspected by the extension.')
    setText(contentStatus, 'Page access unavailable')
  }
}

refreshButton?.addEventListener('click', () => {
  void refreshTabInfo()
})

void refreshTabInfo()
