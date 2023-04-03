<template>
  <navbar></navbar>

  <v-container>
    <v-row>
      <v-col cols="4">
        <h2 class="text-start mt-3">Schedule</h2>
        <h2 class="text-start mb-3">Maintenance</h2>
      </v-col>
      <v-col cols="6" style="background-color: lavender">
        <v-row class="mt-0 mx-2">
          <v-col cols="5" style="padding: 8px 16px">
            <p class="text-right font-weight-regular">EquipmentID:</p>
          </v-col>
          <v-col cols="7" style="padding: 8px 16px">
            <p>{{ equipment_id }}</p>
          </v-col>
        </v-row>
        <v-row class="mt-0 mx-2">
          <v-col cols="5" style="padding: 8px 16px">
            <p class="text-right font-weight-regular">Parts:</p>
          </v-col>
          <v-col cols="7" style="padding: 8px 16px">
            <p v-if="Object.keys(part_list).length != 0">Reserved</p>
            <p v-else>No Parts Reserved</p>
          </v-col>
        </v-row>
        <v-row class="mt-0 mx-2">
          <v-col cols="5" style="padding: 8px 16px 8px 8px">
            <p class="text-right font-weight-regular">Maintenance Datetime:</p>
          </v-col>
          <v-col cols="7" style="padding: 8px 16px">
            <input
              type="date"
              id="maintenance-datetime"
              name="maintenance-datetime"
              v-model="maintenance_date"
              @change="checkBusyTechnicians"
            />
          </v-col>
        </v-row>
        <v-row class="mt-0 mx-2">
          <v-col cols="5" style="padding: 8px 16px 8px 8px">
            <p class="text-right font-weight-regular">Technician:</p>
          </v-col>
          <v-col cols="7" style="padding: 8px 16px">
            <p v-if="availableTechnicians.length == 0" style="color: red">
              No Available Technicians
            </p>
            <select
              v-else
              id="technician-select"
              name="technician-select"
              v-model="selected_technician"
              :disabled="!maintenance_date"
              style="
                width: 200px;
                border: solid 1px gray;
                border-radius: 5px;
                text-align: center;
              "
            >
              <option value="" disabled selected>Select your option</option>
              <option v-for="tech in availableTechnicians" :value="tech['_id']">
                {{ tech.name }}
              </option>
            </select>
          </v-col>
        </v-row>
      </v-col>
      <v-col
        cols="2"
        class="position-relative"
        style="background-color: lavender"
      >
        <v-btn
          class="btn btn-success position-absolute"
          style="right: 5px; bottom: 5px"
          @click="ScheduleMaintenance"
          :disabled="!formValid"
          >Submit</v-btn
        >
      </v-col>
    </v-row>

    <!-- Testing Purpose -->
    <!-- <div>
      {{ availableTechnicians }}
    </div> -->

    <v-tabs class="mx-auto mt-3" align-tabs="center" v-model="tab">
      <v-tab value="equipment"> Equipment </v-tab>
      <v-tab value="parts"> Parts </v-tab>
    </v-tabs>

    <v-window v-model="tab">
      <v-window-item key="equipment" value="equipment">
        <v-container
          class="rounded-lg"
          style="background-color: lightblue; padding: 5px"
        >
          <v-text-field
            :loading="loading"
            density="compact"
            variant="solo"
            label="Search equipment name"
            single-line
            hide-details
            @click:append-inner="onClick"
            class="my-2"
            style="margin: auto; width: 50%"
          ></v-text-field>

          <v-table fixed-header height="400px">
            <thead>
              <tr>
                <th class="text-center">Name</th>
                <th class="text-center">Status</th>
                <th class="text-center">Select</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="eqp in equipment_arr" :key="eqp['_id']">
                <td class="text-center">{{ eqp.equipment_name }}</td>
                <td class="text-center">{{ eqp.equipment_status }}</td>
                <td class="text-center">
                  <input
                    type="radio"
                    :id="eqp['_id']"
                    :value="eqp['_id']"
                    name="eqp_radios"
                    v-model="equipment_id"
                  />
                </td>
              </tr>
            </tbody>
          </v-table>
        </v-container>
      </v-window-item>

      <v-window-item key="parts" value="parts">
        <v-container
          class="rounded-lg"
          style="background-color: lightgray; padding: 5px"
        >
          <v-text-field
            :loading="loading"
            density="compact"
            variant="solo"
            label="Search part name"
            single-line
            hide-details
            @click:append-inner="onClick"
            class="my-2"
            style="margin: auto; width: 50%"
          ></v-text-field>

          <v-table fixed-header height="400px">
            <thead>
              <tr>
                <th class="text-center">Part Name</th>
                <th class="text-center">Quantity</th>
                <th class="text-center" width="20%">Reserve Qty</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="part in inventory_arr">
                <td class="text-center">{{ part["PartName"] }}</td>
                <td class="text-center">{{ part["Qty"] }}</td>
                <td class="text-center">
                  <!-- <v-text-field :id="part['_id']" type="number" width="50" value="0" min="0"></v-text-field> -->
                  <input
                    type="number"
                    :id="part['_id']"
                    :name="part['PartName']"
                    min="0"
                    value="0"
                    style="border: 1px solid gray; width: 50%"
                    class="text-center"
                    @change="UpdatePartList($event.currentTarget)"
                  />
                </td>
              </tr>
            </tbody>
          </v-table>
        </v-container>
      </v-window-item>
    </v-window>
  </v-container>
