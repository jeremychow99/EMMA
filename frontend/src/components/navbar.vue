<template>

<v-toolbar>
        <v-toolbar-title><h1>EMMA</h1></v-toolbar-title>
        <v-toolbar-items>
          <v-btn>Home</v-btn>
          <v-btn>Maintenance</v-btn>
          <v-btn>Inventory</v-btn>
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
  }
  // TODO: Implement Logout button
  // clear localstorage token, redirect to login page.
  }
</script>
