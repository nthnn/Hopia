import { createRouter, createWebHistory } from "vue-router";
import DisclaimerView from "../views/DisclaimerView.vue";
import DocumentationView from "../views/DocumentationView.vue";
import HomeView from "../views/HomeView.vue";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeView
    },
    {
      path: "/docs",
      name: "docs",
      component: DocumentationView
    },
    {
      path: "/disclaimer",
      name: "disclaimer",
      component: DisclaimerView
    }
  ]
});

export default router;
