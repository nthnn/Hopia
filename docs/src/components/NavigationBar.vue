<script setup lang="ts">
import { computed, reactive } from "vue";
import { RouteLocationNormalizedLoaded, RouterLink, useRoute } from "vue-router";
import {
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
    <nav class="navbar navbar-expand-lg bg-light shadow-sm fixed-top" data-bs-theme="light">
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
                </ul>
            </div>
        </div>
    </nav>
    <br/><br/><br/><br/><br/>
</template>