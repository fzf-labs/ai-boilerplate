import type { Ref } from 'vue';

import type { InfraFileApi } from '#/api/infra/file/data';

import { computed, unref } from 'vue';

import TOS from '@volcengine/tos-sdk';

import { getOSSDefaultPolicy } from '#/api/infra/file/data';

import { generateSafeFileName, generateUploadPath } from './helper';

/**
 * 进度回调函数类型
 */
export type ProgressCallback = (e: { loaded: number; total: number }) => void;

export function useUploadType({
  acceptRef,
  helpTextRef,
  maxNumberRef,
  maxSizeRef,
}: {
  acceptRef: Ref<string[]>;
  helpTextRef: Ref<string>;
  maxNumberRef: Ref<number>;
  maxSizeRef: Ref<number>;
}) {
  // 文件类型限制
  const getAccept = computed(() => {
    const accept = unref(acceptRef);
    if (accept && accept.length > 0) {
      return accept;
    }
    return [];
  });
  const getStringAccept = computed(() => {
    return unref(getAccept)
      .map((item) => {
        return item.indexOf('/') > 0 || item.startsWith('.')
          ? item
          : `.${item}`;
      })
      .join(',');
  });

  // 支持jpg、jpeg、png格式，不超过2M，最多可选择10张图片，。
  const getHelpText = computed(() => {
    const helpText = unref(helpTextRef);
    if (helpText) {
      return helpText;
    }
    const helpTexts: string[] = [];

    const accept = unref(acceptRef);
    if (accept.length > 0) {
      helpTexts.push(`支持${accept.join(',')}格式`);
    }

    const maxSize = unref(maxSizeRef);
    if (maxSize) {
      helpTexts.push(`不超过${maxSize}MB`);
    }

    const maxNumber = unref(maxNumberRef);
    if (maxNumber && maxNumber !== Infinity) {
      helpTexts.push(`最多选择${maxNumber}个文件`);
    }
    return helpTexts.join('，');
  });
  return { getAccept, getStringAccept, getHelpText };
}

/**
 * OSS 对象存储上传 Hook
 */
export function useOSSUpload(scene: string = 'upload') {
  /**
   * 上传文件到对象存储
   */
  async function uploadToOSS(file: File, onProgress?: ProgressCallback) {
    try {
      // 1. 生成文件名称和路径
      const fileName = generateSafeFileName(file.name);
      const filePath = generateUploadPath(scene, fileName);
      // 2. 获取OSS上传策略
      const policyInfo = await getOSSDefaultPolicy(
        fileName,
        filePath,
        file.size,
      );
      // 3. 根据存储引擎进行上传 - 当前只支持火山云
      if (policyInfo.storage !== 'volcengine' || !policyInfo.volcengine) {
        throw new Error(
          `当前只支持火山云对象存储，不支持: ${policyInfo.storage}`,
        );
      }

      return await uploadToVolcengine(
        file,
        policyInfo.volcengine,
        policyInfo.fileId,
        filePath,
        onProgress,
      );
    } catch (error) {
      console.error('上传失败:', error);
      throw error;
    }
  }

  return {
    uploadToOSS,
  };
}

/**
 * 上传到火山云 TOS
 */
/**
 * 统一上传 Hook - 提供标准的上传接口
 */
export function useUpload(scene?: string) {
  const { uploadToOSS } = useOSSUpload(scene);

  /**
   * HTTP 请求上传文件
   */
  async function httpRequest(file: File, onProgress?: ProgressCallback) {
    return await uploadToOSS(file, onProgress);
  }

  return {
    httpRequest,
  };
}

async function uploadToVolcengine(
  file: File,
  policy: InfraFileApi.VolcenginePolicy,
  fileId: string,
  filePath: string,
  onProgress?: ProgressCallback,
) {
  // 创建 TOS 客户端
  const client = new TOS({
    accessKeyId: policy.accessKeyId,
    accessKeySecret: policy.secretAccessKey,
    stsToken: policy.sessionToken,
    region: policy.region,
    endpoint: policy.endpoint,
    bucket: policy.bucket,
  });

  try {
    // 执行上传
    const result = await client.putObject({
      key: filePath,
      body: file,
      contentType: file.type,
      progress: onProgress
        ? (progressEvent: any) => {
            if (progressEvent.total > 0) {
              onProgress({
                loaded: progressEvent.loaded || 0,
                total: progressEvent.total,
              });
            }
          }
        : undefined,
    });

    const fileUrl = policy.customDomain
      ? `https://${policy.customDomain}/${filePath}`
      : `https://${policy.bucket}.${policy.endpoint}/${filePath}`;

    return {
      id: fileId,
      url: fileUrl,
      requestId: result.requestId || '',
    };
  } catch (error) {
    console.error('火山云上传失败:', error);
    throw error;
  }
}
