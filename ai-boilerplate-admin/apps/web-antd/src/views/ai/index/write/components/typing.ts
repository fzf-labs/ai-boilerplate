/** 写作点击示例时的数据 */
export const WriteExample = {
  write: {
    prompt: '产品介绍',
    data: '请帮我把这段产品介绍改写得更简洁清晰，适合放在页面详情中。',
  },
  reply: {
    originalContent: '您好，想确认一下明天会议是否按原计划进行。',
    prompt: '请简洁确认并补充提醒事项',
    data: '您好，明天的会议按原计划进行，时间和地点保持不变。请提前准备相关资料，感谢配合。',
  },
};

/**
 * AI 写作类型的枚举
 */
export enum AiWriteTypeEnum {
  WRITING = 1, // 撰写
  REPLY, // 回复
}
