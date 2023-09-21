<script setup lang="ts">
import { RouterView } from "vue-router";

import NavigationBar from "./components/NavigationBar.vue";
import Footer from "./components/Footer.vue";
</script>

<template>
    <div class="bg-dark text-white">
        <NavigationBar />

        <div class="container">
            <RouterView />
            <Footer />
        </div>
    </div>
</template>