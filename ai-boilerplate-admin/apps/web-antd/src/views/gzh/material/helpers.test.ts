import { describe, expect, it } from 'vitest';

import type { WxGzhMaterialInfo } from '#/api/v1/wx-gzh-material';

import {
  buildMaterialCsv,
  buildMaterialTagOptions,
  filterMaterials,
  formatMaterialTags,
  groupMaterials,
  normalizeMaterialTags,
} from './helpers';

const materials: WxGzhMaterialInfo[] = [
  {
    id: '1',
    name: 'Alpha "One"',
    type: 'image',
    tags: [' news ', '', 'release'],
    updateTime: '2024-01-15 08:30:00',
    mediaId: 'media-1',
    URL: 'https://example.com/a.jpg',
  },
  {
    id: '2',
    name: 'Beta Clip',
    type: 'video',
    tags: ['promo'],
    updateTime: '2024-02-20 12:00:00',
    mediaId: 'media-2',
    URL: 'https://example.com/b.mp4',
  },
  {
    id: '3',
    name: 'Gamma Voice',
    type: 'voice',
    tags: [],
    updateTime: '2024-03-05 09:10:00',
    mediaId: 'media-3',
    URL: 'https://example.com/c.mp3',
  },
];

describe('material helpers', () => {
  it('normalizes and formats tags', () => {
    expect(normalizeMaterialTags(materials[0]?.tags)).toEqual([
      'news',
      'release',
    ]);
    expect(formatMaterialTags(materials[2]?.tags)).toBe('-');
  });

  it('builds csv rows with escaped values and tags', () => {
    const csv = buildMaterialCsv([materials[0] as WxGzhMaterialInfo]);

    expect(csv).toContain('"名称","类型","标签","更新时间","Media ID","访问链接","封面链接","描述"');
    expect(csv).toContain('"Alpha ""One""","图片","news / release","2024-01-15 08:30:00"');
  });

  it('builds sorted unique tag options from the current materials', () => {
    expect(buildMaterialTagOptions(materials)).toEqual([
      { label: 'news', value: 'news' },
      { label: 'promo', value: 'promo' },
      { label: 'release', value: 'release' },
    ]);
  });

  it('filters materials by keyword, type, tag, and date range', () => {
    expect(
      filterMaterials(materials, {
        keyword: 'alpha',
        type: 'image',
        tag: 'news',
        dateRange: ['2024-01-01', '2024-01-31'],
      }),
    ).toEqual([materials[0]]);

    expect(
      filterMaterials(materials, {
        keyword: '',
        type: 'video',
        tag: 'promo',
        dateRange: ['2024-02-01', '2024-02-28'],
      }),
    ).toEqual([materials[1]]);
  });

  it('groups materials by tag and update date', () => {
    expect(groupMaterials(materials, 'tag')).toEqual([
      {
        key: 'news',
        label: 'news',
        items: [materials[0]],
      },
      {
        key: 'promo',
        label: 'promo',
        items: [materials[1]],
      },
      {
        key: '未分组',
        label: '未分组',
        items: [materials[2]],
      },
    ]);

    expect(groupMaterials(materials, 'date')).toEqual([
      {
        key: '2024-01-15',
        label: '2024-01-15',
        items: [materials[0]],
      },
      {
        key: '2024-02-20',
        label: '2024-02-20',
        items: [materials[1]],
      },
      {
        key: '2024-03-05',
        label: '2024-03-05',
        items: [materials[2]],
      },
    ]);
  });
});
