import {createRouter, createWebHashHistory, RouteRecordRaw} from 'vue-router'

export const constantRoutes: Array<RouteRecordRaw> = [

    {
        path: '/',
        component: () => import('@/components/layouts/Main.vue'),
        redirect: '/home', meta: {title: '首页', icon: 'el-icon-s-home', affix: true},
        children: [
            {
                path: '/home',
                component: () => import('@/components/HelloWorld.vue'),
                name: 'Home',

            },

        ]
    }
]

export const asyncRoutes = []

export const router = createRouter({
    history: createWebHashHistory(),
    routes: constantRoutes
})

export default router
