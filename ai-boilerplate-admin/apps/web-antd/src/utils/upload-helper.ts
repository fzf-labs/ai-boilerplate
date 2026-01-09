/**
 * 上传工具函数
 */

/**
 * 格式化文件大小
 * @param bytes 字节数
 * @returns 格式化后的大小字符串
 */
export function formatFileSize(bytes: number): string {
  if (bytes === 0) return '0 Bytes';

  const k = 1024;
  const sizes = ['Bytes', 'KB', 'MB', 'GB', 'TB'];
  const i = Math.floor(Math.log(bytes) / Math.log(k));

  return `${Number.parseFloat((bytes / k ** i).toFixed(2))} ${sizes[i]}`;
}

/**
 * 获取文件扩展名
 * @param filename 文件名
 * @returns 扩展名（包含点）
 */
export function getFileExtension(filename: string): string {
  const lastDotIndex = filename.lastIndexOf('.');
  // 确保扩展名不是文件名本身（处理以点开头的文件）
  return lastDotIndex > 0 && lastDotIndex < filename.length - 1
    ? filename.slice(lastDotIndex)
    : '';
}

/**
 * 生成安全的文件名
 * @param originalName 原始文件名
 * @returns 安全的文件名
 */
export function generateSafeFileName(originalName: string): string {
  const timestamp = Date.now();
  const ext = getFileExtension(originalName);
  const lastDotIndex = originalName.lastIndexOf('.');
  const nameWithoutExt =
    lastDotIndex > 0 ? originalName.slice(0, lastDotIndex) : originalName;

  // 清理文件名中的特殊字符
  const safeName = nameWithoutExt.replaceAll(/[^a-z0-9\u4E00-\u9FA5]/g, '_');

  return `${safeName}_${timestamp}${ext}`;
}

/**
 * 验证文件类型
 * @param file 文件对象
 * @param allowedTypes 允许的文件类型数组
 * @returns 是否允许的文件类型
 */
export function validateFileType(file: File, allowedTypes: string[]): boolean {
  if (!allowedTypes || allowedTypes.length === 0) return true;

  const fileExtension = getFileExtension(file.name).toLowerCase();

  return allowedTypes.some((type) => {
    // 支持 .ext 格式
    if (type.startsWith('.')) {
      return fileExtension === type.toLowerCase();
    }
    // 支持 MIME 类型
    if (type.includes('/')) {
      return file.type === type;
    }
    // 支持扩展名（不带点）
    return fileExtension === `.${type.toLowerCase()}`;
  });
}

/**
 * 验证文件大小
 * @param file 文件对象
 * @param maxSizeMB 最大大小（MB）
 * @returns 是否符合大小限制
 */
export function validateFileSize(file: File, maxSizeMB: number): boolean {
  const maxSizeBytes = maxSizeMB * 1024 * 1024;
  return file.size <= maxSizeBytes;
}

/**
 * 生成上传路径
 * @param prefix 路径前缀
 * @param fileName 文件名
 * @returns 完整的上传路径
 */
export function generateUploadPath(
  prefix: string = 'upload',
  fileName: string,
): string {
  const now = new Date();
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, '0');
  const day = String(now.getDate()).padStart(2, '0');

  return `${prefix}/${year}/${month}/${day}/${fileName}`;
}
