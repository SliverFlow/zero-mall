import { defineStore } from 'pinia'
import { ref } from 'vue'

export const useActiveStore = defineStore('active', () => {
  const active = ref(false)

  return {
    active
  }
})
