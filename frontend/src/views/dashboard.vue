<template>
  <navbar></navbar>

  <v-container>

    <v-card
      title="Equipment Status"
      subtitle="Proportion of Operational Equipment"
      text=""
      width="300px"
    >

      <div class="px-5 py-2">
        <v-progress-circular
          color="#3F75FC"
          :model-value="eqOperational"
          :size="128"
          :width="20">

          <p class="font-weight-black text-h5">{{ eqOperational }}%</p>
        </v-progress-circular>
      </div>

      <v-card-actions class="pt-5">
        <v-btn
          @click="showReport = !showReport">
          View Full Report
        </v-btn>
      </v-card-actions>

      <v-expand-transition>
        <div v-show="showReport">
          <v-divider></v-divider>

        <v-card-text>
          <v-table>
            <thead>
              <tr>
                <th class="text-left">
                  Status
                </th>
                <th class="text-left">
                  Number of Equipment
                </th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td class="font-weight-bold text-green-darken-4">Operational</td>
                <td>87</td>
              </tr>
              <tr>
                <td class="font-weight-bold text-amber-darken-1">Warning</td>
                <td>10</td>
              </tr>
              <tr>
                <td class="font-weight-bold text-red-darken-4">Down</td>
                <td>3</td>
              </tr>
            </tbody>
          </v-table>
        </v-card-text>

        </div>
      </v-expand-transition>

    </v-card>

  </v-container>

</template>



<script>
import navbar from "../components/navbar.vue";

export default {
  components: {
      navbar
  },

  data() {
      return {
        // value to be calculated
        eqOperational: 87,
        showReport: false,
      }
  }
};

</script>
