<template>
  <navbar></navbar>

  <!-- Header -->
  <v-container>
    <div class="text-h4 text-center font-weight-black my-5">Inventory</div>
  </v-container>

  <!-- Available Parts Table-->
  <v-container>
    <div class="text-h6">All Parts Overview</div>
    <v-table fixed-header height="300px">
      <thead>
        <tr>
          <th class="text-left">ID</th>
          <th class="text-left">Part Name</th>
          <th class="text-left">Quantity</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="part in inventory_arr" :key="part['_id']">
          <td>{{ part._id }}</td>
          <td>{{ part.PartName }}</td>
          <td>{{ part.Qty }}</td>
        </tr>
      </tbody>
    </v-table>
  </v-container>
</template>

<script>
import navbar from "../components/navbar.vue";
import { inventoryURL } from "../../api";
import axios from "axios";

export default {
  components: {
    navbar,
  },

  data() {
    return {
      inventory_arr: [],
    };
  },

  async mounted() {
    let inv_result = await axios.get(inventoryURL, {
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
        Authorization: "Bearer " + localStorage.token,
      },
    });
    this.inventory_arr = inv_result.data.data.parts;
  },
};
</script>
