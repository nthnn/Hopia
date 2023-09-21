<script setup lang="ts">
import { computed, reactive } from "vue";
import { RouteLocationNormalizedLoaded, RouterLink, useRoute } from "vue-router";

const route: RouteLocationNormalizedLoaded = useRoute();
const pageData = reactive({
    currentRoute: computed(()=> route.name.toString())
});
</script>

<template>
    <nav class="navbar navbar-expand-lg bg-light shadow-sm fixed-top" data-bs-theme="light">
        <div class="container-fluid">
            <a class="navbar-brand" href="#">Hopia</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarColor01" aria-controls="navbarColor01" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>

            <div class="collapse navbar-collapse" id="navbarColor01">
                <ul class="navbar-nav me-auto">
                    <li class="nav-item">
                        <RouterLink to="/" :class="pageData.currentRoute.toString() == 'home' ? 'active nav-link' : 'nav-link'">Home</RouterLink>
                    </li>
                    <li class="nav-item">
                        <RouterLink to="/docs" :class="pageData.currentRoute.toString() == 'docs' ? 'active nav-link' : 'nav-link'">Documentation</RouterLink>
                    </li>
                    <li class="nav-item">
                        <RouterLink to="/about" :class="pageData.currentRoute.toString() == 'about' ? 'active nav-link' : 'nav-link'">About</RouterLink>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
    <br/><br/><br/>
</template>