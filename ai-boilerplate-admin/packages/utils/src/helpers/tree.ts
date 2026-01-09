interface TreeNode {
  [key: string]: any;
  children?: TreeNode[];
}

/**
 * 构造树型结构数据
 *
 * @param {*} data 数据源
 * @param {*} id id字段 默认 'id'
 * @param {*} pid 父节点字段 默认 'pid'
 * @param {*} children 孩子节点字段 默认 'children'
 */
function handleTree(
  data: TreeNode[],
  id: string = 'id',
  pid: string = 'pid',
  children: string = 'children',
): TreeNode[] {
  if (!Array.isArray(data)) {
    console.warn('data must be an array');
    return [];
  }
  const config = {
    id,
    pid,
    childrenList: children,
  };
  const childrenListMap: Record<number | string, TreeNode[]> = {};
  const nodeIds: Record<number | string, TreeNode> = {};
  const tree: TreeNode[] = [];

  // 1. 数据预处理
  // 1.1 第一次遍历，生成 childrenListMap 和 nodeIds 映射
  for (const d of data) {
    const pId = d[config.pid];
    if (childrenListMap[pId] === undefined) {
      childrenListMap[pId] = [];
    }
    nodeIds[d[config.id]] = d;
    childrenListMap[pId].push(d);
  }
  // 1.2 第二次遍历，找出根节点
  for (const d of data) {
    const pId = d[config.pid];
    if (nodeIds[pId] === undefined) {
      tree.push(d);
    }
  }

  // 2. 构建树结：递归构建子节点
  const adaptToChildrenList = (node: TreeNode): void => {
    const nodeId = node[config.id];
    if (childrenListMap[nodeId]) {
      node[config.childrenList] = childrenListMap[nodeId];
      // 递归处理子节点
      for (const child of node[config.childrenList]) {
        adaptToChildrenList(child);
      }
    }
  };

  // 3. 从根节点开始构建完整树
  for (const rootNode of tree) {
    adaptToChildrenList(rootNode);
  }

  return tree;
}

export { handleTree };
