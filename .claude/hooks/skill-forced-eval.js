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
  // ============ 后端开发技能 ============
  {
    name: 'backend-dev',
    priority: 11,
    group: 'backend-workflow',
    keywords: [
      // 核心触发词 - 完整后端开发流程
      ['帮我开发后端', 5], ['后端开发流程', 5], ['完整后端', 5],
      ['从零开发后端', 5], ['后端CRUD开发', 5], ['开发后端功能', 5],
      ['后端功能开发', 4], ['帮我做后端', 4], ['帮我写后端', 4],
      // 场景触发
      ['新功能后端', 3], ['后端全流程', 4], ['整个后端', 3],
      // 英文支持
      ['backend development', 4], ['full backend', 4],
      ['develop backend', 4], ['backend workflow', 4]
    ]
  },
  {
    name: 'backend-audit',
    priority: 10,
    group: 'backend-init',
    keywords: [
      // 核心触发词
      ['审计', 3], ['工件检查', 3], ['开发起点', 3], ['从哪开始', 3],
      ['前置条件', 3], ['开发前', 2], ['检查状态', 2],
      // 流程相关
      ['后端开发', 2], ['开发流程', 2], ['后端任务', 1],
      // 英文支持
      ['audit', 3], ['backend dev', 2]
    ]
  },
  {
    name: 'backend-database',
    priority: 9,
    group: 'backend-db',
    keywords: [
      // 核心触发词
      ['建表', 3], ['创建表', 3], ['新建表', 3], ['表设计', 3],
      ['数据库设计', 3], ['表结构', 3], ['修改表', 3],
      // 技术词
      ['PostgreSQL', 2], ['SQL', 2], ['DDL', 3], ['migration', 2],
      ['字段', 1], ['索引', 2], ['schema', 2],
      // 通用
      ['数据库', 2], ['database', 2]
    ]
  },
  {
    name: 'backend-gorm',
    priority: 8,
    group: 'backend-orm',
    keywords: [
      // 核心触发词
      ['GORM', 3], ['gorm生成', 3], ['make gorm', 3],
      ['DAO', 3], ['Repo', 2], ['数据访问层', 3],
      // 技术词
      ['ORM', 2], ['Model', 1], ['gen.go', 3],
      // 英文支持
      ['gorm code', 3]
    ]
  },
  {
    name: 'backend-proto-gen',
    priority: 8,
    group: 'backend-proto',
    keywords: [
      // 核心触发词 - 必须用 sqltopb
      ['sqltopb', 3], ['proto生成', 3], ['make sqltopb', 3],
      ['生成Proto', 3], ['创建Proto', 3],
      // 技术词
      ['Proto', 2], ['Protobuf', 3], ['protobuf', 3],
      ['gRPC', 2], ['grpc', 2], ['API定义', 2],
      // 接口定义
      ['接口定义', 2], ['.proto', 3]
    ]
  },
  {
    name: 'backend-proto-edit',
    priority: 7,
    group: 'backend-proto',
    keywords: [
      // 核心触发词
      ['Proto编辑', 3], ['proto修改', 3], ['修改Proto', 3],
      ['编辑Proto', 3],
      // 具体场景
      ['过滤条件', 2], ['验证规则', 2], ['RPC方法', 2],
      ['buf.validate', 3], ['添加RPC', 3], ['删除RPC', 3],
      ['ListReq', 2]
    ]
  },
  {
    name: 'backend-api-gen',
    priority: 7,
    group: 'backend-api',
    keywords: [
      // 核心触发词
      ['make api', 3], ['pbtocode', 3], ['API代码生成', 3],
      // 生成文件
      ['pb.go', 3], ['http.pb.go', 3], ['grpc.pb.go', 3],
      // 场景
      ['Service桩', 2], ['生成Go代码', 2]
    ]
  },
  {
    name: 'backend-codeing',
    priority: 6,
    group: 'backend-code',
    keywords: [
      // 核心触发词
      ['后端功能', 3], ['业务逻辑', 3], ['后端代码', 2],
      ['CRUD', 2], ['增删改查', 2],
      // 具体场景
      ['Service实现', 3], ['业务模块', 2], ['后端实现', 3],
      // 英文支持
      ['backend feature', 3], ['backend implement', 3],
      ['service logic', 2], ['data layer', 2]
    ]
  },
  {
    name: 'backend-quality',
    priority: 5,
    group: 'backend-quality',
    keywords: [
      // 核心触发词
      ['代码质量', 3], ['质量检查', 3], ['make wire', 3],
      ['make lint', 3], ['make gci', 3],
      // 技术词
      ['wire', 2], ['gci', 2], ['lint', 2], ['依赖注入', 2],
      // 场景
      ['格式化', 1], ['代码检查', 2]
    ]
  },
  {
    name: 'backend-api-schema-test',
    priority: 5,
    group: 'backend-test',
    keywords: [
      // 核心触发词
      ['契约测试', 3], ['Schema测试', 3], ['schemathesis', 3],
      ['API测试', 3], ['接口测试', 3],
      // 技术词
      ['Swagger测试', 3], ['OpenAPI测试', 3],
      ['Schema验证', 3], ['回归测试', 2],
      // 英文支持
      ['schema test', 3], ['api test', 2]
    ]
  },

  // ============ 前端开发技能 ============
  {
    name: 'admin-dev',
    priority: 11,
    group: 'frontend-admin-workflow',
    keywords: [
      // 核心触发词 - 完整管理后台开发流程
      ['帮我开发管理后台', 5], ['开发Admin页面', 5], ['完整前端开发', 5],
      ['前端CRUD开发', 5], ['帮我开发后台', 5], ['开发后台功能', 5],
      ['管理后台开发', 4], ['帮我做后台页面', 4], ['帮我写后台', 4],
      // 场景触发
      ['新功能前端', 3], ['前端全流程', 4], ['整个后台', 3],
      ['后台开发流程', 4], ['管理端开发', 4],
      // 英文支持
      ['admin development', 4], ['full admin', 4],
      ['develop admin', 4], ['admin workflow', 4]
    ]
  },
  {
    name: 'admin-codeing',
    priority: 6,
    group: 'frontend-admin',
    keywords: [
      // 核心触发词
      ['管理后台', 3], ['后台页面', 3], ['Admin页面', 3],
      ['CRUD页面', 3], ['管理页面', 3],
      // 技术栈
      ['Ant Design Vue', 3], ['Vxe-Table', 3],
      ['Vue3', 2], ['Vue', 1],
      // 具体场景
      ['表单开发', 2], ['表格开发', 2], ['权限控制', 2],
      ['前端页面', 2], ['组件封装', 2],
      // 英文支持
      ['admin panel', 3], ['admin page', 3]
    ]
  },
  {
    name: 'admin-api-gen',
    priority: 7,
    group: 'frontend-admin-api',
    keywords: [
      // 核心触发词
      ['Admin API', 3], ['管理后台API', 3], ['后台接口同步', 3],
      ['pnpm api:gen', 3], ['admin api:gen', 3],
      // 场景
      ['后台API同步', 3], ['管理端API', 3],
      ['Swagger同步', 2], ['接口类型生成', 2]
    ]
  },
  {
    name: 'app-dev',
    priority: 11,
    group: 'frontend-app-workflow',
    keywords: [
      // 核心触发词 - 完整移动端开发流程
      ['帮我开发App', 5], ['帮我开发移动端', 5], ['完整移动端开发', 5],
      ['移动端开发流程', 5], ['帮我做App', 5], ['开发App功能', 5],
      ['uni-app开发', 4], ['开发小程序', 4], ['帮我写App', 4],
      // 场景触发
      ['新功能App', 3], ['App全流程', 4], ['整个App', 3],
      ['移动端全流程', 4], ['手机端开发', 3],
      // 英文支持
      ['app development', 4], ['full app', 4],
      ['develop app', 4], ['app workflow', 4],
      ['mobile development', 4], ['full mobile', 4]
    ]
  },
  {
    name: 'app-codeing',
    priority: 6,
    group: 'frontend-app',
    keywords: [
      // 核心触发词
      ['uni-app', 3], ['UniApp', 3], ['移动端', 3],
      ['小程序', 3], ['移动应用', 3], ['App开发', 3],
      // 技术栈
      ['wot-design-uni', 3], ['wot-design', 3], ['z-paging', 3],
      // 具体场景
      ['H5', 2], ['手机端', 2], ['跨端', 2], ['移动开发', 3],
      // 英文支持
      ['mobile app', 3], ['mobile page', 3]
    ]
  },
  {
    name: 'app-api-gen',
    priority: 7,
    group: 'frontend-app-api',
    keywords: [
      // 核心触发词
      ['App API', 3], ['移动端API', 3], ['app api:gen', 3],
      ['openapi-ts-request', 3],
      // 场景
      ['TypeScript API', 2], ['API客户端', 2],
      ['API同步', 2], ['API类型', 2], ['接口同步', 2]
    ]
  },

  // ============ 通用技能 ============
  {
    name: 'ui-ux-pro-max',
    priority: 4,
    group: 'design',
    keywords: [
      // 核心触发词
      ['UI设计', 3], ['UX设计', 3], ['界面设计', 3],
      // 设计风格
      ['glassmorphism', 3], ['neumorphism', 3], ['minimalism', 2],
      ['bento grid', 3], ['dark mode', 2],
      // 具体元素
      ['配色', 2], ['字体搭配', 2], ['动画', 2], ['响应式', 2],
      ['布局', 1], ['样式', 1], ['Tailwind', 2],
      // 英文支持
      ['design', 2], ['color palette', 3]
    ]
  },
  {
    name: 'interview',
    priority: 3,
    group: 'planning',
    keywords: [
      // 核心触发词
      ['需求分析', 3], ['需求澄清', 3], ['方案讨论', 3],
      ['头脑风暴', 3], ['功能设计', 2],
      // 场景
      ['探索', 1], ['想法', 1], ['创意', 2], ['澄清', 2],
      // 英文支持
      ['discovery', 3], ['clarify', 2], ['brainstorm', 3]
    ]
  },
  {
    name: 'tech-decision',
    priority: 4,
    group: 'planning',
    keywords: [
      // 核心触发词
      ['技术选型', 3], ['技术对比', 3], ['架构决策', 3],
      ['选择框架', 3], ['选择库', 2],
      // 场景
      ['技术方案', 2], ['技术评审', 3], ['评估', 1],
      ['推荐', 1], ['对比', 2],
      // 英文支持
      ['tech decision', 3], ['compare', 2]
    ]
  },
  {
    name: 'skill-creator',
    priority: 5,
    group: 'meta',
    keywords: [
      // 核心触发词
      ['创建技能', 3], ['技能开发', 3], ['技能编写', 3],
      ['自定义技能', 3], ['新建技能', 3],
      // 技术词
      ['SKILL.md', 3], ['Skill', 2],
      // 场景
      ['扩展能力', 2], ['技能更新', 3], ['更新技能', 3],
      // 英文支持
      ['create skill', 3], ['skill creator', 3]
    ]
  },
  {
    name: 'git-workflow',
    priority: 4,
    group: 'vcs',
    keywords: [
      // 核心触发词
      ['Git', 2], ['版本控制', 2], ['git操作', 3],
      // 具体操作
      ['分支', 1], ['提交', 1], ['合并', 1],
      ['commit', 2], ['merge', 2], ['rebase', 2],
      ['PR', 2], ['pull request', 3],
      // 场景
      ['代码管理', 2], ['冲突解决', 3], ['conflict', 2]
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
