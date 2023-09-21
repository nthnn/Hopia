<script setup lang="ts">
import { computed, reactive } from "vue";
import { type RouteLocationNormalizedLoaded, RouterLink, useRoute } from "vue-router";
import {
    CodeBracketIcon,
    DocumentDuplicateIcon,
    GlobeAsiaAustraliaIcon,
    HomeModernIcon
} from "@heroicons/vue/24/outline";

const route: RouteLocationNormalizedLoaded = useRoute();
var pageData = reactive({
    currentRoute: computed(()=> route.name)
});
</script>

<template>
    <nav class="navbar navbar-expand-lg bg-dark shadow-lg fixed-top blur-navbar" data-bs-theme="dark">
        <div class="container-fluid">
            <a class="navbar-brand" href="#">Hopia</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbar-main" aria-controls="navbar-main" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>

            <div class="collapse navbar-collapse" id="navbar-main">
                <ul class="navbar-nav mx-auto">
                    <li class="nav-item">
                        <RouterLink to="/" v-bind:class="pageData.currentRoute == 'home' ? 'active nav-link' : 'nav-link'">
                            <HomeModernIcon class="icon-18" />
                            Home
                        </RouterLink>
                    </li>
                    <li class="nav-item">
                        <RouterLink to="/docs" v-bind:class="pageData.currentRoute == 'docs' ? 'active nav-link' : 'nav-link'">
                            <DocumentDuplicateIcon class="icon-18" />
                            Documentation
                        </RouterLink>
                    </li>
                    <li class="nav-item">
                        <RouterLink to="/about" v-bind:class="pageData.currentRoute == 'about' ? 'active nav-link' : 'nav-link'">
                            <GlobeAsiaAustraliaIcon class="icon-18" />
                            About
                        </RouterLink>
                    </li>
                    <li class="nav-item">
                        <a href="https://github.com/nthnn/Hopia" target="_blank" class="nav-link">
                            <CodeBracketIcon class="icon-18" />
                            GitHub
                        </a>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
    <br/><br/><br/><br/><br/>
</template>

<style scoped>
    .blur-navbar {
        backdrop-filter: blur(5px);
        background-color: rgba(52, 58, 64, 0.8) !important;
    }
</style>