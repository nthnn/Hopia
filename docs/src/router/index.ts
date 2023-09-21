import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import DocumentationView from "../views/DocumentationView.vue";
import AboutView from "../views/AboutView.vue";

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
      path: "/about",
      name: "about",
      component: AboutView
    }
  ]
});

export default router;
