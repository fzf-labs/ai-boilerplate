#!/usr/bin/env node
/**
 * Stop Hook
 * AI 完成回答后触发
 */

const { execSync } = require('child_process');

function getCodeChanges() {
  try {
    const status = execSync('git status --porcelain', { encoding: 'utf8' });
    if (!status.trim()) {
      return { hasChanges: false, files: [] };
    }

    const files = status.split('\n')
      .filter(Boolean)
      .map(line => ({
        status: line.substring(0, 2).trim(),
        file: line.substring(3)
      }));

    return { hasChanges: true, files };
  } catch {
    return { hasChanges: false, files: [] };
  }
}

function main() {
  const changes = getCodeChanges();

  // 构建通知消息
  const notificationMessage = changes.hasChanges
    ? `修改了 ${changes.files.length} 个文件`
    : '任务完成,无文件变更';

  // 发送系统通知
  try {
    // 构建通知标题
    const title = `Claude Code Complete`;

    // 转义特殊字符，避免命令注入
    const escapedTitle = title.replace(/"/g, '\\"').replace(/`/g, '\\`').replace(/\$/g, '\\$');
    const escapedMessage = notificationMessage.replace(/"/g, '\\"').replace(/`/g, '\\`').replace(/\$/g, '\\$');
    const cmd = `terminal-notifier -title "${escapedTitle}" -message "${escapedMessage}" -sound default`;
    console.log('🔔 发送通知:', title, '-', notificationMessage);
    execSync(cmd, { encoding: 'utf8' });
  } catch (error) {
    console.error('❌ 通知发送失败:', error.message);
  }

  process.exit(0);
}

main();
