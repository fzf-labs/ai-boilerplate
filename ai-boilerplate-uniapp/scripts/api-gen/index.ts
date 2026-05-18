import type { APIDataType } from 'openapi-ts-request/dist/generator/type'

import { existsSync } from 'node:fs'
import { mkdir, readdir, stat, writeFile } from 'node:fs/promises'
import { basename, dirname, join, relative, resolve } from 'node:path'
import * as process from 'node:process'
import { fileURLToPath } from 'node:url'

import { generateService } from 'openapi-ts-request'

const __filename = fileURLToPath(import.meta.url)
const __dirname = dirname(__filename)

/**
 * 将 swagger 文件名转换为输出目录名
 * 例如: user.swagger.json -> user
 */
function convertFileNameToDirName(fileName: string): string {
  // 移除 .swagger.json 后缀
  const nameWithoutExt = fileName.replace(/\.swagger\.json$/, '')
  // 将下划线替换为中横线并转小写
  return nameWithoutExt.replace(/_/g, '-').toLowerCase()
}

/**
 * 将大驼峰转换为小驼峰
 * 例如: CreateUser -> createUser
 */
function toCamelCase(str: string): string {
  if (!str)
    return str
  return str.charAt(0).toLowerCase() + str.slice(1)
}

/**
 * 将 kebab-case 转换为 PascalCase
 * 例如: ai-audio-record -> AiAudioRecord
 */
function kebabToPascalCase(str: string): string {
  if (!str)
    return str
  return str
    .split('-')
    .map(word => word.charAt(0).toUpperCase() + word.slice(1))
    .join('')
}

/**
 * 自定义函数名称 hook
 * 从 operationId 中提取后半段并转换为驼峰命名
 * 例如: User_CreateUser -> createUser
 */
function customFunctionName(
  data: APIDataType,
  _prefix?: string,
): string {
  const { operationId } = data
  if (!operationId) {
    return ''
  }

  // 使用下划线分割，取后半段
  const parts = operationId.split('_')
  if (parts.length > 1) {
    // 取后半部分（去掉前缀）
    const suffix = parts.slice(1).join('_')
    // 转换为小驼峰命名
    return toCamelCase(suffix)
  }

  // 如果没有下划线，直接使用整个 operationId 并转为小驼峰
  return toCamelCase(operationId)
}

/**
 * 获取项目根目录路径
 */
function getProjectRoot(): string {
  // 从 scripts/api-gen/index.ts 到项目根目录
  return resolve(__dirname, '../../')
}

/**
 * 递归查找所有 swagger 文件
 */
async function findSwaggerFiles(
  dir: string,
  baseDir: string,
): Promise<Array<{ filePath: string, relativePath: string }>> {
  const files: Array<{ filePath: string, relativePath: string }> = []
  const entries = await readdir(dir)

  for (const entry of entries) {
    const fullPath = join(dir, entry)
    const fileStat = await stat(fullPath)

    if (fileStat.isDirectory()) {
      // 递归处理子目录
      const subFiles = await findSwaggerFiles(fullPath, baseDir)
      files.push(...subFiles)
    }
    else if (entry.endsWith('.swagger.json')) {
      // 计算相对于 baseDir 的路径
      const relativePath = relative(baseDir, dir)
      files.push({
        filePath: fullPath,
        relativePath: relativePath || '.',
      })
    }
  }

  return files
}

/**
 * 生成 index.ts 导出文件
 */
