<template>

<v-toolbar>
        <v-toolbar-title><h1>EMMA</h1></v-toolbar-title>
        <v-toolbar-items v-if="isAdmin">
          <router-link class="d-flex justify-center text-black" :to="{ name: 'dashboard', params: {} }" style="text-decoration:none">
            <v-btn>Dashboard</v-btn>
          </router-link>
          <router-link class="d-flex justify-center text-black" :to="{ name: 'schedule', params: {} }" style="text-decoration:none">
            <v-btn>Maintenance</v-btn>
          </router-link>
          <router-link class="d-flex justify-center text-black" :to="{ name: 'inventory', params: {} }" style="text-decoration:none">
            <v-btn>Inventory</v-btn>
          </router-link>
          <router-link class="d-flex justify-center text-black" :to="{ name: 'equipment', params: {} }" style="text-decoration:none">
            <v-btn>Equipment</v-btn>
          </router-link>
          <v-btn @click="logout">Logout</v-btn>
        </v-toolbar-items>

        <v-toolbar-items v-else>
          <router-link class="d-flex justify-center text-black" :to="{ name: 'techDashboard', params: {} }" style="text-decoration:none">
            <v-btn>Tasks</v-btn>
          </router-link>
          <v-btn @click="logout">Logout</v-btn>
        </v-toolbar-items>
      </v-toolbar>
</template>

<script>
export default {
    data () {
      return {
        nav: [
          {
            text: 'Home',
          },
          {
            text: 'Maintenance',
          },
          {
            text: 'Inventory',
          },
          {
            text: 'Logout',
          }
        ]
      }
  },
  methods: {
    logout () {
      localStorage.clear()
      this.$router.push({name: 'login'})
    }
  },
  computed: {
    isAdmin() {
      return sessionStorage.getItem('role') === "ADMIN"
    }
  }
  // TODO: Implement Logout button
  // clear localstorage token, redirect to login page.
  }
</script>
