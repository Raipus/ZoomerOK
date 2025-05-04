<script setup lang="ts">
import { useDark, useToggle } from '@vueuse/core'

const isDark = useDark({
  selector: 'body',
  attribute: 'data-theme',
  valueDark: 'dark',
  valueLight: 'light',
})
const toggleDark = useToggle(isDark)
</script>

<template>
  <button @click="toggleDark()">
    <v-fade-transition>
      <v-avatar v-if="isDark" color="#D9D9D9" size="45" style="position: absolute; z-index: 99">
        <v-icon icon="dark_mode" />
      </v-avatar>
    </v-fade-transition>
    <v-avatar color="#D9D9D9" size="45">
      <v-icon icon="light_mode" />
    </v-avatar>
  </button>
</template>
