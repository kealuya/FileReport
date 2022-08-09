import { createRouter, createWebHashHistory } from 'vue-router';
export const constantRoutes = [
    {
        path: '/',
        component: () => import('@/components/layouts/Main.vue'),
        redirect: '/home', meta: { title: '首页', icon: 'el-icon-s-home', affix: true },
        children: [
            {
                path: '/home',
                component: () => import('@/components/MainPage.vue'),
                name: 'Home',
            },
            {
                path: '/detail',
                component: () => import('@/components/ProjectDetail.vue'),
                name: 'Detail',
            },
        ]
    }
];
export const asyncRoutes = [];
export const router = createRouter({
    history: createWebHashHistory(),
    routes: constantRoutes
});
export default router;
//# sourceMappingURL=index.js.map