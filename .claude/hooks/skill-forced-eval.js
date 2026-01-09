#!/usr/bin/env node
/**
 * Skill Forced Evaluation Hook
 * 用户提交问题时触发,评估并激活相关技能
 */

const skills = [
  { name: 'backend-database', keywords: ['数据库', 'SQL', '建表', 'PostgreSQL', 'MySQL', '表结构', '字段', '索引', 'DDL', 'schema', '数据库设计', 'migration'] },
  { name: 'backend-codeing', keywords: ['CRUD', '增删改查', '后端开发', '后端代码', 'Service', 'GORM', 'Protobuf', '代码生成', '业务模块', 'API接口', '后端功能', '业务逻辑', 'pbtocode', 'sqltopb'] },
  { name: 'api-schema-test', keywords: ['API测试', 'Swagger', '接口测试', '契约测试', 'Schema验证', 'schemathesis', '回归测试', 'OpenAPI'] },
  { name: 'admin-codeing', keywords: ['管理后台', '前端页面', '表单', '表格', 'Vue', 'Ant Design', '组件', '路由', 'CRUD页面', '权限控制'] },
  { name: 'app-codeing', keywords: ['移动端', 'App', 'UniApp', '小程序', 'H5', '移动应用', '手机端', 'uni-app', 'wot-design', 'z-paging', '移动开发', '跨端'] },
  { name: 'ui-ux-pro-max', keywords: ['UI设计', 'UX设计', '界面', '样式', '布局', '组件设计', '响应式', '动画', '交互', 'Tailwind', '配色', '字体'] },
  { name: 'product-requirements', keywords: ['产品需求', 'PRD', '需求文档', '功能规格', '用户故事', '验收标准', '需求分析', '产品设计'] },
  { name: 'interview', keywords: ['头脑风暴', '想法', '设计方案', '创意', '探索', '需求分析', '方案讨论', '功能设计', '澄清', '确认'] },
  { name: 'prd-to-testcase', keywords: ['测试用例', '测试计划', 'PRD转测试', '测试文档', '测试场景', '边界测试', 'UI测试'] },
  { name: 'prototype-design', keywords: ['原型', '原型设计', 'HTML原型', 'CSS', '界面设计', '交互设计', '移动端', 'PC端', '高保真原型'] },
  { name: 'tech-decision', keywords: ['技术选型', '技术对比', '架构决策', '选择框架', '选择库', '技术方案', '推荐', '评估', '技术评审'] },
  { name: 'webapp-testing', keywords: ['Web测试', '自动化测试', 'Playwright', '浏览器测试', 'E2E测试', '前端测试', 'UI自动化'] },
  { name: 'skill-creator', keywords: ['创建技能', 'Skill', '技能开发', '扩展能力', '自定义技能', '技能编写'] },
  { name: 'prompt-optimizer', keywords: ['Prompt优化', '提示词工程', '提示词优化', 'AI指令', '提示词框架', 'Prompt框架'] },
  { name: 'git-workflow', keywords: ['Git', '版本控制', '分支', '提交', 'commit', 'PR', 'merge', 'rebase', '代码管理', 'pull request'] },
  { name: 'bug-detective', keywords: ['Bug', '问题', '错误', '排查', '调试', 'debug', '故障', '异常', 'panic', '日志', '报错', '不工作'] }
];

async function readStdin() {
  if (process.stdin.isTTY) return '';

  return new Promise((resolve) => {
    let data = '';
    const timeout = setTimeout(() => resolve(data), 100);

    process.stdin.setEncoding('utf8');
    process.stdin.on('data', (chunk) => { clearTimeout(timeout); data += chunk; });
    process.stdin.on('end', () => { clearTimeout(timeout); resolve(data); });
    process.stdin.on('error', () => { clearTimeout(timeout); resolve(data); });
    process.stdin.resume();
  });
}

async function main() {
  try {
    let prompt = '';
    const stdinData = await readStdin();

    if (stdinData.trim()) {
      try {
        const input = JSON.parse(stdinData);
        prompt = input.prompt || input.user_prompt || '';
      } catch {
        prompt = stdinData;
      }
    }

    if (!prompt && process.argv.length > 2) {
      prompt = process.argv.slice(2).join(' ');
    }

    if (!prompt.trim() || /^\/[^\/\s]+/.test(prompt.trim())) {
      process.exit(0);
    }

    const lowerPrompt = prompt.toLowerCase();
    const matched = skills.filter(s =>
      s.keywords.some(k => lowerPrompt.includes(k.toLowerCase()))
    );

    if (matched.length > 0) {
      console.log(`[技能激活] 检测到 ${matched.length} 个相关技能：${matched.map(s => s.name).join(', ')}
请读取以下技能文件获取规范：
${matched.map(s => `- .claude/skills/${s.name}/SKILL.md`).join('\n')}`);
    }
  } catch {}
  process.exit(0);
}

main();
