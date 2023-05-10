<script setup>
defineProps({
  detail: {
    type: Object,
    required: true
  }
});
</script>

<template>
  <section class="bg-white flex w-2/5 p-4">
    <div class="overflow-x-auto overflow-y-auto shadow-xl px-8 py-12 sm:px-12 lg:px-8 w-full">
      <h1 class="text-2xl font-semibold">
        {{ detail.title }}
      </h1>
      <p>
        {{ detail.body }}
      </p>
    </div>
  </section>
</template>
