export const AiModelTypeEnum = {
  CHAT: 1, // 聊天
  IMAGE: 2, // 图像
  VOICE: 3, // 音频
  VIDEO: 4, // 视频
  EMBEDDING: 5, // 向量
  RERANK: 6, // 重排
};
/**
 * AI 平台的枚举
 */
export const AiPlatformEnum = {
  TONG_YI: 'TongYi', // 阿里
  YI_YAN: 'YiYan', // 百度
  DEEP_SEEK: 'DeepSeek', // DeepSeek
  ZHI_PU: 'ZhiPu', // 智谱 AI
  XING_HUO: 'XingHuo', // 讯飞
  SiliconFlow: 'SiliconFlow', // 硅基流动
  OPENAI: 'OpenAI',
  Ollama: 'Ollama',
  STABLE_DIFFUSION: 'StableDiffusion', // Stability AI
  MIDJOURNEY: 'Midjourney', // Midjourney
  SUNO: 'Suno', // Suno AI
};

export interface ImageModel {
  key: string;
  name: string;
  image?: string;
}

export interface ImageSize {
  key: string;
  name?: string;
  style: string;
  width: string;
  height: string;
}

export const OtherPlatformEnum: ImageModel[] = [
  {
    key: AiPlatformEnum.TONG_YI,
    name: '通义万相',
  },
  {
    key: AiPlatformEnum.YI_YAN,
    name: '百度千帆',
  },
  {
    key: AiPlatformEnum.ZHI_PU,
    name: '智谱 AI',
  },
  {
    key: AiPlatformEnum.SiliconFlow,
    name: '硅基流动',
  },
];
export const ImageHotWords = [
  '中国旗袍',
  '古装美女',
  '卡通头像',
  '机甲战士',
  '童话小屋',
  '中国长城',
]; // 图片热词

export const ImageHotEnglishWords = [
  'Chinese Cheongsam',
  'Ancient Beauty',
  'Cartoon Avatar',
  'Mech Warrior',
  'Fairy Tale Cottage',
  'The Great Wall of China',
]; // 图片热词（英文）

export const StableDiffusionSamplers: ImageModel[] = [
  {
    key: 'DDIM',
    name: 'DDIM',
  },
  {
    key: 'DDPM',
    name: 'DDPM',
  },
  {
    key: 'K_DPMPP_2M',
    name: 'K_DPMPP_2M',
  },
  {
    key: 'K_DPMPP_2S_ANCESTRAL',
    name: 'K_DPMPP_2S_ANCESTRAL',
  },
  {
    key: 'K_DPM_2',
    name: 'K_DPM_2',
  },
  {
    key: 'K_DPM_2_ANCESTRAL',
    name: 'K_DPM_2_ANCESTRAL',
  },
  {
    key: 'K_EULER',
    name: 'K_EULER',
  },
  {
    key: 'K_EULER_ANCESTRAL',
    name: 'K_EULER_ANCESTRAL',
  },
  {
    key: 'K_HEUN',
    name: 'K_HEUN',
  },
  {
    key: 'K_LMS',
    name: 'K_LMS',
  },
];

export const StableDiffusionStylePresets: ImageModel[] = [
  {
    key: '3d-model',
    name: '3d-model',
  },
  {
    key: 'analog-film',
    name: 'analog-film',
  },
  {
    key: 'anime',
    name: 'anime',
  },
  {
    key: 'cinematic',
    name: 'cinematic',
  },
  {
    key: 'comic-book',
    name: 'comic-book',
  },
  {
    key: 'digital-art',
    name: 'digital-art',
  },
  {
    key: 'enhance',
    name: 'enhance',
  },
  {
    key: 'fantasy-art',
    name: 'fantasy-art',
  },
  {
    key: 'isometric',
    name: 'isometric',
  },
  {
    key: 'line-art',
    name: 'line-art',
  },
  {
    key: 'low-poly',
    name: 'low-poly',
  },
  {
    key: 'modeling-compound',
    name: 'modeling-compound',
  },
  // neon-punk origami photographic pixel-art tile-texture
  {
    key: 'neon-punk',
    name: 'neon-punk',
  },
  {
    key: 'origami',
    name: 'origami',
  },
  {
    key: 'photographic',
    name: 'photographic',
  },
  {
    key: 'pixel-art',
    name: 'pixel-art',
  },
  {
    key: 'tile-texture',
    name: 'tile-texture',
  },
];

