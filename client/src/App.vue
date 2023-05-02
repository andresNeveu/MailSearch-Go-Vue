<script setup>
import { ref } from "vue";
const mails = ref([])

const loadData = async () => {
  const response = await fetch("https://jsonplaceholder.typicode.com/posts")
  const data = await response.json()

  mails.value = data.slice(0, 10)
}

loadData()
</script>

<template>
  <header>
  </header>

  <main>
    <h1>
      <h1 class="text-3xl font-bold text-center p-3">
        Mail Search
      </h1>
    </h1>
    <div class="bg-white flex">
      <section class="bg-white flex w-1/2 p-6">
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
            <tbody v-for="mail in mails" :key="mail.id">
              <tr class="border-b hover:bg-gray-50">
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
      <div class="w-1/2 p-6">
        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Nunc pharetra venenatis justo, non placerat ipsum
        hendrerit et. Interdum et malesuada fames ac ante ipsum primis in faucibus. Nullam eu dui vitae nisi vehicula
        aliquet eu sit amet velit. Fusce eget augue id lectus ultricies accumsan. Donec elit felis, fermentum dictum odio
        ut, ultricies faucibus leo. Morbi vitae est a arcu ultricies laoreet varius sit amet diam. Duis lobortis vehicula
        nibh, et cursus nibh feugiat eu. Sed imperdiet elit nulla, non vehicula leo suscipit et.

        Sed ac pellentesque magna, eget ultrices mauris. Ut nisl risus, sodales ullamcorper ex at, feugiat semper nunc.
        Suspendisse turpis risus, pharetra a nisi id, venenatis volutpat urna. Donec dapibus arcu sit amet semper ornare.
        Aliquam egestas eros nec vestibulum tempor. In hac habitasse platea dictumst. Maecenas porttitor augue lorem, sit
        amet porta tortor pretium eu. Sed tortor libero, blandit vel rhoncus et, dignissim vitae ipsum. Donec dictum
        sapien a felis bibendum ornare a et sapien. Nulla feugiat ultrices sapien quis viverra. Aliquam fermentum nec nibh
        non ultricies. Sed sit amet sem condimentum, volutpat neque a, rutrum nunc.

        Morbi ultrices auctor commodo. In in turpis ac erat tristique ornare eget a ipsum. Maecenas hendrerit iaculis nisl
        eget volutpat. Quisque laoreet nibh at orci vestibulum, quis pellentesque dui venenatis. Praesent rhoncus lacinia
        nisi, quis cursus est. Phasellus laoreet, arcu sit amet laoreet commodo, magna velit eleifend elit, vitae finibus
        odio risus porttitor urna. Vivamus ultricies hendrerit velit, a cursus eros rhoncus eu.

        Aliquam sapien sem, fringilla sed libero id, porttitor pellentesque tortor. Cras accumsan dui placerat
        pellentesque pretium. Mauris fermentum eget tortor eget congue. Ut sit amet dapibus turpis, in auctor nisi.
        Pellentesque habitant morbi tristique senectus et netus et malesuada fames ac turpis egestas. Nam sagittis nunc
        enim, sed ullamcorper turpis pretium vel. Etiam sed augue a dolor cursus eleifend. Nullam ullamcorper in ligula id
        feugiat.

        Nunc molestie, magna vitae tempus eleifend, ante arcu ullamcorper dolor, sed molestie odio justo ut urna. Praesent
        id consequat eros. Nunc et odio vitae leo finibus ultrices. Etiam nec mi nibh. In porta at magna sed iaculis. Nam
        blandit iaculis velit et placerat. Praesent viverra purus ipsum, at venenatis lectus vehicula sit amet. Ut id
        finibus libero. Nam ornare pellentesque libero, id congue sem luctus vitae. Suspendisse sed egestas purus. Duis
        convallis lacus id lorem congue, id ultricies quam malesuada. Donec a ipsum vitae nunc accumsan ultricies. Proin
        dui mauris, mattis vel risus venenatis, dapibus hendrerit dolor.
      </div>
    </div>

  </main>
</template>

<style scoped></style>
