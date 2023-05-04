<script setup>
defineProps({
    detail: {
        type: String,
        required: true
    }
})
</script>

<template>
    <div class="w-1/2 p-5">
        {{ detail }}
    </div>
</template>