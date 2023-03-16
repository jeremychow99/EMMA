<script setup>

</script>

<template>
  <v-app>
    <router-view v-slot="{ Component }">
          <div><component :is="Component"/></div> 
      </router-view>
    </v-app>
</template>

