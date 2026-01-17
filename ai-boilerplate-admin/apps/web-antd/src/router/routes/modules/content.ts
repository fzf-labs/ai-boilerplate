import type { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    name: 'Content',
    path: '/content',
    meta: {
      title: '内容管理',
      icon: 'lucide:file-text',
      order: 20,
    },
    children: [
      {
        name: 'ContentArticle',
        path: '/content/article',
        component: () => import('#/views/content/article/index.vue'),
        meta: {
          title: '文章管理',
          icon: 'lucide:newspaper',
        },
      },
    ],
  },
];

export default routes;
