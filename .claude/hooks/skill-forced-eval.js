#!/usr/bin/env node
/**
 * Skill Forced Evaluation Hook
 * 用户提交问题时触发,评估并激活相关技能
 *
 * 特性:
 * - 支持关键词权重匹配
 * - 技能分组互斥(同组只激活最高分)
 * - 最大激活数量限制
 */

// 配置
const CONFIG = {
  maxSkills: 3,           // 最多激活技能数
  minScore: 1,            // 最低激活分数
  stdinTimeout: 100,      // stdin 读取超时(ms)
};

// 技能定义
// priority: 优先级(越高越优先), group: 互斥分组, keywords: [关键词, 权重]
const skills = [
  // 后端开发技能 (group: backend-*)
  {
    name: 'backend-audit',
    priority: 10,
    group: 'backend-init',
    keywords: [
      ['后端开发', 2], ['开发流程', 2], ['审计', 3], ['工件检查', 3],
      ['开发起点', 3], ['后端任务', 1], ['从哪开始', 2]
    ]
  },
  {
    name: 'backend-database',
    priority: 9,
    group: 'backend-db',
    keywords: [
      ['数据库', 2], ['SQL', 2], ['建表', 3], ['PostgreSQL', 2],
      ['表结构', 3], ['字段', 1], ['索引', 2], ['DDL', 3],
      ['schema', 2], ['数据库设计', 3], ['migration', 2]
    ]
  },
  {
    name: 'backend-gorm',
    priority: 8,
    group: 'backend-orm',
    keywords: [
      ['GORM', 3], ['Model', 1], ['DAO', 3], ['Repo', 2],
      ['ORM', 2], ['数据访问层', 3], ['gorm生成', 3]
    ]
  },
  {
    name: 'backend-proto-gen',
    priority: 8,
    group: 'backend-proto',
    keywords: [
      ['Proto', 2], ['Protobuf', 3], ['sqltopb', 3], ['gRPC', 2],
      ['HTTP接口', 1], ['API定义', 2], ['proto生成', 3]
    ]
  },
  {
    name: 'backend-proto-edit',
    priority: 7,
    group: 'backend-proto',
    keywords: [
      ['Proto编辑', 3], ['过滤条件', 2], ['验证规则', 2],
      ['RPC方法', 2], ['buf.validate', 3], ['proto修改', 3]
    ]
  },
  {
    name: 'backend-api-gen',
    priority: 7,
    group: 'backend-api',
    keywords: [
      ['API生成', 3], ['pbtocode', 3], ['pb.go', 3],
      ['http.pb.go', 3], ['Service桩', 2], ['make api', 3]
    ]
  },
  {
    name: 'backend-codeing',
    priority: 6,
    group: 'backend-code',
    keywords: [
      ['CRUD', 2], ['增删改查', 2], ['后端代码', 2], ['Service', 1],
      ['业务模块', 2], ['API接口', 1], ['后端功能', 2], ['业务逻辑', 2]
    ]
  },
  {
    name: 'backend-quality',
    priority: 5,
    group: 'backend-quality',
    keywords: [
      ['代码质量', 2], ['wire', 2], ['gci', 2], ['lint', 2],
      ['依赖注入', 2], ['格式化', 1], ['质量检查', 3]
    ]
  },
  {
    name: 'backend-api-schema-test',
    priority: 5,
    group: 'backend-test',
    keywords: [
      ['API测试', 3], ['Swagger', 2], ['接口测试', 3], ['契约测试', 3],
      ['Schema验证', 3], ['schemathesis', 3], ['回归测试', 2], ['OpenAPI', 2]
    ]
  },

  // 前端开发技能
  {
    name: 'admin-codeing',
    priority: 6,
    group: 'frontend-admin',
    keywords: [
      ['管理后台', 3], ['前端页面', 2], ['表单', 1], ['表格', 1],
      ['Vue', 1], ['Ant Design', 2], ['组件', 1], ['路由', 1],
      ['CRUD页面', 3], ['权限控制', 2], ['admin', 2]
    ]
  },
  {
    name: 'admin-api-gen',
    priority: 7,
    group: 'frontend-admin-api',
    keywords: [
      ['Admin API', 3], ['管理后台API', 3], ['后台接口同步', 3]
    ]
  },
  {
    name: 'app-codeing',
    priority: 6,
    group: 'frontend-app',
    keywords: [
      ['移动端', 3], ['App', 2], ['UniApp', 3], ['小程序', 3],
      ['H5', 2], ['移动应用', 3], ['手机端', 2], ['uni-app', 3],
      ['wot-design', 3], ['z-paging', 3], ['移动开发', 3], ['跨端', 2]
    ]
  },
  {
    name: 'app-api-gen',
    priority: 7,
    group: 'frontend-app-api',
    keywords: [
      ['App API', 3], ['TypeScript API', 2], ['API客户端', 2],
      ['openapi-ts-request', 3], ['API同步', 2], ['API类型', 2]
    ]
  },

  // 通用技能
  {
    name: 'ui-ux-pro-max',
    priority: 4,
    group: 'design',
    keywords: [
      ['UI设计', 3], ['UX设计', 3], ['界面', 1], ['样式', 1],
      ['布局', 1], ['组件设计', 2], ['响应式', 2], ['动画', 2],
      ['交互', 1], ['Tailwind', 2], ['配色', 2], ['字体', 1]
    ]
  },
  {
    name: 'interview',
    priority: 3,
    group: 'planning',
    keywords: [
      ['头脑风暴', 3], ['想法', 1], ['设计方案', 2], ['创意', 2],
      ['探索', 1], ['需求分析', 3], ['方案讨论', 3], ['功能设计', 2],
      ['澄清', 2], ['确认', 1]
    ]
  },
  {
    name: 'tech-decision',
    priority: 4,
    group: 'planning',
    keywords: [
      ['技术选型', 3], ['技术对比', 3], ['架构决策', 3], ['选择框架', 3],
      ['选择库', 2], ['技术方案', 2], ['推荐', 1], ['评估', 1], ['技术评审', 3]
    ]
  },
  {
    name: 'skill-creator',
    priority: 5,
    group: 'meta',
    keywords: [
      ['创建技能', 3], ['Skill', 2], ['技能开发', 3],
      ['扩展能力', 2], ['自定义技能', 3], ['技能编写', 3]
    ]
  },
  {
    name: 'git-workflow',
    priority: 4,
    group: 'vcs',
    keywords: [
      ['Git', 2], ['版本控制', 2], ['分支', 1], ['提交', 1],
      ['commit', 2], ['PR', 2], ['merge', 2], ['rebase', 2],
      ['代码管理', 2], ['pull request', 3]
    ]
  }
];

