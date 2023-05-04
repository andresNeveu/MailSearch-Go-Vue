<script setup>
defineProps({
    data: {
        type: Array,
        required: true

    }
})
const emit = defineEmits(['onclick'])

const sendIndex = (index) => {
    console.log(index);
    emit('onclick', index)
}

</script>

<template>
    <section class="bg-white flex w-1/2 p-5">
        <div class="overflow-x-auto border-x border-t ">
            <table class="table-auto w-full h-full">
                <thead class="border-b">
                    <tr class="bg-gray-100">
                        <th class="text-left p-4 font-medium">
                            Name
                        </th>
                        <th class="text-left p-4 font-medium">
                            Email
                        </th>
                        <th class="text-left p-4 font-medium">
                            Role
                        </th>
                    </tr>
                </thead>
                <tbody v-for="(mail, index) in data" :key="mail.id">
                    <tr class="border-b hover:bg-gray-50" @click="sendIndex(index)">
                        <td class="p-4">
                            {{ mail.userId }}
                        </td>
                        <td class="p-4">
                            {{ mail.title }}
                        </td>
                        <td class="p-4">
                            {{ mail.id }}
                        </td>
                    </tr>
                </tbody>
            </table>
        </div>
    </section>
</template>