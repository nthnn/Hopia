<script setup lang="ts">
import { RouterLink } from "vue-router";
</script>

<template>
    <nav class="navbar navbar-expand-lg bg-primary shadow-sm fixed-bottom" data-bs-theme="dark">
        <div class="container-fluid">
            <a class="navbar-brand" href="#">Hopia</a>
            <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarColor01" aria-controls="navbarColor01" aria-expanded="false" aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>

            <div class="collapse navbar-collapse" id="navbarColor01">
                <ul class="navbar-nav me-auto">
                    <li class="nav-item">
                        <RouterLink to="/" class="nav-link">Home</RouterLink>
                    </li>
                    <li class="nav-item">
                        <RouterLink to="/docs" class="nav-link">Documentation</RouterLink>
                    </li>
                    <li class="nav-item">
                        <RouterLink to="/about" class="nav-link">About</RouterLink>
                    </li>
                </ul>
            </div>
        </div>
    </nav>
</template>