/**
 * 计算技能匹配分数
 */
function calculateScore(skill, prompt) {
  const lowerPrompt = prompt.toLowerCase();
  let score = 0;

  for (const [keyword, weight] of skill.keywords) {
    if (lowerPrompt.includes(keyword.toLowerCase())) {
      score += weight;
    }
  }

  return score;
}

/**
 * 选择最佳技能(考虑分组互斥)
 */
function selectBestSkills(scoredSkills) {
  // 按分数降序,同分按优先级降序
  scoredSkills.sort((a, b) => {
    if (b.score !== a.score) return b.score - a.score;
    return b.skill.priority - a.skill.priority;
  });

  const selected = [];
  const usedGroups = new Set();

  for (const item of scoredSkills) {
    if (selected.length >= CONFIG.maxSkills) break;

    // 同组只取最高分的
    if (usedGroups.has(item.skill.group)) continue;

    selected.push(item);
    usedGroups.add(item.skill.group);
  }

  return selected;
}

/**
 * 读取 stdin
 */
async function readStdin() {
  if (process.stdin.isTTY) return '';

  return new Promise((resolve) => {
    let data = '';
    const timeout = setTimeout(() => resolve(data), CONFIG.stdinTimeout);

    process.stdin.setEncoding('utf8');
    process.stdin.on('data', (chunk) => {
      clearTimeout(timeout);
      data += chunk;
    });
    process.stdin.on('end', () => {
      clearTimeout(timeout);
      resolve(data);
    });
    process.stdin.on('error', () => {
      clearTimeout(timeout);
      resolve(data);
    });
    process.stdin.resume();
  });
}

/**
 * 解析输入获取 prompt
 */
function parsePrompt(stdinData, argv) {
  let prompt = '';

  if (stdinData.trim()) {
    try {
      const input = JSON.parse(stdinData);
      prompt = input.prompt || input.user_prompt || '';
    } catch {
      prompt = stdinData;
    }
  }

  if (!prompt && argv.length > 2) {
    prompt = argv.slice(2).join(' ');
  }

  return prompt.trim();
}

/**
 * 主函数
 */
async function main() {
  try {
    const stdinData = await readStdin();
    const prompt = parsePrompt(stdinData, process.argv);

    // 跳过空输入或斜杠命令
    if (!prompt || /^\/[^\/\s]+/.test(prompt)) {
      process.exit(0);
    }

    // 计算所有技能的匹配分数
    const scoredSkills = skills
      .map(skill => ({ skill, score: calculateScore(skill, prompt) }))
      .filter(item => item.score >= CONFIG.minScore);

    if (scoredSkills.length === 0) {
      process.exit(0);
    }

    // 选择最佳技能
    const selected = selectBestSkills(scoredSkills);

    // 输出结果
    const names = selected.map(s => s.skill.name).join(', ');
    const files = selected.map(s => `- .claude/skills/${s.skill.name}/SKILL.md`).join('\n');

    console.log(`[技能激活] 检测到 ${selected.length} 个相关技能：${names}
请读取以下技能文件获取规范：
${files}`);

  } catch (err) {
    // 静默失败,不影响用户体验
  }
  process.exit(0);
}

main();
