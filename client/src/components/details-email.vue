<script setup>
defineProps({
    detail: {
        type: Object,
        required: true
    }
})
</script>

<template>
    <section class="bg-white flex w-2/5 p-4">
        <div class="overflow-x-auto">
            <h1>
                <strong>
                    {{ detail.title }}
                </strong>
            </h1>
            <p>
                {{ detail.body }}
            </p>
        </div>
    </section>
</template>