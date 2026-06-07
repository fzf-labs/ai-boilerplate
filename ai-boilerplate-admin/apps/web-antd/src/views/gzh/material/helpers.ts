import type { WxGzhMaterialInfo } from '#/api/v1/wx-gzh-material'

export const MaterialType = {
  IMAGE: 'image',
  VOICE: 'voice',
  VIDEO: 'video',
} as const

export type MaterialTypeValue = (typeof MaterialType)[keyof typeof MaterialType]

export type MaterialGroupMode = 'date' | 'none' | 'tag'

export type MaterialGroup = {
  key: string
  label: string
  items: WxGzhMaterialInfo[]
}

export type MaterialToolbarFilters = {
  keyword: string
  type?: MaterialTypeValue
  dateRange?: [string, string]
  groupMode?: MaterialGroupMode
  tag?: string
}

export type MaterialTagOption = {
  label: string
  value: string
}

export const materialTypeLabels: Record<MaterialTypeValue, string> = {
  [MaterialType.IMAGE]: '图片',
  [MaterialType.VOICE]: '语音',
  [MaterialType.VIDEO]: '视频',
}

export const materialTypeColors: Record<MaterialTypeValue, string> = {
  [MaterialType.IMAGE]: 'blue',
  [MaterialType.VOICE]: 'green',
  [MaterialType.VIDEO]: 'orange',
}

export const materialTypeOptions = [
  { label: materialTypeLabels[MaterialType.IMAGE], value: MaterialType.IMAGE },
  { label: materialTypeLabels[MaterialType.VOICE], value: MaterialType.VOICE },
  { label: materialTypeLabels[MaterialType.VIDEO], value: MaterialType.VIDEO },
]

export const materialGroupModeOptions: Array<{
  label: string
  value: MaterialGroupMode
}> = [
  { label: '平铺', value: 'none' },
  { label: '按标签', value: 'tag' },
  { label: '按日期', value: 'date' },
]

export const getMaterialTypeLabel = (type?: string) => {
  if (!type) {
    return ''
  }

  return materialTypeLabels[type as MaterialTypeValue] || type
}

export const getMaterialTypeColor = (type?: string) => {
  if (!type) {
    return 'default'
  }

  return materialTypeColors[type as MaterialTypeValue] || 'default'
}

export const normalizeMaterialTags = (tags?: string[]) => {
  return (tags || []).map((tag) => tag.trim()).filter(Boolean)
}

export const formatMaterialTags = (tags?: string[]) => {
  const normalized = normalizeMaterialTags(tags)

  return normalized.length > 0 ? normalized.join(' / ') : '-'
}

const normalizeSearchText = (value?: string) => {
  return value?.trim().toLowerCase() || ''
}

const getMaterialDateKey = (material: WxGzhMaterialInfo) => {
  return (
    material.updateTime ||
    material.updatedAt ||
    material.createdAt ||
    ''
  ).slice(0, 10)
}

const matchesMaterialDateRange = (
  material: WxGzhMaterialInfo,
  dateRange?: [string, string],
) => {
  if (!dateRange) {
    return true
  }

  const [startDate, endDate] = dateRange
  const materialDate = getMaterialDateKey(material)

  if (!materialDate) {
    return false
  }

  return (!startDate || materialDate >= startDate) && (!endDate || materialDate <= endDate)
}

export const buildMaterialTagOptions = (materials: WxGzhMaterialInfo[]) => {
  const tags = new Set<string>()

  materials.forEach((material) => {
    normalizeMaterialTags(material.tags).forEach((tag) => tags.add(tag))
  })

  return Array.from(tags)
    .sort((left, right) => left.localeCompare(right))
    .map((tag) => ({
      label: tag,
      value: tag,
    }))
}

export const filterMaterials = (
  materials: WxGzhMaterialInfo[],
  filters: MaterialToolbarFilters,
) => {
  const keyword = normalizeSearchText(filters.keyword)
  const type = filters.type
  const selectedTag = normalizeSearchText(filters.tag)

  return materials.filter((material) => {
    if (type && material.type !== type) {
      return false
    }

    if (selectedTag) {
      const tags = normalizeMaterialTags(material.tags).map((tag) =>
        normalizeSearchText(tag),
      )

      if (!tags.includes(selectedTag)) {
        return false
      }
    }

    if (!matchesMaterialDateRange(material, filters.dateRange)) {
      return false
    }

    if (!keyword) {
      return true
    }

    const searchableText = [
      material.name,
      material.mediaId,
      material.description,
      ...normalizeMaterialTags(material.tags),
    ]
      .filter(Boolean)
      .map((value) => normalizeSearchText(value))
      .join(' ')

    return searchableText.includes(keyword)
  })
}

const escapeCsvCell = (value: string) => {
  const escaped = value.replaceAll('"', '""')
  return `"${escaped}"`
}

const materialCsvColumns = [
  '名称',
  '类型',
  '标签',
  '更新时间',
  'Media ID',
  '访问链接',
  '封面链接',
  '描述',
]

export const buildMaterialCsv = (materials: WxGzhMaterialInfo[]) => {
  const rows = materials.map((material) => {
    const tags = normalizeMaterialTags(material.tags).join(' / ')
    const updatedAt = material.updateTime || material.updatedAt || material.createdAt || '-'

    return [
      material.name || '-',
      getMaterialTypeLabel(material.type),
      tags || '-',
      updatedAt,
      material.mediaId || '-',
      material.URL || '-',
      material.coverURL || '-',
      material.description || '-',
    ].map(escapeCsvCell).join(',')
  })

  return [materialCsvColumns.map(escapeCsvCell).join(','), ...rows].join('\n')
}

const groupByTag = (materials: WxGzhMaterialInfo[]) => {
  const map = new Map<string, WxGzhMaterialInfo[]>()

  materials.forEach((material) => {
    const tags = normalizeMaterialTags(material.tags)
    const bucket = tags[0] || '未分组'

    if (!map.has(bucket)) {
      map.set(bucket, [])
    }

    map.get(bucket)?.push(material)
  })

  return Array.from(map.entries()).map(([label, items]) => ({
    key: label,
    label,
    items,
  }))
}

const groupByDate = (materials: WxGzhMaterialInfo[]) => {
  const map = new Map<string, WxGzhMaterialInfo[]>()

  materials.forEach((material) => {
    const rawDate = material.updateTime || material.updatedAt || material.createdAt || ''
    const bucket = rawDate ? rawDate.slice(0, 10) : '未分组'

    if (!map.has(bucket)) {
      map.set(bucket, [])
    }

    map.get(bucket)?.push(material)
  })

  return Array.from(map.entries()).map(([label, items]) => ({
    key: label,
    label,
    items,
  }))
}

export const groupMaterials = (
  materials: WxGzhMaterialInfo[],
  mode: MaterialGroupMode,
) => {
  if (mode === 'tag') {
    return groupByTag(materials)
  }

  if (mode === 'date') {
    return groupByDate(materials)
  }

  return [{
    key: 'all',
    label: '全部素材',
    items: materials,
  }]
}
