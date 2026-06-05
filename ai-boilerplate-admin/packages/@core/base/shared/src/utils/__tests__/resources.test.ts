import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';

import { loadScript } from '../resources';

const testJsPath = 'test-script.js';

describe('loadScript', () => {
  let scripts: HTMLScriptElement[];

  beforeEach(() => {
    scripts = [];

    vi.spyOn(document.head, 'append').mockImplementation((...nodes) => {
      scripts.push(...(nodes as HTMLScriptElement[]));
    });
    vi.spyOn(document.head, 'querySelectorAll').mockImplementation(
      (selector: string) =>
        findScripts(selector) as unknown as NodeListOf<HTMLScriptElement>,
    );
    vi.spyOn(document, 'querySelector').mockImplementation(
      (selector: string) => findScripts(selector)[0] ?? null,
    );
  });

  afterEach(() => {
    vi.restoreAllMocks();
  });

  it('should resolve when the script loads successfully', async () => {
    const promise = loadScript(testJsPath);

    // 此时脚本元素已被创建并插入
    const script = document.querySelector(
      `script[src="${testJsPath}"]`,
    ) as HTMLScriptElement;
    expect(script).toBeTruthy();

    // 模拟加载成功
    script.dispatchEvent(new Event('load'));

    // 等待 promise resolve
    await expect(promise).resolves.toBeUndefined();
  });

  it('should not insert duplicate script and resolve immediately if already loaded', async () => {
    // 先手动插入一个相同 src 的 script
    const duplicatePath = 'bar.js';
    const existing = document.createElement('script');
    existing.src = duplicatePath;
    scripts.push(existing);

    // 再次调用
    const promise = loadScript(duplicatePath);

    // 立即 resolve
    await expect(promise).resolves.toBeUndefined();

    // head 中只保留一个
    const matchedScripts = document.head.querySelectorAll(
      `script[src="${duplicatePath}"]`,
    );
    expect(matchedScripts).toHaveLength(1);
  });

  it('should reject when the script fails to load', async () => {
    const errorPath = 'error.js';
    const promise = loadScript(errorPath);

    const script = document.querySelector(
      `script[src="${errorPath}"]`,
    ) as HTMLScriptElement;
    expect(script).toBeTruthy();

    // 模拟加载失败
    script.dispatchEvent(new Event('error'));

    await expect(promise).rejects.toThrow(
      `Failed to load script: ${errorPath}`,
    );
  });

  it('should handle multiple concurrent calls and only insert one script tag', async () => {
    const p1 = loadScript(testJsPath);
    const p2 = loadScript(testJsPath);

    const script = document.querySelector(
      `script[src="${testJsPath}"]`,
    ) as HTMLScriptElement;
    expect(script).toBeTruthy();

    // 触发一次 load，两个 promise 都应该 resolve
    script.dispatchEvent(new Event('load'));

    await expect(p1).resolves.toBeUndefined();
    await expect(p2).resolves.toBeUndefined();

    // 只插入一次
    const matchedScripts = document.head.querySelectorAll(
      `script[src="${testJsPath}"]`,
    );
    expect(matchedScripts).toHaveLength(1);
  });

  function findScripts(selector: string): HTMLScriptElement[] {
    const match = selector.match(/^script\[src="(.+)"\]$/);
    if (!match) {
      return [];
    }
    return scripts.filter((script) => script.getAttribute('src') === match[1]);
  }
});
