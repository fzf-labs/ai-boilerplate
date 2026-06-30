declare module 'wot-design-uni' {
  export type ToastOptions = Record<string, unknown>
  export type MessageOptions = Record<string, unknown>

  export type ToastInstance = {
    success: (message?: string, options?: ToastOptions) => void
    error: (message?: string, options?: ToastOptions) => void
    warning: (message?: string, options?: ToastOptions) => void
    info: (message?: string, options?: ToastOptions) => void
    loading: (message?: string, options?: ToastOptions) => void
    close: () => void
  }

  export type MessageInstance = {
    confirm: (options?: MessageOptions) => Promise<void>
  }

  export function useToast(): ToastInstance
  export function useMessage(): MessageInstance
}
