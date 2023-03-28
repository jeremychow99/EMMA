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
          v-for="eq in equipment_arr"
          :key="eq._id"
        >
          <td>{{ eq._id }}</td>
          <td>{{ eq.equipment_name }}</td>
          <td>{{ eq.equipment_location }}</td>
          <td>{{ eq.equipment_status }}</td>
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
                <td>{{eqOperational}}</td>
              </tr>
              <tr>
                <td class="font-weight-bold text-amber-darken-1">Warning</td>
                <td>{{eqWarning}}</td>
              </tr>
              <tr>
                <td class="font-weight-bold text-red-darken-4">Down</td>
                <td>{{eqDown}}</td>
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
// import { equipmentURL } from '../../api'
// import axios from "axios";

export default {
  components: {
      navbar
  },

  data() {
      return {
        // values to be calculated
        equipment_arr: [],
        eqNoMaintain: 0,
        eqMaintain: 0,

        // test values
        // eqOperational: 87,
        // eqWarning: 10,
        // eqDown: 3,

        showReport: false,

        //  test values
        // desserts: [
        //   {
        //     name: 'Frozen Yogurt',
        //     calories: 159,
        //   },
        //   {
        //     name: 'Ice cream sandwich',
        //     calories: 237,
        //   },
        //   {
        //     name: 'Eclair',
        //     calories: 262,
        //   },
        //   {
        //     name: 'Cupcake',
        //     calories: 305,
        //   },
        //   {
        //     name: 'Gingerbread',
        //     calories: 356,
        //   },
        //   {
        //     name: 'Jelly bean',
        //     calories: 375,
        //   },
        //   {
        //     name: 'Lollipop',
        //     calories: 392,
        //   },
        //   {
        //     name: 'Honeycomb',
        //     calories: 408,
        //   },
        //   {
        //     name: 'Donut',
        //     calories: 452,
        //   },
        //   {
        //     name: 'KitKat',
        //     calories: 518,
        //   },
        // ],
      }
  },

  // have not tested yet!
  methods: {
    updateEquipmentStatus() {
      for (equipment in this.equipment_arr) {
        if (equipment.equipment_status == "No maintenance required") {
          this.eqNoMaintain++
        }
        else {
          this.eqMaintain++
        }
      }
    },

  },

  async mounted() {
        let eqp_result = await axios.get(equipmentURL)

        this.equipment_arr = eqp_result.data.data.equipment
    }

  
};

</script>
