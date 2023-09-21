<script setup lang="ts">
import { RouterView } from "vue-router";

import NavigationBar from "./components/NavigationBar.vue";
import Footer from "./components/Footer.vue";
</script>

<template>
    <NavigationBar />

    <div class="container">
        <RouterView />
        <Footer />
    </div>
</template>