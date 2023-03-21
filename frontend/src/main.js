import { createApp } from 'vue'
import router from "./router";
import App from './App.vue'

// import axios from 'redaxios';

// axios.defaults.baseURL ="http://localhost:8080"


// app.axios = axios;
// app.$http = axios;
// app.config.globalProperties.axios = axios;
// app.config.globalProperties.$http = axios;

// //create an axios instance in order to use it globally with same config
// const instance = axios.create({
//   //  baseURL: process.env.VUE_APP_API_URL,
//   withCredentials: false,
//   headers: {
//     Accept: 'application/json',
//     'Content-Type': 'application/json',
//   },

// })

// app.config.globalProperties.axios = instance;

// Vuetify
import 'vuetify/styles'
import { createVuetify } from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'

const vuetify = createVuetify({
  components,
  directives,
})

createApp(App).use(vuetify).use(router).mount('#app')
