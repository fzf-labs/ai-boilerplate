import {
  uploadFileUrl,
  useUpload as useFilePickerUpload,
} from '@/utils/uploadFile'

type TfileType = 'image' | 'file'
type TImage = 'png' | 'jpg' | 'jpeg' | 'webp' | '*'
type TFile = 'doc' | 'docx' | 'ppt' | 'zip' | 'xls' | 'xlsx' | 'txt' | TImage

interface TOptions<T extends TfileType> {
  formData?: Record<string, any>
  maxSize?: number
  accept?: T extends 'image' ? TImage[] : TFile[]
  fileType?: T
  success?: (params: any) => void
  error?: (err: any) => void
}

export default function useUpload<T extends TfileType>(options: TOptions<T> = {} as TOptions<T>) {
  const {
    formData = {},
    maxSize = 5 * 1024 * 1024,
    fileType = 'image',
    success,
    error,
  } = options

  const upload = useFilePickerUpload(
    uploadFileUrl.DEFAULT,
    formData,
    {
      maxSize: maxSize / 1024 / 1024,
      onSuccess: success,
      onError: error,
    },
  )

  if (fileType !== 'image') {
    console.warn('useUpload currently delegates to uni.chooseImage; non-image file selection is not supported by this wrapper.')
  }

  return upload
}
