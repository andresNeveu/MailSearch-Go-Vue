<script setup>
import { ref } from "vue";
import TableEmails from "./components/table-emails.vue";
import DetailsEmail from "./components/details-email.vue";
const mails = ref([])
const search = ref('')
const emailSelect = ref('Ejemplo')

const loadData = async () => {
  const response = await fetch("https://jsonplaceholder.typicode.com/posts")
  const data = await response.json()

  mails.value = data.slice(0, 10)
}

const submit = () => {
  console.log(search.value);
}

const handleIndex = (index) => {
  emailSelect.value = mails.value[index].body
  console.log(mails.value[index].title);
}

loadData()
</script>

<template>
  <header>
  </header>

  <main>
    <h1>
      <h1 class="text-3xl font-bold text-center p-3">
        Mail Search {{ search }}
      </h1>
    </h1>
    <form @submit.prevent="submit" class="ml-4 mr-4 flex items-center">
      <div class="flex items-center w-full">
        <input type=" text" placeholder="Search..." v-model="search" class="w-full rounded-lg border border-gray-400 p-2">
        <button type="submit" value="Submit"
          class="ml-4 rounded-lg bg-blue-500 p-2 text-white hover:bg-blue-600">submit</button>
      </div>
    </form>
    <div class="bg-white flex">
      <TableEmails :data="mails" @onclick="handleIndex" />
      <DetailsEmail :detail="emailSelect" />
    </div>

  </main>
</template>

<style scoped></style>
