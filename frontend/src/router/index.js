import { createRouter, createWebHistory } from "vue-router";

import login from "../components/login.vue";
import main from "../views/main.vue"
import schedule from "../views/schedule.vue";
import startMaintenance from "../views/startMaintenance.vue";
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
    }
]


const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
