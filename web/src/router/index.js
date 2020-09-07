import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from '../components/Index.vue'
import Task from "@/views/Task";
import System from "@/views/System";

Vue.use(VueRouter)

const routes = [
    {
        path: '/',
        name: 'Index',
        component: Index,
        children: [
            {
                path: '/task',
                name: 'Task',
                component: Task,
            },
            {
                path: '/system',
                name: 'System',
                component: System,
            }
        ]
    }
]

const router = new VueRouter({
    mode: 'history',
    base: process.env.BASE_URL,
    routes
})

export default router
