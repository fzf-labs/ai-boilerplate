/**
 * 默认图片类型
 */
export const defaultImageAccepts = ['jpg', 'jpeg', 'png', 'gif', 'webp'];

export function checkFileType(file: File, accepts: string[]) {
  if (!accepts || accepts.length === 0) {
    return true;
  }

  const fileExtension = file.name
    .slice(Math.max(0, file.name.lastIndexOf('.')))
    .toLowerCase();

  return accepts.some((type) => {
    // 支持 .ext 格式
    if (type.startsWith('.')) {
      return fileExtension === type.toLowerCase();
    }
    // 支持扩展名（不带点）
    return fileExtension === `.${type.toLowerCase()}`;
  });
}

export function checkImgType(
  file: File,
  accepts: string[] = defaultImageAccepts,
) {
  return checkFileType(file, accepts);
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
  const safeName = nameWithoutExt.replaceAll(/[^a-z0-9\u4E00-\u9FA5]/gi, '_');

  return `${safeName}_${timestamp}${ext}`;
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

  return `${prefix}/${year}${month}${day}/${fileName}`;
}
