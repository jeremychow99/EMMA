import { createRouter, createWebHistory } from "vue-router";

import login from "../components/login.vue";
import main from "../views/main.vue"
import schedule from "../views/schedule.vue";
import startMaintenance from "../views/startMaintenance.vue";
import endMaintenance from "../views/endMaintenance.vue";
import dashboard from "../views/dashboard.vue";
import equipment from "../views/equipment.vue";
import inventory from "../views/inventory.vue";

const routes = [
    {
        path: "/login",
        name: "login",
        component: login,
    },
    {
        path: "/",
        name: "main",
        component: main
    },
    {
        path: "/schedule",
        name: "schedule",
        component: schedule
    },
    {
        path: "/startMaintenance",
        name: "startMaintenance",
        component: startMaintenance
    },
    {
        path: "/endMaintenance",
        name: "endMaintenance",
        component: endMaintenance
    },
    {
        path: "/dashboard",
        name: "dashboard",
        component: dashboard
    },
    {
        path: "/equipment",
        name: "equipment",
        component: equipment
    },
    {
        path: "/inventory",
        name: "inventory",
        component: inventory
    },
]


const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach( async (to, from) => {
    if(localStorage.token == null && to.name !== 'login'){
        return { name: 'login'}
    }
})

export default router;