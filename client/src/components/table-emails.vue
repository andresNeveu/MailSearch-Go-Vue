<script setup>
defineProps({
  data: {
    type: Array,
    required: true
  }
});
const emit = defineEmits(['onclick']);

const sendIndex = (index) => {
  emit('onclick', index);
};
</script>

<template>
  <section class="bg-white flex w-3/5 p-4">
    <div class="overflowx-x-auto overflow-y-auto border-x border-t shadow-xl">
      <table class="table-auto w-full">
        <thead class="border-b">
          <tr class="bg-gray-100">
            <th class="text-left p-4 font-medium">Subject</th>
            <th class="text-left p-4 font-medium">To</th>
            <th class="text-left p-4 font-medium">From</th>
          </tr>
        </thead>
        <tbody v-for="(mail, index) in data" :key="mail.id">
          <tr class="border-b hover:bg-gray-50" @click="sendIndex(index)">
            <td class="p-4">
              {{ mail._source.Subject }}
            </td>
            <td class="p-4">
              {{ mail._source.To }}
            </td>
            <td class="p-4">
              {{ mail._source.From }}
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </section>
</template>
