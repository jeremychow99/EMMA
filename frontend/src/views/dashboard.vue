<template>
  <navbar></navbar>

  <v-container>
    <div class="text-h4 font-weight-bold">Equipment Overview</div>
    <v-table
    fixed-header
    height="300px"
    >
      <thead>
        <tr>
          <th class="text-left">
            ID
          </th>
          <th class="text-left">
            Name
          </th>
          <th class="text-left">
            Location
          </th>
          <th class="text-left">
            Status
          </th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="eqp in equipment_arr"
          :key="eqp['_id']"
        >
          <td>{{ eqp._id }}</td>
          <td>{{ eqp.equipment_name }}</td>
          <td>{{ eqp.equipment_location }}</td>
          <td>{{ eqp.equipment_status }}</td>
        </tr>
      </tbody>
    </v-table>
  
  </v-container>

  <!-- Equipment Status Card -->
  
  <v-container>
    <v-divider class="mb-5"></v-divider>
    <v-card
      title="Equipment Status"
      subtitle="Proportion of Operational Equipment"
      text=""
      width="300px"
    >

      <div class="px-5 py-2">
        <v-progress-circular
          color="#3F75FC"
          :model-value="eqpPercentOperational"
          :size="128"
          :width="20">

          <p class="font-weight-black text-h5">{{ eqpPercentOperational }}%</p>
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
                <td>{{eqpNoMaintain}}</td>
              </tr>
              <tr>
                <td class="font-weight-bold text-red-darken-4">Down</td>
                <td>{{eqpMaintain}}</td>
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
import { equipmentURL } from '../../api'
import axios from "axios";

export default {
  components: {
      navbar
  },

  data() {
      return {
        // values to be calculated
        equipment_arr: [],
        eqpNoMaintain: 0,
        eqpMaintain: 0,
        eqpPercentOperational: 0,

        // test values
        eqOperational: 87,
        eqWarning: 10,
        eqDown: 3,

        showReport: false,
      }
  },

  methods: {
    updateEquipmentStatus() {
      // console.log(this.equipment_arr)

      this.equipment_arr.forEach((eqp) => {
        // console.log(eqp)
        if (eqp.equipment_status == "Undergoing Maintenance" || eqp.equipment_status == "Down") {
          this.eqpMaintain ++
        }
        else {
          this.eqpNoMaintain ++
        }
      })

      this.eqpPercentOperational = (this.eqpNoMaintain / (this.eqpMaintain + this.eqpNoMaintain)) * 100
      this.eqpPercentOperational = this.eqpPercentOperational.toFixed(0)
    },

  },

  async mounted() {
        let eqp_result = await axios.get(equipmentURL)
        this.equipment_arr = eqp_result.data.data.equipment
        this.updateEquipmentStatus()
    }

  
};

</script>
