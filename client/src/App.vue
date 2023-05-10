<script setup>
import { ref } from 'vue';
import TableEmails from './components/table-emails.vue';
import DetailsEmail from './components/details-email.vue';
const mails = ref([]);
const search = ref('');
const emailSelect = ref({ title: '', body: '' });

const loadData = async () => {
  const value = search.value === '' ? 'a' : search.value;
  const response = await fetch('http://localhost:3333/search?value=' + value);
  const data = await response.json();

  if (data.hits.total.value === 0) {
    alert('Error: match not found ');
  } else {
    mails.value = data.hits.hits;
  }
};

const submit = async () => {
  emailSelect.value.title = '';
  emailSelect.value.body = '';
  await loadData();
};

const handleIndex = (index) => {
  emailSelect.value.title = mails.value[index]._source.Subject;
  emailSelect.value.body = mails.value[index]._source.Body;
};

loadData();
</script>

<template>
  <header>
    <h1 class="text-3xl font-bold text-center p-3">Mail Search</h1>
  </header>
  <main class="h-5/6">
    <form @submit.prevent="submit" class="ml-4 mr-4 pb-2">
      <div class="flex items-center mx-auto bg-white rounded-lg w-full">
        <input
          type=" text"
          placeholder="Type to search..."
          v-model="search"
          class="w-full px-4 py-1 border border-gray-400 text-gray-800 focus:outline-none h-12"
        />
        <button
          type="submit"
          value="Submit"
          class="flex items-center bg-blue-500 justify-center w-24 h-12 text-white rounded-r-lg"
        >
          <p class="pr-2">Search</p>
          <svg
            class="w-5 h-5"
            fill="none"
            stroke="currentColor"
            viewBox="0 0 24 24"
            xmlns="http://www.w3.org/2000/svg"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
            ></path>
          </svg>
        </button>
      </div>
    </form>
    <div class="flex h-full pt-4">
      <TableEmails :data="mails" @onclick="handleIndex" />
      <DetailsEmail :detail="emailSelect" />
    </div>
  </main>
</template>

<style scoped></style>
