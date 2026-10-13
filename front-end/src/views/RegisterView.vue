<script setup lang="ts">
import router from '@/router'
import { ref } from 'vue'
const name = ref<string | null>(null)
const email = ref<string | null>(null)
const department = ref<string | null>(null)
const password = ref<string | null>(null)
const confirmPassword = ref<string | null>(null)

const isRegistered = ref(false)
function registerUser() {
  const missingFields: string[] = []

  if (!name.value) {
    missingFields.push('Name')
  }

  if (!email.value) {
    missingFields.push('Email')
  }

  if (!department.value) {
    missingFields.push('Department')
  }

  if (!password.value) {
    missingFields.push('Password')
  }

  if (!confirmPassword.value) {
    missingFields.push('Confirm Password')
  }

  if (missingFields.length > 0) {
    alert(`Please fill in: ${missingFields.join(', ')}`)
    return
  }

  if (password.value !== confirmPassword.value) {
    alert(`Make sure Password and Confirmed Password are the same`)
    return
  }
  isRegistered.value = true
}

function moveToLogin() {
  router.push('/login')
}

// TODO: pastikan nama memenuhi kriteria tertentu (minimal berapa huruf, tidak boleh mengandung karakter aneh-aneh spt emoji)
// TODO: pastikan password memenuhi kriteria tertentu
// TODO: kasih tanda pada kolom (misal label jadi warna merah) jika nilai input tidak sesuai harapan
// TODO: login with google
</script>
<template>
  <div class="min-h-screen flex flex-row justify-center items-center bg-blue-100">
    <div class="w-120 md:w-130 lg:w-140 px-6 py-10 flex flex-col rounded-2xl bg-white shadow-xl">
      <template v-if="!isRegistered">
        <form @submit.prevent="registerUser">
          <p class="w-full text-3xl text-center font-light">IT Helpdesk</p>
          <p class="w-full text-xl text-center font-bold my-6">Register</p>
          <div class="w-full flex flex-col items-start gap-4">
            <div class="flex flex-col gap-2 w-full">
              <label>Name</label>
              <div class="w-full rounded-sm border px-4 py-2">
                <input
                  v-model="name"
                  placeholder="Enter your name..."
                  type="text"
                  class="w-full outline-none"
                />
              </div>
            </div>

            <div class="flex flex-col gap-2 w-full">
              <label>Email</label>
              <div class="w-full rounded-sm border px-4 py-2">
                <input
                  v-model="email"
                  placeholder="Enter your email..."
                  type="email"
                  class="w-full outline-none"
                />
              </div>
            </div>

            <div class="flex flex-col gap-2 w-full">
              <label>Department</label>

              <select class="w-full rounded-sm border px-4 py-2" v-model="department">
                <option value="business-development">Business Development</option>
                <option value="customer-service">Customer Service</option>
                <option value="finance-accounting">Finance & Accounting</option>
                <option value="human-resources">Human Resources (HR)</option>
                <option value="it">IT</option>
                <option value="legal-compliance">Legal & Compliance</option>
                <option value="logistics">Logistics</option>
                <option value="management">Management</option>
                <option value="marketing">Marketing</option>
                <option value="operations">Operations</option>
                <option value="procurement">Procurement</option>
                <option value="sales">Sales</option>
              </select>
            </div>

            <div class="flex flex-col gap-2 w-full">
              <label>Password</label>
              <div class="w-full rounded-sm border px-4 py-2">
                <input
                  v-model="password"
                  placeholder="Enter new password..."
                  type="password"
                  class="w-full outline-none"
                />
              </div>
            </div>

            <div class="flex flex-col gap-2 w-full">
              <label>Confirm Password</label>
              <div class="w-full rounded-sm border px-4 py-2">
                <input
                  v-model="confirmPassword"
                  placeholder="Confirm password..."
                  type="password"
                  class="w-full outline-none"
                />
              </div>
            </div>

            <button
              type="submit"
              class="w-full my-6 bg-blue-900 py-2 rounded-lg text-lg text-white font-bold cursor-pointer transition-colors duration-150 hover:bg-blue-600"
            >
              Sign up
            </button>
          </div>
        </form>
        <div class="w-full justify-center flex flex-row gap-2">
          <p>Already have an account?</p>
          <button class="font-semibold text-blue-500 cursor-pointer" @click="moveToLogin">
            Log in
          </button>
        </div>
      </template>
      <div v-else class="flex flex-col items-center gap-4 text-center">
        <!-- TODO: ganti v-else dengan v-if akun telah berhasil dibuat -->
        <p class="text-2xl font-bold">Account created successfully!</p>
        <p class="text-gray-600">Please log in to continue.</p>

        <button
          type="button"
          class="w-full rounded-lg bg-blue-900 py-2 font-bold text-white transition-colors hover:bg-blue-600"
          @click="moveToLogin"
        >
          Go to Log In page
        </button>
      </div>
    </div>
  </div>
</template>
