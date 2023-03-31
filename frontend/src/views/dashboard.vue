<template>
  <navbar></navbar>

  <!-- Header -->
  <v-container>
    <div class="text-h4 text-center font-weight-black my-5">Dashboard</div>
  </v-container>

  <!-- 1st Row of Content -->
  <v-container>
    <div class="text-h6">Equipment Overview</div>
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

  <!-- 2nd Row of Content -->
  <v-container fluid>
    <v-divider class="mb-5" :thickness="3"></v-divider>
    <v-row dense>
      <v-col cols="3">
        <!-- Equipment Status Card -->
        <v-card
          title="Equipment Status"
          subtitle= "Proportion of Operational Equipment"
          text=""
          width="300px"
          class="mx-2"
          variant="tonal"
        >

          <div class="px-5 py-2">
            <v-progress-circular
              color="#3F75FC"
              :model-value="eqpPercentOperational"
              :size="128"
              :width="20"
              class="ml-16">

              <p class="font-weight-black text-h5">{{ eqpPercentOperational }}%</p>
            </v-progress-circular>
          </div>

          <v-card-actions class="pt-5">
            <v-btn color="#3F75FC" class="font-weight-bold"
              @click="showReport = !showReport">
              View Full Report
            </v-btn>
          </v-card-actions>

          <!-- Equipment Status Card - Full Report -->
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

      </v-col>

      <v-col cols="9">
      <!-- Maintenance Table -->
      <div v-if="upcomingExists">
        <div class="text-h6 mx-3">Upcoming Maintenance</div>
            <v-table
            fixed-header
            height="300px"
            >
              <thead>
                <tr>
                  <th class="text-left">
                    Maintenance ID
                  </th>
                  <th class="text-left">
                    Equipment Name
                  </th>
                  <th class="text-left">
                    Scheduled Date
                  </th>
                  <th class="text-left">
                    Status
                  </th>
                  <!-- <th class="text-left">
                    Technician ID
                  </th> -->
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="mtn in upcomingMtn_arr"
                  :key="mtn['_id']"
                >
                  <td>{{ mtn._id }}</td>
                  <td>{{ mtn.equipment.equipment_name }}</td>
                  <td>{{ mtn.schedule_date }}</td>
                  <td>{{ mtn.status }}</td>
                  <!-- <td>{{ mtn.technician_id }}</td> -->
                </tr>
              </tbody>
            </v-table>
        </div>
        <!-- If no upcoming maintenance -->
        <div v-else>
          <p class="text-h4 font-weight-bold text-red-darken-3 px-10 py-16">No Upcoming Maintenance!</p>
        </div>
      </v-col>
    </v-row>
  </v-container>

  <!-- Maintenance History Chart -->
  <v-container>
    <v-divider class="mb-5" :thickness="3"></v-divider>
    <div class="text-h6">Maintenance History</div>
    <Bar
      v-if="chartLoaded"
      id="my-chart-id"
      :options="chartOptions"
      :data="chartData"
    />
  </v-container>
  
</template>



<script>
import navbar from "../components/navbar.vue";
import { equipmentURL, maintenanceURL } from '../../api'
import axios from "axios";
import { Bar } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale, scales } from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)


