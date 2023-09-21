import "./assets/main.css";
import "../node_modules/bootstrap/dist/js/bootstrap";
import "../node_modules/bootswatch/dist/lux/bootstrap.min.css";

import { createApp, type App } from "vue";
import WebApp from "./App.vue";
import router from "./router";

const app: App<Element> = createApp(WebApp);
app.use(router);
app.mount("#app");