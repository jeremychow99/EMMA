<template>
  <navbar></navbar>

  <!-- Header -->
  <v-container>
    <div class="text-h4 text-center font-weight-black my-5">Dashboard</div>
  </v-container>

  <!-- 1st Row of Content -->
  <v-container>
    <div v-if="upcomingExists">
      <div class="text-h6 mx-3">Upcoming Maintenance</div>
      <v-table fixed-header height="320px">
        <thead>
          <tr>
            <th class="text-left">Maintenance ID</th>
            <th class="text-left">Equipment Name</th>
            <th class="text-left">Scheduled Date</th>
            <th class="text-left">Status</th>
            <!-- <th class="text-left">
                    Technician ID
                  </th> -->
          </tr>
        </thead>
        <tbody>
          <tr v-for="mtn in upcomingMtn_arr" :key="mtn['_id']">
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
      <p class="text-h4 font-weight-bold text-red-darken-3 px-10 py-16 text-center">
        No Upcoming Maintenance!
      </p>
    </div>
  </v-container>

  <!-- 2nd Row of Content -->
  <v-container fluid>
    <v-divider class="mb-5" :thickness="3"></v-divider>
    <v-row dense>
      <v-col cols="3">
        <!-- Equipment Status Card -->
        <v-card
          title="Equipment Status"
          subtitle="Proportion of Operational Equipment"
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
              class="ml-16"
            >
              <p class="font-weight-black text-h5">
                {{ eqpPercentOperational }}%
              </p>
            </v-progress-circular>
          </div>

          <v-card-actions class="pt-5">
            <v-btn
              color="#3F75FC"
              class="font-weight-bold"
              @click="showReport = !showReport"
            >
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
                      <th class="text-left">Status</th>
                      <th class="text-left">Number of Equipment</th>
                    </tr>
                  </thead>
                  <tbody>
                    <tr>
                      <td class="font-weight-bold text-green-darken-4">
                        Operational
                      </td>
                      <td>{{ eqpNoMaintain }}</td>
                    </tr>
                    <tr>
                      <td class="font-weight-bold text-red-darken-4">Down</td>
                      <td>{{ eqpMaintain }}</td>
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
        <div class="text-h6">Equipment Overview</div>
        <v-table fixed-header height="300px">
          <thead>
            <tr>
              <th class="text-left">ID</th>
              <th class="text-left">Name</th>
              <th class="text-left">Location</th>
              <th class="text-left">Status</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="eqp in equipment_arr" :key="eqp['_id']">
              <td>{{ eqp._id }}</td>
              <td>{{ eqp.equipment_name }}</td>
              <td>{{ eqp.equipment_location }}</td>
              <td>{{ eqp.equipment_status }}</td>
            </tr>
          </tbody>
        </v-table>
      </v-col>
    </v-row>
    <v-divider class="mt-5" :thickness="3"></v-divider>
  </v-container>

  <!-- Maintenance History Chart -->
  <v-container>
    <div class="text-h6">Maintenance Frequency</div>
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
import { equipmentURL, maintenanceURL } from "../../api";
import axios from "axios";
import { Bar } from "vue-chartjs";
import {
  Chart as ChartJS,
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale,
  scales,
  elements,
} from "chart.js";

ChartJS.register(
  Title,
  Tooltip,
  Legend,
  BarElement,
  CategoryScale,
  LinearScale
);

