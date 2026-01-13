import type { AiProviderModelInfo } from '#/api/v1/ai-provider-model';
import type { AiProviderPlatformInfo } from '#/api/v1/ai-provider-platform';

import { getAiProviderModelList } from '#/api/v1/ai-provider-model';
import { getAiProviderPlatformList } from '#/api/v1/ai-provider-platform';

export type ProviderModelOption = AiProviderModelInfo & {
  platformName?: string;
};

export async function fetchProviderModels(
  modelType?: string,
): Promise<ProviderModelOption[]> {
  const [platformRes, modelRes] = await Promise.all([
    getAiProviderPlatformList({
      params: { page: 1, pageSize: 200 },
    }),
    getAiProviderModelList({
      params: { page: 1, pageSize: 200 },
    }),
  ]);

  const platformMap = new Map<string, string>();
  (platformRes?.list || []).forEach((item: AiProviderPlatformInfo) => {
    if (item.id) {
      platformMap.set(item.id, item.platform || item.name || '');
    }
  });

  return (modelRes?.list || [])
    .filter((item) => !modelType || item.modelType === modelType)
    .map((item) => ({
      ...item,
      platformName: item.platformId
        ? platformMap.get(item.platformId)
        : undefined,
    }));
}
