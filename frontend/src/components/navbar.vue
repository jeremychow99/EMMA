<template>

<v-toolbar>
        <v-toolbar-title><h1>EMMA</h1></v-toolbar-title>
        <v-toolbar-items>
         <v-btn
           v-for="item in nav"
         >{{ item.text }}</v-btn>
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
  }
  }
</script>