export const StableDiffusionClipGuidancePresets: ImageModel[] = [
  {
    key: 'NONE',
    name: 'NONE',
  },
  {
    key: 'FAST_BLUE',
    name: 'FAST_BLUE',
  },
  {
    key: 'FAST_GREEN',
    name: 'FAST_GREEN',
  },
  {
    key: 'SIMPLE',
    name: 'SIMPLE',
  },
  {
    key: 'SLOW',
    name: 'SLOW',
  },
  {
    key: 'SLOWER',
    name: 'SLOWER',
  },
  {
    key: 'SLOWEST',
    name: 'SLOWEST',
  },
];

export const Dall3SizeList: ImageSize[] = [
  {
    key: '1024x1024',
    name: '1:1',
    width: '1024',
    height: '1024',
    style: 'width: 30px; height: 30px;background-color: #dcdcdc;',
  },
  {
    key: '1024x1792',
    name: '3:5',
    width: '1024',
    height: '1792',
    style: 'width: 30px; height: 50px;background-color: #dcdcdc;',
  },
  {
    key: '1792x1024',
    name: '5:3',
    width: '1792',
    height: '1024',
    style: 'width: 50px; height: 30px;background-color: #dcdcdc;',
  },
];

export const Dall3Models: ImageModel[] = [
  {
    key: 'dall-e-3',
    name: 'DALL·E 3',
    image: `/static/imgs/ai/dall2.jpg`,
  },
  {
    key: 'dall-e-2',
    name: 'DALL·E 2',
    image: `/static/imgs/ai/dall3.jpg`,
  },
];

export const Dall3StyleList: ImageModel[] = [
  {
    key: 'vivid',
    name: '清晰',
    image: `/static/imgs/ai/qingxi.jpg`,
  },
  {
    key: 'natural',
    name: '自然',
    image: `/static/imgs/ai/ziran.jpg`,
  },
];
export const MidjourneyModels: ImageModel[] = [
  {
    key: 'midjourney',
    name: 'MJ',
    image: 'https://bigpt8.com/pc/_nuxt/mj.34a61377.png',
  },
  {
    key: 'niji',
    name: 'NIJI',
    image: 'https://bigpt8.com/pc/_nuxt/nj.ca79b143.png',
  },
];
export const MidjourneyVersions = [
  {
    value: '6.0',
    label: 'v6.0',
  },
  {
    value: '5.2',
    label: 'v5.2',
  },
  {
    value: '5.1',
    label: 'v5.1',
  },
  {
    value: '5.0',
    label: 'v5.0',
  },
  {
    value: '4.0',
    label: 'v4.0',
  },
];

export const NijiVersionList = [
  {
    value: '5',
    label: 'v5',
  },
];

export const MidjourneySizeList: ImageSize[] = [
  {
    key: '1:1',
    width: '1',
    height: '1',
    style: 'width: 30px; height: 30px;background-color: #dcdcdc;',
  },
  {
    key: '3:4',
    width: '3',
    height: '4',
    style: 'width: 30px; height: 40px;background-color: #dcdcdc;',
  },
  {
    key: '4:3',
    width: '4',
    height: '3',
    style: 'width: 40px; height: 30px;background-color: #dcdcdc;',
  },
  {
    key: '9:16',
    width: '9',
    height: '16',
    style: 'width: 30px; height: 50px;background-color: #dcdcdc;',
  },
  {
    key: '16:9',
    width: '16',
    height: '9',
    style: 'width: 50px; height: 30px;background-color: #dcdcdc;',
  },
];

/**
 * AI 图像生成状态的枚举
 */
export const AiImageStatusEnum = {
  IN_PROGRESS: 10, // 进行中
  SUCCESS: 20, // 已完成
  FAIL: 30, // 已失败
};