export default {
  components: {
    navbar,
    Bar,
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
        datasets: [
          {
            data: [],
          },
        ],
      },
      chartOptions: {
        responsive: true,
        plugins: {
          legend: {
            display: false,
          },
        },
        scales: {
          y: {
            title: {
              text: "Number of Maintenances",
              display: true,
            },
          },
          x: {
            title: {
              text: "Date",
              // display: true,
            },
          },
        },
        elements: {
          bar: {
            backgroundColor: [
              "rgba(0,105,92,0.4)",
              "rgba(0,105,92,0.4)",
              "rgba(0,105,92,0.8)",
              "rgba(0,105,92,0.4)",
              "rgba(0,105,92,0.4)",
            ],
            borderColor: [
              "rgba(0,105,92,1)",
              "rgba(0,105,92,1)",
              "rgba(0,105,92,1)",
              "rgba(0,105,92,1)",
              "rgba(0,105,92,1)",
            ],
            borderWidth: 1,
          },
        },
      },
    };
  },

  methods: {
    // Converts maintenance date format from DB to Time thing
    convertMtnDate(mtn_date) {
      var mtnDate = mtn_date.replace(/-/g, "/");
      // var mtnDate = mtnDate.slice(0,-3)
      // console.log(mtnDate)
      // var changedDate = mtnDate.replace(/(..)\/(..)\/(....) (..):(..)/, '$3-$2-$1 $4:$5')
      // var mtnTime = new Date(changedDate)
      var mtnDate = new Date(mtnDate);
      var mtnTime = mtnDate.getTime();
      // console.log(mtnTime)
      return mtnTime;
    },

    updateEquipmentStatus() {
      // console.log(this.equipment_arr)

      this.equipment_arr.forEach((eqp) => {
        console.log(eqp.equipment_status);
        if (
          eqp.equipment_status == "Undergoing Maintenance" ||
          eqp.equipment_status == "Down"
        ) {
          this.eqpMaintain++;
        } else {
          this.eqpNoMaintain++;
        }
      });

      this.eqpPercentOperational =
        (this.eqpNoMaintain / (this.eqpMaintain + this.eqpNoMaintain)) * 100;
      this.eqpPercentOperational = this.eqpPercentOperational.toFixed(0);
    },

    checkUpcomingMaintenance() {
      console.log(this.maintenance_arr)
      console.log("Checking Upcoming Maintenance")
      var todayDate = new Date();
      var todayTime = todayDate.getTime();
      // console.log(todayDate.getTime());

      // Loop through each maintenance record, check if scheduled date has already passed
      this.maintenance_arr.forEach((mtn) => {
        // console.log(mtn.schedule_date)
        if (mtn.schedule_date) {
          // console.log(mtn.schedule_date)
          var mtnTime = this.convertMtnDate(mtn.schedule_date);
          // console.log(mtnTime)
          if (mtnTime >= todayTime) {
            this.upcomingExists = true;
            this.upcomingMtn_arr.push(mtn);
          }
        }
      });

      this.upcomingMtn_arr.sort((a, b) => a.schedule_date > b.schedule_date ? 1 : -1);

    },

    populateBarChart() {
      var numMtnLast = 0;
      var numMtnCurrent = 0;
      var numMtnNext = 0;
      var numMtnFollowing = 0;
      var numMtnFollowingNext = 0;

      // function to get First and Last days of each week
      Date.prototype.GetFirstDayOfWeek = function () {
        return new Date(
          this.setDate(
            this.getDate() - this.getDay() + (this.getDay() == 0 ? -6 : 1)
          )
        );
      };
      Date.prototype.GetLastDayOfWeek = function () {
        return new Date(this.setDate(this.getDate() - this.getDay() + 7));
      };

      var currentWeek = new Date();
      var lastWeek = new Date(
        currentWeek.getFullYear(),
        currentWeek.getMonth(),
        currentWeek.getDate() - 7
      );
      var nextWeek = new Date(
        currentWeek.getFullYear(),
        currentWeek.getMonth(),
        currentWeek.getDate() + 7
      );
      var followingWeek = new Date(
        currentWeek.getFullYear(),
        currentWeek.getMonth(),
        currentWeek.getDate() + 14
      );
      var followingNextWeek = new Date(
        currentWeek.getFullYear(),
        currentWeek.getMonth(),
        currentWeek.getDate() + 21
      );

      // Get time for start and end of each week
      var currentStartTime = currentWeek.GetFirstDayOfWeek().getTime();
      var currentEndTime = currentWeek.GetLastDayOfWeek().getTime();
      // console.log(`current week start time: ${currentStartTime}`)
      // console.log(`current week end time: ${currentEndTime}`)

      var lastStartTime = lastWeek.GetFirstDayOfWeek().getTime();
      var lastEndTime = lastWeek.GetLastDayOfWeek().getTime();
      // console.log(`last week start time: ${lastStartTime}`)
      // console.log(`last week end time: ${lastEndTime}`)

      var nextStartTime = nextWeek.GetFirstDayOfWeek().getTime();
      var nextEndTime = nextWeek.GetLastDayOfWeek().getTime();
      // console.log(`next week start time: ${nextStartTime}`)
      // console.log(`next week end time: ${nextEndTime}`)

      var followingStartTime = followingWeek.GetFirstDayOfWeek().getTime();
      var followingEndTime = followingWeek.GetLastDayOfWeek().getTime();

      var followingNextStartTime = followingNextWeek
        .GetFirstDayOfWeek()
        .getTime();
      var followingNextEndTime = followingNextWeek.GetLastDayOfWeek().getTime();

      // console.log(currentWeek)
      // console.log(currentWeek.GetFirstDayOfWeek())
      // console.log(currentWeek.GetLastDayOfWeek())

      // Generate labels for X-Axis (weeks)
      this.chartData.labels = [
        `${lastWeek.GetFirstDayOfWeek().getDate()}/${
          lastWeek.GetFirstDayOfWeek().getMonth() + 1
        } - ${lastWeek.GetLastDayOfWeek().getDate()}/${
          lastWeek.GetLastDayOfWeek().getMonth() + 1
        }`,
        `${currentWeek.GetFirstDayOfWeek().getDate()}/${
          currentWeek.GetFirstDayOfWeek().getMonth() + 1
        } - ${currentWeek.GetLastDayOfWeek().getDate()}/${
          currentWeek.GetLastDayOfWeek().getMonth() + 1
        }`,
        `${nextWeek.GetFirstDayOfWeek().getDate()}/${
          nextWeek.GetFirstDayOfWeek().getMonth() + 1
        } - ${nextWeek.GetLastDayOfWeek().getDate()}/${
          nextWeek.GetLastDayOfWeek().getMonth() + 1
        }`,
        `${followingWeek.GetFirstDayOfWeek().getDate()}/${
          followingWeek.GetFirstDayOfWeek().getMonth() + 1
        } - ${followingWeek.GetLastDayOfWeek().getDate()}/${
          followingWeek.GetLastDayOfWeek().getMonth() + 1
        }`,
        `${followingNextWeek.GetFirstDayOfWeek().getDate()}/${
          followingNextWeek.GetFirstDayOfWeek().getMonth() + 1
        } - ${followingNextWeek.GetLastDayOfWeek().getDate()}/${
          followingNextWeek.GetLastDayOfWeek().getMonth() + 1
        }`,
      ];

      // Loop through maintenance records and how many maintenance in each week
      this.maintenance_arr.forEach((mtn) => {
        if (mtn.schedule_date) {
          var mtnTime = this.convertMtnDate(mtn.schedule_date);
          // console.log(`maintenance time: ${mtnTime}`)
          if (lastStartTime <= mtnTime && mtnTime <= lastEndTime) {
            numMtnLast++;
            // console.log('Last')
          } else if (currentStartTime <= mtnTime && mtnTime <= currentEndTime) {
            numMtnCurrent++;
            // console.log('Current')
          } else if (nextStartTime <= mtnTime && mtnTime <= nextEndTime) {
            numMtnNext++;
            // console.log('Next')
          } else if (
            followingStartTime <= mtnTime &&
            mtnTime <= followingEndTime
          ) {
            numMtnFollowing++;
            // console.log('Next')
          } else if (
            followingNextStartTime <= mtnTime &&
            mtnTime <= followingNextEndTime
          ) {
            numMtnFollowingNext++;
            // console.log('Next')
          }
        }
      });

      this.chartData.datasets = [
        {
          data: [
            numMtnLast,
            numMtnCurrent,
            numMtnNext,
            numMtnFollowing,
            numMtnFollowingNext,
          ],
        },
      ];
      this.chartLoaded = true;
    },
  },

  async mounted() {
    console.log("mounting");
    let eqp_result = await axios.get(equipmentURL, {
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
        Authorization: "Bearer " + localStorage.token,
      },
    });
    this.equipment_arr = eqp_result.data.data.equipment;

    let mtn_result = await axios
      .get(maintenanceURL, {
        headers: {
          "Content-Type": "application/json",
          "Access-Control-Allow-Origin": "*",
          Authorization: "Bearer " + localStorage.token,
        },
      })
      .then((response) => {
        if (response) console.log(response.data.data.maintenance);
        this.maintenance_arr = response.data.data.maintenance
      })
      .catch((error) => {
        if (error.response) {
          console.log(error.response.data);
          console.log(error.response.status);
        }
      });

    this.updateEquipmentStatus();
    this.checkUpcomingMaintenance();
    this.populateBarChart();
  },
};
</script>
