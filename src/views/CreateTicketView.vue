<script setup lang="ts">
import { ref } from 'vue'
const title = ref<string | null>(null)
const category = ref<string | null>(null)
const priority = ref<string | null>(null)
const affectedAsset = ref<string | null>(null)
const description = ref<string | null>(null)

function createTicket() {
  const missingFields: string[] = []

  if (!title.value) {
    missingFields.push('Title')
  }

  if (!category.value) {
    missingFields.push('Category')
  }

  if (!priority.value) {
    missingFields.push('Priority')
  }

  if (!description.value) {
    missingFields.push('Description')
  }

  if (missingFields.length > 0) {
    alert(`Please fill in: ${missingFields.join(', ')}`)
  }
}
</script>

<template>
  <div
    class="h-screen p-4 grid lg:grid-cols-[minmax(0,2fr)_minmax(0,1fr)] md:grid-cols-[minmax(0,1fr)] items-center"
  >
    <form class="w-full rounded-lg border shadow-md p-4" @submit.prevent="createTicket">
      <p class="text-lg font-bold mb-6">Create New Ticket</p>
      <div class="flex flex-row gap-4">
        <div class="w-80 md:w-150 h-full flex flex-col gap-2">
          <div class="flex flex-col gap-2 mb-2">
            <label class="px-1">Title</label>
            <input
              v-model="title"
              placeholder="Enter title..."
              type="text"
              class="w-full outline-none rounded-sm border px-4 py-2"
            />
          </div>
          <div class="flex flex-col gap-2 mb-2">
            <label class="px-1">Category</label>
            <select class="w-full rounded-sm border px-4 py-2" v-model="category">
              <option value="hardware">Hardware</option>
              <option value="software">Software</option>
              <option value="network-internet">Network & Internet</option>
              <option value="account-access">Account & Access</option>
              <option value="email-communication">Email & Communication</option>
              <option value="security">Security</option>
              <option value="file-data">File & Data</option>
              <option value="other">Other</option>
            </select>
          </div>
          <div class="flex flex-col gap-2 mb-2">
            <label class="px-1">Priority</label>
            <select class="w-full rounded-sm border px-4 py-2" v-model="priority">
              <option value="low">Low</option>
              <option value="medium">Medium</option>
              <option value="high">High</option>
              <option value="critical">Critical</option>
            </select>
          </div>
          <div class="flex flex-col gap-2 mb-2">
            <label class="px-1">Affected Asset</label>
            <input
              v-model="affectedAsset"
              placeholder="Enter assett name..."
              type="text"
              class="w-full outline-none rounded-sm border px-4 py-2"
            />
          </div>
        </div>
        <div
          class="w-full aspect-[1/0.8] flex items-center justify-center rounded-md border border-dashed"
        >
          <p class="text-gray-500">Attachment</p>
        </div>
      </div>
      <div class="flex flex-col gap-2">
        <label class="px-1">Description</label>
        <textarea
          v-model="description"
          placeholder="Describe problem..."
          class="w-full h-50 outline-none rounded-sm border px-4 py-2"
        >
        </textarea>
      </div>
      <button
        type="submit"
        class="w-full my-6 bg-blue-900 py-2 rounded-lg text-lg text-white font-bold cursor-pointer transition-colors duration-150 hover:bg-blue-600"
      >
        Submit Ticket
      </button>
    </form>
    <div class="w-full hidden lg:flex flex-col justify-center">
      <img
        src="/images/helpdeskflipped.jpg"
        alt="helpdesk recipient"
        class="items-center justify-center opacity-70"
      />
    </div>
  </div>
</template>
