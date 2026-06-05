import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';

import { dataURLtoBlob, urlToBase64 } from '../download';

describe('download helpers', () => {
  let originalCreateElement: typeof document.createElement;

  beforeEach(() => {
    const canvas = {
      getContext: vi.fn(() => ({
        drawImage: vi.fn(),
      })),
      height: 0,
      toDataURL: vi.fn(() => 'data:image/png;base64,Zm9v'),
      width: 0,
    };

    class MockImage {
      crossOrigin = '';
      height = 1;
      width = 1;

      private listeners: Record<string, Array<(event: Event) => void>> = {};

      addEventListener(
        type: string,
        listener: (event: Event) => void,
      ): void {
        if (!this.listeners[type]) {
          this.listeners[type] = [];
        }
        this.listeners[type].push(listener);
      }

      set src(_value: string) {
        queueMicrotask(() => {
          this.listeners.error?.forEach((listener) => {
            listener(new Event('error'));
          });
        });
      }
    }

    originalCreateElement = document.createElement.bind(document);
    vi.stubGlobal('Image', MockImage as unknown as typeof Image);
    vi.spyOn(document, 'createElement').mockImplementation((tagName: string) => {
      if (String(tagName).toLowerCase() === 'canvas') {
        return canvas as unknown as HTMLCanvasElement;
      }
      return originalCreateElement(tagName);
    });
  });

  afterEach(() => {
    vi.restoreAllMocks();
    vi.unstubAllGlobals();
  });

  it('converts a data URL to a Blob', async () => {
    const blob = dataURLtoBlob('data:text/plain;base64,SGVsbG8=');

    expect(blob.type).toBe('text/plain');
    await expect(blob.text()).resolves.toBe('Hello');
  });

  it('rejects when the image fails to load', async () => {
    await expect(
      urlToBase64('https://example.com/broken.png'),
    ).rejects.toThrow('Failed to load image: https://example.com/broken.png');
  });
});