</template>

<script>
import navbar from "../components/navbar.vue";
import { equipmentURL, inventoryURL, maintenanceURL, userURL } from "../../api";
import axios from "axios";

export default {
  components: {
    navbar,
  },

  data() {
    return {
      tab: null,
      date: null,
      equipment_arr: [],
      inventory_arr: [],
      technician_arr: [],

      busytech_arr: [],

      // Form Inputs
      equipment_id: "",
      part_list: [],
      maintenance_date: "",
      selected_technician: "",
    };
  },

  computed: {
    formValid() {
      return (
        this.equipment_id != "" &&
        this.maintenance_date != "" &&
        this.selected_technician != ""
      );
    },

    availableTechnicians() {
      let availabletech_arr = [];

      for (let technician of this.technician_arr) {
        if (!this.busytech_arr.includes(technician["_id"])) {
          availabletech_arr.push(technician);
        }
      }

      return availabletech_arr;
    },
  },

  methods: {
    UpdatePartList(e) {
      // console.log(e)
      // console.log(e.id)
      // console.log(e.value)
      let part_id = e.id;
      let reserve_qty = e.value;
      let part_name = e.name;

      let partObj = {
        PartName: part_name,
        Qty: reserve_qty,
        _id: part_id,
      };

      let temp_part_list = this.part_list;

      let part_idx = temp_part_list.findIndex(
        (part) => part["_id"] === part_id
      );

      console.log(part_idx);

      // Remove part from part list
      if (reserve_qty === "0") {
        temp_part_list = temp_part_list.filter(
          (part) => part["_id"] != part_id
        );
      } else if (part_idx == -1) {
        // console.log("Added Parts")
        temp_part_list.push(partObj);
      } else {
        temp_part_list[part_idx]["Qty"] = reserve_qty;
      }

      this.part_list = temp_part_list;
    },

    ScheduleMaintenance() {
      // console.log(this.equipment_id)
      // console.log(this.part_list)
      console.log(this.maintenance_date);

      // var datetime_arr = this.maintenance_date.split("T");
      // var date = datetime_arr[0]
      // var time = datetime_arr[1]
      // var date_arr = date.split("-")

      // var month = date_arr[1]
      // var day = date_arr[2]
      // var year = date_arr[0]

      // // console.log(month)
      // // console.log(day)
      // // console.log(year)

      // var formatted_datetime = day + '-' + month + '-' + year + " " + time + ":00"
      // console.log(formatted_datetime)

      let eqpObj = this.equipment_arr.find((eqp) => {
        return eqp["_id"] == this.equipment_id;
      });
      let techObj = this.technician_arr.find((tech) => {
        return tech["_id"] == this.selected_technician;
      });

      // console.log(eqpObj)
      // console.log(techObj)

      let tempEqp = {
        equipment_id: eqpObj["_id"],
        equipment_name: eqpObj["equipment_name"],
        equipment_location: eqpObj["equipment_location"],
      };

      let tempTech = {
        technician_id: techObj["_id"],
        name: techObj["name"],
        phone: techObj["phone"],
      };

      if (
        this.equipment_id != "" &&
        this.maintenance_date != "" &&
        this.selected_technician != ""
      ) {
        console.log("Submitted schedule maintenance request");
        let data = {
          equipment: tempEqp,
          schedule_date: this.maintenance_date,
          partlist: this.part_list,
          technician: tempTech,
          // Need to include creator_id from sessionStorage
        };

        // console.log(this.part_list)
        console.log(data);

        axios
          .post(`http://127.0.0.1:8000/api/v1/schedule_maintenance`, data, {
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "*",
              Authorization: "Bearer " + localStorage.token,
            },
          })
          .then((response) => {
            console.log(response.data);

            // Reset Inputs
            this.equipment_id = "";
            this.part_list = [];
            this.maintenance_date = "";
            this.selected_technician = "";

            this.$router.go();
          })
          .catch((error) => {});
      }
    },

    async checkBusyTechnicians() {
      // console.log(this.maintenance_date)
      await axios
        .get(`${maintenanceURL}/busy_technicians/${this.maintenance_date}`, {
          headers: {
            "Content-Type": "application/json",
            "Access-Control-Allow-Origin": "*",
            Authorization: "Bearer " + localStorage.token,
          },
        })
        .then((response) => {
          console.log(response.data);

          let busy_arr = response.data.data;

          this.busytech_arr = busy_arr;
        });
    },
  },
  async mounted() {
    let eqp_result = await axios.get(equipmentURL, {
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
        Authorization: "Bearer " + localStorage.token,
      },
    });
    let inv_result = await axios.get(inventoryURL, {
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
        Authorization: "Bearer " + localStorage.token,
      },
    });
    let user_result = await axios.get(`${userURL}`, {
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
        Authorization: "Bearer " + localStorage.token,
      },
    });

    let user_arr = user_result.data.users;
    console.log(user_arr);
    let temp_user = [];

    for (let user of user_arr) {
      console.log(user);
      if (user.role === "TECHNICIAN") {
        temp_user.push(user);
      }
    }

    this.technician_arr = temp_user;
    this.equipment_arr = eqp_result.data.data.equipment;
    this.inventory_arr = inv_result.data.data.parts;
  },
};
</script>
