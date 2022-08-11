import InputView from "@/views/InputView.vue";
import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import ReCheckView from "@/views/ReCheckView.vue";
import ScannerLoginView from "@/views/ScannerLoginView.vue";
import ScannerView from "@/views/ScannerView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView,
    },
    {
      path: "/start",
      name: "input",
      component: InputView,
    },
    {
      path: "/recheck",
      name: "recheck",
      component: ReCheckView,
    },
    {
      path: "/login-scan",
      name: "loginScan",
      component: ScannerLoginView,
    },
    {
      path: "/scanner",
      name: "scanner",
      component: ScannerView,
    },
  ],
});

export default router;
