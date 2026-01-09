// @ts-check

import { defineConfig } from '@vben/eslint-config';

export default defineConfig([
  {
    // 忽略自动生成的 API 文件,node_modules目录
    ignores: ['apps/web-antd/src/api/v1/**', 'apps/web-antd/node_modules/**'],
  },
]);
