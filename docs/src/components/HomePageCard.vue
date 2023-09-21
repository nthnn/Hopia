<script setup lang="ts">
const props = defineProps<{
    title: string,
    content: string
}>();
</script>

<template>
    <div class="col-lg-3">
        <div class="card card-body border border-1 bg-dark text-white">
            <h2>{{ props.title }}</h2>
            <p>{{ props.content }}</p>
        </div>
        <br/>
    </div>
</template>