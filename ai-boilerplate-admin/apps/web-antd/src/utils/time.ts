/**
 * 将秒数转换为可读的时间格式
 * @param seconds 秒数
 * @returns 格式化后的时间字符串，如："1小时30分钟15秒"
 */
export function formatSecondsToTime(seconds: number): string {
  if (!seconds || seconds < 0) {
    return '0秒';
  }

  const hours = Math.floor(seconds / 3600);
  const minutes = Math.floor((seconds % 3600) / 60);
  const remainingSeconds = seconds % 60;

  const parts: string[] = [];

  if (hours > 0) {
    parts.push(`${hours}小时`);
  }

  if (minutes > 0) {
    parts.push(`${minutes}分钟`);
  }

  if (remainingSeconds > 0 || parts.length === 0) {
    parts.push(`${remainingSeconds}秒`);
  }

  return parts.join('');
}

/**
 * 将毫秒转换为可读的时间格式
 * @param milliseconds 毫秒数
 * @returns 格式化后的时间字符串
 */
export function formatMillisecondsToTime(milliseconds: number): string {
  return formatSecondsToTime(Math.floor(milliseconds / 1000));
}

// 返回今日的00:00:00和23:59:59
export function getTodayTimeRange(): [string, string] {
  const today = new Date().toISOString().split('T')[0]; // YYYY-MM-DD
  return [`${today} 00:00:00`, `${today} 23:59:59`];
}
