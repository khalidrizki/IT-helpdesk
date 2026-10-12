<script setup lang="ts">
import router from '@/router'
import { ref } from 'vue'

const email = ref<string | null>(null)
const password = ref<string | null>(null)

function login() {
  const missingFields: string[] = []

  if (!email.value) {
    missingFields.push('Email')
  }

  if (!password.value) {
    missingFields.push('Password')
  }
  if (missingFields.length > 0) {
    alert(`Please fill in: ${missingFields.join(', ')}`)
    return
  }
  router.push('/dashboard')
}

function moveToForgotPassword() {
  router.push('/forgot-password')
}
function moveToSignUp() {
  router.push('/register')
}
</script>

<template>
  <div class="h-screen flex flex-row bg-blue-100">
    <div class="w-160 px-16 py-10 bg-white rounded-r-2xl shadow-xl">
      <form @submit.prevent="login">
        <p class="text-3xl font-light my-6">IT Helpdesk</p>
        <p class="text-xl font-bold mt-16 mb-6">Login</p>
        <div class="flex flex-col gap-4">
          <div class="flex flex-col gap-2">
            <p>Email</p>
            <div class="rounded-sm border px-4 py-2">
              <input
                v-model="email"
                placeholder="Enter your email..."
                type="email"
                class="w-full outline-none"
              />
            </div>
          </div>
          <div class="flex flex-col gap-2">
            <p>Password</p>
            <div class="rounded-sm border px-4 py-2">
              <input
                v-model="password"
                placeholder="Enter your password..."
                type="password"
                class="w-full outline-none"
              />
            </div>
          </div>
          <button
            type="button"
            class="text-sm font-semibold text-blue-500 self-start cursor-pointer"
            @click="moveToForgotPassword"
          >
            Forgot Password
          </button>
          <button
            type="submit"
            class="my-6 bg-blue-900 py-2 rounded-lg text-lg text-white font-bold cursor-pointer transition-colors duration-150 hover:bg-blue-600"
          >
            Login
          </button>
        </div>
      </form>
      <div class="flex flex-row justify-center gap-2">
        <p>Don't have an account?</p>
        <button class="font-semibold text-blue-500 cursor-pointer" @click="moveToSignUp">
          Sign Up
        </button>
      </div>
    </div>
  </div>
</template>