async function generateIndexFiles(
  outputBaseDir: string,
  moduleDirs: Map<string, string[]>,
) {
  // 为每个版本目录生成 index.ts
  for (const [versionDir, modules] of moduleDirs.entries()) {
    const indexPath = join(outputBaseDir, versionDir, 'index.ts')
    const exports = modules
      .sort()
      .map(
        module =>
          `export * as ${kebabToPascalCase(module)} from './${module}'`,
      )
      .join('\n')

    await writeFile(indexPath, `${exports}\n`, 'utf8')
    console.log(`  ✅ 生成 ${versionDir}/index.ts`)
  }

  // 为根目录生成 index.ts（如果有多个版本目录）
  if (moduleDirs.size > 1) {
    const rootIndexPath = join(outputBaseDir, 'index.ts')
    const rootExports = [...moduleDirs.keys()]
      .sort()
      .filter(version => version !== '.') // 排除根目录下的模块
      .map(version => `export * from './${version}'`)
      .join('\n')

    if (rootExports) {
      await writeFile(rootIndexPath, `${rootExports}\n`, 'utf8')
      console.log('  ✅ 生成 index.ts')
    }
  }
}

/**
 * 主生成函数
 */
async function main() {
  const projectRoot = getProjectRoot()
  const swaggerBaseDir = resolve(
    projectRoot,
    '../ai-boilerplate-backend/doc/swagger/app',
  )
  const outputBaseDir = resolve(projectRoot, 'src/api')

  console.log('Swagger 基础目录:', swaggerBaseDir)
  console.log('输出基础目录:', outputBaseDir)

  // 检查 swagger 目录是否存在
  if (!existsSync(swaggerBaseDir)) {
    console.error(`错误: Swagger 目录不存在: ${swaggerBaseDir}`)
    process.exit(1)
  }

  // 递归查找所有 swagger 文件
  const swaggerFiles = await findSwaggerFiles(swaggerBaseDir, swaggerBaseDir)

  if (swaggerFiles.length === 0) {
    console.warn('警告: 未找到任何 .swagger.json 文件')
    return
  }

  console.log(`找到 ${swaggerFiles.length} 个 swagger 文件\n`)

  // 用于收集所有生成的模块目录，按版本分组
  const moduleDirs = new Map<string, string[]>()

  // 为每个 swagger 文件逐个生成
  for (const { filePath, relativePath } of swaggerFiles) {
    const fileName = basename(filePath)
    const dirName = convertFileNameToDirName(fileName)

    // 构建输出目录，保持相对路径结构
    // 例如: v1/user 或 user
    const outputRelativeDir
      = relativePath === '.' ? dirName : join(relativePath, dirName)
    const outputDir = join(outputBaseDir, outputRelativeDir)

    console.log(`\n处理文件: ${relativePath}/${fileName}`)
    console.log(`  输出目录: ${outputRelativeDir}`)

    // 确保输出目录存在
    if (!existsSync(outputDir)) {
      await mkdir(outputDir, { recursive: true })
    }

    // 为单个文件生成接口
    try {
      await generateService({
        describe: `生成 ${fileName} 的接口文件`,
        schemaPath: filePath,
        serversPath: outputDir,
        requestLibPath:
          'import request from \'@/http/vue-query\';\n import { CustomRequestOptions_ } from \'@/http/types\';',
        requestOptionsType: 'CustomRequestOptions_',
        isGenReactQuery: false,
        reactQueryMode: 'vue',
        isGenJavaScript: false,
        isCamelCase: true, // 确保生成的文件和函数都是驼峰命名
        hook: {
          customFunctionName,
        },
      })
      console.log(`  ✅ ${fileName} 生成完成`)

      // 收集模块目录信息
      const versionDir = relativePath === '.' ? '.' : relativePath
      if (!moduleDirs.has(versionDir)) {
        moduleDirs.set(versionDir, [])
      }
      const modules = moduleDirs.get(versionDir)
      if (modules) {
        modules.push(dirName)
      }
    }
    catch (error) {
      console.error(`  ❌ ${fileName} 生成失败:`, error)
      // 继续处理下一个文件，不中断整个流程
    }
  }

  // 生成 index.ts 导出文件
  if (moduleDirs.size > 0) {
    console.log('\n生成导出文件...')
    await generateIndexFiles(outputBaseDir, moduleDirs)
  }

  console.log('\n✅ 所有接口文件生成完成!')
}

// 执行主函数
main().catch((error) => {
  console.error('脚本执行失败:', error)
  process.exit(1)
})
