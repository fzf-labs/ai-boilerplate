export const extensionStatusRequestType = 'GET_EXTENSION_STATUS'

export type ExtensionStatusRequest = {
  type: typeof extensionStatusRequestType
}

export type ExtensionStatusResponse = {
  ok: true
  tabId?: number
}

export const isExtensionStatusRequest = (message: unknown): message is ExtensionStatusRequest => {
  return (
    typeof message === 'object'
    && message !== null
    && 'type' in message
    && message.type === extensionStatusRequestType
  )
}
