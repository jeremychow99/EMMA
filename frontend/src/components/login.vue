<template>
  <div class="d-flex align-center justify-center" style="height: 100vh">
    <v-sheet width="400" class="mx-auto">
      <button @click="test">test</button>
      <v-form fast-fail @submit.prevent="login">
        <v-text-field v-model="username" label="User Name"></v-text-field>
        <v-text-field v-model="password" label="Password"></v-text-field>
        <a href="#" class="text-body-2 font-weight-regular">Forgot Password?</a>

        <v-btn type="submit" color="primary" block class="mt-2">Sign in</v-btn>
      </v-form>
      <div class="mt-2">
        <p class="text-body-2">
          Don't have an account? <a href="#">Sign Up</a>
        </p>
      </div>
    </v-sheet>
  </div>
</template>
<script>
import axios from 'axios'
export default {
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    test() {
      axios
        .get(
          "http://localhost:8000/api/v1/equipment",
          {
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "*",
            },
          }
        )
        .then((response) => {
          console.log(response);

          }
        )
        .catch((err) => {
          console.log("ERROR" + err.message);
        });
    },
    login() {
      // if auth store jwt, go to main screen
      console.log("login invoked");
      console.log(this.password);
      axios
        .post(
          "http://localhost:3001/login",
          {
            name: this.username,
            password: this.password,
          },
          {
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "*",
            },
          }
        )
        .then((response) => {
          console.log(response);
          if (response.status == 200) {
            localStorage.token = response.data.token;
            console.log(localStorage.token);
            this.$router.push({ name: "main" });
          }
        })
        .catch((err) => {
          console.log("ERROR" + err.message);
        });
    },
  },
};
</script>
