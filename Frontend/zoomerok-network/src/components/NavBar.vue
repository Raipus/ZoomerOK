<script setup lang="ts">
import { RouterLink } from 'vue-router'
import { useAccountStore } from '@/stores/accountStore'
import { storeToRefs } from 'pinia'
import ThemeButton from './ThemeButton.vue'

const userStore = useAccountStore()
const { isAuthenticated } = storeToRefs(userStore)

const logout = () => {
  //userStore.logout()
}
</script>

<template>
  <!-- Условие отключено на время dev -->
  <nav v-if="!isAuthenticated" className="h-2">
    <div class="navdiv">
      <div class="logodiv">
        <v-avatar color="#D9D9D9" size="45">
          <p style="font-weight: 1000; font-size: 150%">Z</p>
        </v-avatar>
      </div>
      <div class="linksdiv">
        <router-link to="/news">
          <v-avatar color="#D9D9D9" size="45">
            <v-icon alt="John" icon="newspaper"></v-icon>
          </v-avatar>
        </router-link>
        <router-link to="/my_page">
          <v-avatar color="#D9D9D9" size="45">
            <v-icon icon="home"></v-icon>
          </v-avatar>
        </router-link>
        <router-link to="/friends">
          <v-avatar color="#D9D9D9" size="45">
            <v-icon icon="people_alt"></v-icon>
          </v-avatar>
        </router-link>
      </div>
      <div class="accountdiv">
        <ThemeButton />
        <button @click="logout">
          <v-hover v-slot="{ isHovering, props }">
            <div v-bind="props">
              <v-fade-transition>
                <v-avatar
                  v-if="isHovering"
                  color="#acacac"
                  size="45"
                  style="position: absolute; z-index: 99"
                >
                  <v-icon icon="exit_to_app" />
                </v-avatar>
              </v-fade-transition>
              <v-avatar color="#D9D9D9" size="45">
                <v-icon icon="person" />
              </v-avatar>
            </div>
          </v-hover>
        </button>
      </div>
    </div>
  </nav>
</template>

<style lang="scss">
nav {
  height: 91px;
  width: 1118px;
  margin: auto;
  margin-top: 26px;
  margin-bottom: 26px;
  display: flex;
  place-items: center;
  background-color: #ff00a9;
  border-radius: 45px;
  .navdiv {
    height: 67px;
    width: 1097px;
    margin: 0 auto;
    display: grid;
    place-items: center;
    grid-template-columns: 1fr auto 1fr;
    background-color: #ffffff;
    border-radius: 45px;
    .logodiv {
      margin-left: 18px;
      justify-self: start;
    }
    .linksdiv {
      display: grid;
      justify-self: center;
      grid-auto-flow: column;
      gap: 22px;
    }
    .accountdiv {
      margin-right: 18px;
      justify-self: end;
      display: grid;
      grid-auto-flow: column;
      gap: 22px;
    }
  }
}
</style>
