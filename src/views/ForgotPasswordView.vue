<script setup lang="ts">
import { ref } from 'vue'

const email = ref<string | null>(null)

function inputEmail() {
  if (!email.value) {
    alert(`Please input your email`)
  }
}
</script>
<template>
  <div class="min-h-screen flex flex-row justify-center items-center bg-blue-100">
    <div class="w-100 px-6 py-2 flex flex-col rounded-2xl bg-white shadow-xl">
      <p class="w-full text-xl text-center font-bold my-6">Forgot Password?</p>
      <p class="text-center">
        Enter your email and we'll send you a link to get back into your account
      </p>
      <form @submit.prevent="inputEmail">
        <div class="flex flex-col gap-2 w-full">
          <label class="font-bold">Email</label>
          <div class="w-full rounded-sm border px-4 py-2">
            <input
              v-model="email"
              placeholder="Enter your email..."
              type="text"
              class="w-full outline-none"
            />
          </div>
        </div>
        <button
          type="submit"
          class="w-full my-6 bg-blue-900 py-2 rounded-lg text-lg text-white font-bold cursor-pointer transition-colors duration-150 hover:bg-blue-600"
        >
          Reset Password
        </button>
      </form>
    </div>
  </div>
</template>
