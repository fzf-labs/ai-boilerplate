declare module 'wot-design-uni/components/wd-form/types' {
  export type FormItemRule = {
    required?: boolean
    message?: string
    trigger?: string | string[]
    validator?: (...args: unknown[]) => boolean | Promise<boolean>
  }

  export type FormRules = Record<string, FormItemRule[]>

  export type FormInstance = {
    validate: () => Promise<{ valid: boolean }>
    reset?: () => void
  }
}
