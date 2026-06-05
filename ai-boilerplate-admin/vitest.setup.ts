function createMemoryStorage(): Storage {
  const store = new Map<string, string>();

  return {
    get length() {
      return store.size;
    },
    clear() {
      store.clear();
    },
    getItem(key: string) {
      return store.get(key) ?? null;
    },
    key(index: number) {
      return [...store.keys()][index] ?? null;
    },
    removeItem(key: string) {
      store.delete(key);
    },
    setItem(key: string, value: string) {
      store.set(key, String(value));
    },
  };
}

for (const storageName of ['localStorage', 'sessionStorage'] as const) {
  const storage = createMemoryStorage();

  Object.defineProperty(globalThis, storageName, {
    configurable: true,
    value: storage,
  });

  if (globalThis.window) {
    Object.defineProperty(globalThis.window, storageName, {
      configurable: true,
      value: storage,
    });
  }
}
