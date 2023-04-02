<template>
  <div class="d-flex align-center justify-center" style="height: 100vh">
    <v-container>
      <v-sheet width="800" class="mx-auto rounded-xl" color="teal-lighten-3" >
        <v-row class="align-center justify-center">
          <img
            src="/title.png"
            style="width: 60%; height: auto; margin-bottom: 100px"
          />
        </v-row>
        <v-row class="justify-center">
          <div>
            <v-form fast-fail @submit.prevent="login" style="width: 500px;">
              <v-text-field v-model="username" label="User Name" style="background-color: white;" ></v-text-field>
              <v-text-field type="password" v-model="password" label="Password" style="background-color: white;"></v-text-field>

              <v-btn type="submit" block class="mt-2">Sign in</v-btn>
            </v-form>
            <div class="my-2">
              <p class="text-body-2">
                Don't have an account? <a href="#">Sign Up</a>
              </p>
            </div>
          </div>
        </v-row>
      </v-sheet>
    </v-container>
  </div>
</template>
<script>
import axios from "axios";
export default {
  data() {
    return {
      username: "",
      password: "",
    };
  },
  methods: {
    reset() {
      localStorage.clear();
      console.log(localStorage.token);
    },
    login() {
      console.log("login invoked");
      console.log(this.password);
      axios
        .post(
          "http://localhost:8000/api/v1/login",
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
            localStorage.userId = response.data.user.userId
            localStorage.token = response.data.token;
            localStorage.role = response.data.user.role
            console.log(localStorage.token);
            console.log(localStorage.userId);
            console.log(localStorage.role)

            if (localStorage.role === "ADMIN") {
              this.$router.push({ name: "dashboard" });
            } else {
              this.$router.push({ name: "techDashboard" });
            }
            
          }
        })
        .catch((err) => {
          console.log("ERROR" + err.message);
        });
    },
  },
  created() {
    console.log(localStorage.token);
  },
};
</script>