export default {
  components: {
      navbar,
      Bar
  },

  data() {
      return {
        equipment_arr: [],
        eqpNoMaintain: 0,
        eqpMaintain: 0,
        eqpPercentOperational: 0,
        showReport: false,

        maintenance_arr: [],
        upcomingMtn_arr: [],
        upcomingExists: false,

        chartLoaded: false,
        chartData: {
          labels: [],
          datasets: [ { data: [] } ]
        },
        chartOptions: {
          responsive: true,
          plugins: {
            legend: {
              display: false,
            }
          },
          scales: {
            y: {
              title: {
                text: 'Number of Maintenances',
                display: true,
              },
            },
            x: {
              title: {
                text: 'Date',
                // display: true,
              }
            }
          }
        }
      }
  },

  methods: {
    // Converts maintenance date format from DB to Time thing
    convertMtnDate(mtn_date) {
      var mtnDate = mtn_date.replace(/-/g, '/')
      // var mtnDate = mtnDate.slice(0,-3)
      // console.log(mtnDate)
      // var changedDate = mtnDate.replace(/(..)\/(..)\/(....) (..):(..)/, '$3-$2-$1 $4:$5')
      // var mtnTime = new Date(changedDate)
      var mtnDate = new Date(mtnDate)
      var mtnTime = mtnDate.getTime()
      // console.log(mtnTime)
      return mtnTime
    },

    updateEquipmentStatus() {
      // console.log(this.equipment_arr)

      this.equipment_arr.forEach((eqp) => {
        // console.log(eqp.equipment_status)
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
  
  checkUpcomingMaintenance() {
    // console.log(this.maintenance_arr)

    var todayDate = new Date();
    var todayTime = todayDate.getTime()
    // console.log(todayDate.getTime());

    // Loop through each maintenance record, check if scheduled date has already passed
    this.maintenance_arr.forEach((mtn) => {
      if (mtn.schedule_date) {
        // console.log(mtn.schedule_date)
        var mtnTime = this.convertMtnDate(mtn.schedule_date)
        if (mtnTime >= todayTime) {
          this.upcomingExists = true
          this.upcomingMtn_arr.push(mtn)
        }
      }
    })
  },

  populateBarChart() {
    var numMtnLast = 0
    var numMtnCurrent = 0
    var numMtnNext = 0

    // function to get First and Last days of each week
    Date.prototype.GetFirstDayOfWeek = function() {
    return (new Date(this.setDate(this.getDate() - this.getDay()+ (this.getDay() == 0 ? -6:1) )));
    }
    Date.prototype.GetLastDayOfWeek = function() {
        return (new Date(this.setDate(this.getDate() - this.getDay() +7)));
    }

    var currentWeek = new Date();
    var lastWeek = new Date(currentWeek.getFullYear(), currentWeek.getMonth(), currentWeek.getDate() - 7);
    var nextWeek = new Date(currentWeek.getFullYear(), currentWeek.getMonth(), currentWeek.getDate() + 7);

    // Get time for start and end of each week
    var currentStartTime = currentWeek.GetFirstDayOfWeek().getTime()
    var currentEndTime = currentWeek.GetLastDayOfWeek().getTime()
    // console.log(`current week start time: ${currentStartTime}`)
    // console.log(`current week end time: ${currentEndTime}`)

    var lastStartTime = lastWeek.GetFirstDayOfWeek().getTime()
    var lastEndTime = lastWeek.GetLastDayOfWeek().getTime()
    // console.log(`last week start time: ${lastStartTime}`)
    // console.log(`last week end time: ${lastEndTime}`)

    var nextStartTime = nextWeek.GetFirstDayOfWeek().getTime()
    var nextEndTime = nextWeek.GetLastDayOfWeek().getTime()
    // console.log(`next week start time: ${nextStartTime}`)
    // console.log(`next week end time: ${nextEndTime}`)

    // console.log(currentWeek)
    // console.log(currentWeek.GetFirstDayOfWeek())
    // console.log(currentWeek.GetLastDayOfWeek())

    // Generate labels for X-Axis (weeks)
    this.chartData.labels = [
      `${lastWeek.GetFirstDayOfWeek().getDate()}/${lastWeek.GetFirstDayOfWeek().getMonth()} - ${lastWeek.GetLastDayOfWeek().getDate()}/${lastWeek.GetLastDayOfWeek().getMonth()}`,
      `${currentWeek.GetFirstDayOfWeek().getDate()}/${currentWeek.GetFirstDayOfWeek().getMonth()} - ${currentWeek.GetLastDayOfWeek().getDate()}/${currentWeek.GetLastDayOfWeek().getMonth()}`,
      `${nextWeek.GetFirstDayOfWeek().getDate()}/${nextWeek.GetFirstDayOfWeek().getMonth()} - ${nextWeek.GetLastDayOfWeek().getDate()}/${nextWeek.GetLastDayOfWeek().getMonth()}`,
      ]
    
    // Loop through maintenance records and how many maintenance in each week
    this.maintenance_arr.forEach((mtn) => {
      if (mtn.schedule_date)  {
        var mtnTime = this.convertMtnDate(mtn.schedule_date)
        // console.log(`maintenance time: ${mtnTime}`)
        if (lastStartTime <= mtnTime && mtnTime <= lastEndTime) {
          numMtnLast++
          // console.log('Last')
        }
        else if (currentStartTime <= mtnTime && mtnTime <= currentEndTime) {
          numMtnCurrent++
          // console.log('Current')

        }
        else if (nextStartTime <= mtnTime && mtnTime <= nextEndTime) {
          numMtnNext++
          // console.log('Next')

        }
      }
    })

    this.chartData.datasets = [ { data: [numMtnLast, numMtnCurrent, numMtnNext] } ]
    this.chartLoaded = true
  }
},

  async mounted() {
        let eqp_result = await axios.get(equipmentURL)
        this.equipment_arr = eqp_result.data.data.equipment

        let mtn_result = await axios.get(maintenanceURL)
        this.maintenance_arr = mtn_result.data.data.maintenance

        this.updateEquipmentStatus()
        this.checkUpcomingMaintenance()
        this.populateBarChart()
    }

  
};

</script>
