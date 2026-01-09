// Android SDK 版本到系统版本的映射表
const ANDROID_SDK_VERSION_MAP: Record<number, string> = {
  1: '1.0',
  2: '1.1',
  3: '1.5',
  4: '1.6',
  5: '2.0',
  6: '2.0.1',
  7: '2.1',
  8: '2.2',
  9: '2.3',
  10: '2.3.3',
  11: '3.0',
  12: '3.1',
  13: '3.2',
  14: '4.0',
  15: '4.0.3',
  16: '4.1',
  17: '4.2',
  18: '4.3',
  19: '4.4',
  20: '4.4W',
  21: '5.0',
  22: '5.1',
  23: '6.0',
  24: '7.0',
  25: '7.1',
  26: '8.0',
  27: '8.1',
  28: '9.0',
  29: '10.0',
  30: '11.0',
  31: '12.0',
  32: '12L',
  33: '13.0',
  34: '14.0',
  35: '15.0',
};

/**
 * 将 SDK 版本号转换为 Android 版本号
 */
export function getAndroidVersionName(sdkVersion: number | string): string {
  const sdk = Number(sdkVersion);
  return ANDROID_SDK_VERSION_MAP[sdk] || `API ${sdk}`;
}

// 动态加载 app-info-parser CDN 脚本
let appInfoParserLoaded = false;

/**
 * 动态加载 app-info-parser 库
 */
export async function loadAppInfoParser(): Promise<void> {
  if (appInfoParserLoaded) return;

  return new Promise<void>((resolve, reject) => {
    const script = document.createElement('script');
    script.src =
      'https://unpkg.com/app-info-parser@1.1.6/dist/app-info-parser.min.js';
    script.addEventListener('load', () => {
      appInfoParserLoaded = true;
      resolve();
    });
    script.addEventListener('error', () =>
      reject(new Error('Failed to load app-info-parser')),
    );
    document.head.append(script);
  });
}

// 动态加载 SparkMD5 库用于计算 MD5
let sparkMD5Loaded = false;

/**
 * 动态加载 SparkMD5 库
 */
export async function loadSparkMD5(): Promise<void> {
  if (sparkMD5Loaded) return;

  return new Promise<void>((resolve, reject) => {
    const script = document.createElement('script');
    script.src = 'https://unpkg.com/spark-md5@3.0.2/spark-md5.min.js';
    script.addEventListener('load', () => {
      sparkMD5Loaded = true;
      resolve();
    });
    script.addEventListener('error', () =>
      reject(new Error('Failed to load spark-md5')),
    );
    document.head.append(script);
  });
}

/**
 * 计算文件 MD5 值
 */
export async function calculateFileMD5(file: File): Promise<string> {
  await loadSparkMD5();

  // @ts-ignore - SparkMD5 is loaded from CDN
  const spark = new window.SparkMD5.ArrayBuffer();
  const chunkSize = 2_097_152; // 2MB chunks
  const chunks = Math.ceil(file.size / chunkSize);

  for (let currentChunk = 0; currentChunk < chunks; currentChunk++) {
    const start = currentChunk * chunkSize;
    const end = Math.min(start + chunkSize, file.size);
    const blob = file.slice(start, end);
    const arrayBuffer = await blob.arrayBuffer();
    spark.append(arrayBuffer);
  }

  return spark.end();
}

/**
 * 解析安装包信息的结果类型
 */
export interface ParsedPackageInfo {
  buildNum?: number;
  version?: string;
  packageSize: number;
  packageMd5: string;
  minOsVersion?: string;
}

/**
 * 解析安装包文件，提取版本信息
 */
export async function parsePackageFile(file: File): Promise<ParsedPackageInfo> {
  // 并行执行：解析安装包信息和计算 MD5
  const [result, packageMd5] = await Promise.all([
    (async () => {
      await loadAppInfoParser();
      // @ts-ignore - AppInfoParser is loaded from CDN
      const parser = new window.AppInfoParser(file);
      return await parser.parse();
    })(),
    calculateFileMD5(file),
  ]);

  // 提取 buildNum、version 和 minOsVersion（仅支持 Android APK）
  let buildNum: number | undefined;
  let version: string | undefined;
  let minOsVersion: string | undefined;

  // 对于 Android APK
  if (result.versionCode) {
    buildNum = Number(result.versionCode);
    version = result.versionName;
    // Android 最低 SDK 版本转换为系统版本名称
    // minSdkVersion 在 usesSdk 对象下
    const sdkVersion = result.usesSdk?.minSdkVersion || result.minSdkVersion;
    if (sdkVersion) {
      minOsVersion = getAndroidVersionName(sdkVersion);
    } else {
      console.warn('未找到 minSdkVersion 字段');
    }
  }

  // 计算文件大小(MB)
  const packageSize = Number((file.size / 1024 / 1024).toFixed(2));

  return {
    buildNum,
    version,
    packageSize,
    packageMd5,
    minOsVersion,
  };
}
