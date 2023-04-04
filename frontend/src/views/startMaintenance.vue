<template>
  <navbar></navbar>

  <v-container>
    <v-row>
      <v-col cols="4">          
        <v-btn style="background-color: grey; color: white; margin-right: 10px;" @click="$router.go(-1)">
        &lt;
        </v-btn>
        </v-col>
      <v-col cols="4">
        <h1 class="text-start mt-3">View Maintenance</h1>
      </v-col>
      <v-col cols="4"></v-col>
    </v-row>

    <v-container class="d-flex justify-center">
      <v-div v-if="selectedRecord.status == 'SCHEDULED'">
        <v-row>
          <v-col>
            <h2>Maintenance Record ID: {{ selectedRecord._id }}</h2>
          </v-col>
        </v-row>

        <v-row style="margin-top: 0">
          <v-col>
            <h3>Equipment ID: {{ selectedRecord.equipment.equipment_name }}</h3>
          </v-col>
        </v-row>

        <v-row style="margin-top: 0">
          <v-col>
            <h3>Current status: {{ selectedRecord.status }}</h3>
          </v-col>
        </v-row>

        <v-row style="margin-top: 0">
          <v-col>
            <h3>Scheduled for: {{ selectedRecord.schedule_date }}</h3>
          </v-col>
        </v-row>

        <v-table>
          <thead>
            <tr>
              <th>Parts</th>
              <th>Quantity Reserved</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="detail in selectedRecord.partlist">
              <td>
                {{ detail.PartName }}
              </td>
              <td>
                {{ detail.Qty }}
              </td>
            </tr>
          </tbody>
        </v-table>

        <v-row class="justify-end">
          <!-- <v-btn style="background-color: #5D5FEF; color: white; margin-right: 10px;">
                        Previous
                    </v-btn> -->
          <v-btn
            @click="update_maintenance_status"
            style="background-color: #5d5fef; color: white"
          >
            Start
          </v-btn>
        </v-row>
      </v-div>

      <v-div v-else-if="selectedRecord.status == 'STARTED'">
        <v-row>
          <v-col>
            <h3>Maintenance Record ID: {{ selectedRecord._id }}</h3>
          </v-col>
        </v-row>

        <v-row style="margin-top: 0">
          <v-col>
            <h4>Equipment ID: {{ selectedRecord.equipment.equipment_name }}</h4>
          </v-col>
        </v-row>

        <v-row style="margin-top: 0">
          <v-col>
            <h4>Current status: {{ selectedRecord.status }}</h4>
          </v-col>
        </v-row>

        <v-row style="margin-top: 0">
          <v-col>
            <h4>Scheduled for: {{ selectedRecord.schedule_date }}</h4>
          </v-col>
        </v-row>

        <v-row>
          <v-col>
            <h3>Unused Parts</h3>
            <!-- <div>
              {{ unusedPartList }}
            </div> -->
          </v-col>
        </v-row>

        <v-row>
          <v-table>
            <thead>
              <tr>
                <th class="text-left">Part Name</th>
                <th class="text-left">Qty Reserved</th>
                <th class="text-left">Qty Unused</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="parts in selectedRecord.partlist" :key="parts">
                <td>{{ parts.PartName }}</td>
                <td>
                  {{ parts.Qty }}
                </td>
                <td>
                  <input
                    type="number"
                    :id="parts['_id']"
                    :name="parts.PartName"
                    value="0"
                    placeholder="Number of unused parts"
                    style="border: 1px solid grey; text-align: center; width: 100px"
                    @change="update_unused_partlist_array($event.currentTarget)"
                  />
                </td>
              </tr>
            </tbody>
          </v-table>
        </v-row>

        <v-row>
          <v-col cols="4">
            <h3>Status</h3>
          </v-col>
        </v-row>

        <v-row>
          <input
            class="mx-4"
            type="radio"
            :value="'COMPLETE - SUCCESSFUL'"
            v-model="status"
            name="success"
          />
          <label for="success" class="my-2">Completed Successfully</label>
        </v-row>

        <v-row>
          <input
            class="mx-4"
            type="radio"
            :value="'COMPLETE - UNSUCCESSFUL'"
            v-model="status"
            name="unsuccessful"
          />
          <label for="unsuccess" class="my-2">Not Successful</label>
        </v-row>

        <v-row>
          <v-col cols="6">
            <h3>Additional Comments</h3>
            <v-text-field v-model="description" class="mt-5"></v-text-field>
          </v-col>
        </v-row>
        <v-row class="justify-end">
          <v-btn
            @click="end_of_maintenance"
            style="background-color: #5d5fef; color: white"
          >
            End Maintenance
          </v-btn>
        </v-row>
      </v-div>

      <v-div
        v-else-if="
          selectedRecord.status == 'COMPLETE - SUCCESSFUL' ||
          selectedRecord.status == 'COMPLETE - UNSUCCESSFUL'
        "
      >
        <v-row>
          <v-col>
            <h2>Maintenance Record ID: {{ selectedRecord._id }}</h2>
          </v-col>
        </v-row>

        <v-row style="margin-top: 0">
          <v-col>
            <h3>Equipment ID: {{ selectedRecord.equipment.equipment_name }}</h3>
          </v-col>
        </v-row>

        <v-row style="margin-top: 0">
          <v-col>
            <h3>Current status: {{ selectedRecord.status }}</h3>
          </v-col>
        </v-row>

        <v-row style="margin-top: 0">
          <v-col>
            <h3>Scheduled for: {{ selectedRecord.schedule_date }}</h3>
          </v-col>
        </v-row>

        <v-table>
          <thead>
            <tr>
              <th>Parts</th>
              <th>Quantity Used</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="detail in selectedRecord.partlist">
              <td>
                {{ detail.PartName }}
              </td>
              <td>
                {{ detail.Qty }}
              </td>
            </tr>
          </tbody>
        </v-table>
      </v-div>

      <v-col
        cols="6"
        v-if="selectedRecord.status == 'STARTED'"
        style="margin-top: 150px"
      >
        <h3>Request for More Parts</h3>
        <v-table>
          <thead>
            <tr>
              <th class="text-center">Part Name</th>
              <th class="text-center">Quantity</th>
              <th class="text-center" width="20%">Request Qty</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="part in inventory_arr">
              <td class="text-center">{{ part["PartName"] }}</td>
              <td class="text-center">{{ part["Qty"] }}</td>
              <td class="text-center">
                <input
                  type="number"
                  :id="part['_id']"
                  :name="part['PartName']"
                  min="0"
                  value="0"
                  style="border: 1px solid gray; width: 50%"
                  class="text-center"
                  @change="request_parts($event.currentTarget)"
                />
              </td>
            </tr>
          </tbody>
        </v-table>
        <v-row class="justify-end mt-4">
          <v-btn
            style="color: white; background-color: #5d5fef"
            @click="submit_request_parts()"
          >
            Request
          </v-btn>
        </v-row>
      </v-col>
    </v-container>
  </v-container>
</template>

<script>
import navbar from "../components/navbar.vue";
import {
  maintenanceURL,
  inventoryURL,
} from "../../api";
import axios from "axios";

export default {
  components: {
    navbar,
  },

    data() {
        return {
            inventory_arr: [],
            selectedRecord: "",
            description: "",
            status: "",
            requestPartsList: [],
            unusedPartList: [],
            current_date: new Date(),
        }
    },

  methods: {
    update_maintenance_status() {
      console.log("=== START update_maintenance_status() ===");

            var now = this.current_date
            now.setHours(now.getHours() + 8.5);
            now = now.toISOString()
            console.log(now)

            var datetime_arr = now.split("T");
            var date = datetime_arr[0]
            var time = datetime_arr[1]
            var date_arr = date.split("-")
            var time_arr = time.split(":")

      var month = date_arr[1];
      var day = date_arr[2];
      var year = date_arr[0];

      var hour = time_arr[0];
      var minute = time_arr[1];

      console.log(month);
      console.log(day);
      console.log(year);

      var formatted_datetime =
        day + "-" + month + "-" + year + " " + hour + ":" + minute + ":00";
      console.log(formatted_datetime);

      // console.log(datetime)
      let data = {
        equipment: this.selectedRecord.equipment,
        start_datetime: formatted_datetime,
      };
      console.log(data);
      axios
        .put(
          `http://127.0.0.1:8000/api/v1/start_maintenance/${this.selectedRecord._id}`,
          data,
          {
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "*",
              Authorization: "Bearer " + localStorage.token,
            },
          }
        )
        .then((response) => {
          console.log(response.data);
          this.$router.go();
        })
        .catch((error) => {
          console.error(error);
        });
    },

    request_parts(e) {
      let part_id = e.id;
      let reqd_qty = e.value;
      let part_name = e.name;

      let partObj = {
        PartName: part_name,
        Qty: reqd_qty,
        _id: part_id,
      };

      let temp_reqd_part_list = this.requestPartsList;

      let part_idx = temp_reqd_part_list.findIndex(
        (part) => part["_id"] === part_id
      );

      console.log(part_idx);

      // Remove part from part list
      if (reqd_qty === "0") {
        temp_reqd_part_list = temp_reqd_part_list.filter(
          (part) => part["_id"] != part_id
        );
      } else if (part_idx == -1) {
        console.log("Added Parts");
        temp_reqd_part_list.push(partObj);
      } else {
        temp_reqd_part_list[part_idx]["Qty"] = reqd_qty;
      }

      this.requestPartsList = temp_reqd_part_list;
    },

    submit_request_parts() {
      let data = {
        req_partlist: this.requestPartsList,
        partlist: this.selectedRecord.partlist,
      };
      console.log(data);
      axios
        .put(
          `http://127.0.0.1:8000/api/v1/request_parts/${this.selectedRecord._id}`,
          data,
          {
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "*",
              Authorization: "Bearer " + localStorage.token,
            },
          }
        )
        .then((response) => {
          console.log(response.data);
          // alert("Successfully reqd")
          this.$router.go();
        })
        .catch((error) => {});
    },

    update_unused_partlist_array(e) {
      let part_id = e.id;
      let unused_qty = e.value;
      let part_name = e.name;

      let partObj = {
        PartName: part_name,
        Qty: unused_qty,
        _id: part_id,
      };

      let temp_part_list = this.unusedPartList;
      let part_idx = temp_part_list.findIndex(
        (part) => part["_id"] === part_id
      );

      console.log(part_idx);

      if (unused_qty === "0") {
        temp_part_list = temp_part_list.filter(
          (part) => part["_id"] != part_id
        );
      } else if (part_idx == -1) {
        console.log("Added Parts");
        temp_part_list.push(partObj);
      } else {
        temp_part_list[part_idx]["Qty"] = unused_qty;
      }

      this.unusedPartList = temp_part_list;
    },

    async end_of_maintenance(){
      console.log("=== START end_of_maintenance ===")

      var now = this.current_date
      now.setHours(now.getHours() + 8.5);
      now = now.toISOString()

      var datetime_arr = now.split("T");
      var date = datetime_arr[0]
      var time = datetime_arr[1]
      var date_arr = date.split("-")
      var time_arr = time.split(":")

      var month = date_arr[1];
      var day = date_arr[2];
      var year = date_arr[0];

      var hour = time_arr[0];
      var minute = time_arr[1];

      // console.log(month)
      // console.log(day)
      // console.log(year)

      var formatted_datetime =
        day + "-" + month + "-" + year + " " + hour + ":" + minute + ":00";
      console.log(formatted_datetime);

      if (this.status != "") {
        let data = {
        equipment: this.selectedRecord.equipment,
        schedule_date: this.selectedRecord.schedule_date,
        start_datetime: this.selectedRecord.equipment.start_datetime,
        partlist: this.selectedRecord.partlist,
        return_partlist: this.unusedPartList,
        end_datetime: formatted_datetime,
        description: this.description,
        maintenance_status: this.status,
      };

      console.log(data);
      axios
        .put(
          `http://127.0.0.1:8000/api/v1/end_maintenance/${this.selectedRecord._id}`,
          data,
          {
            headers: {
              "Content-Type": "application/json",
              "Access-Control-Allow-Origin": "*",
              Authorization: "Bearer " + localStorage.token,
            },
          }
        )
        .then((response) => {
          console.log(response.data);
          this.$router.go();
        })
        .catch((error) => {
          console.error(error);
        });
      } else {
        alert("Please select a status!")
      }

    },
  },

  async mounted() {
    let inv_result = await axios.get(inventoryURL, {
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
        Authorization: "Bearer " + localStorage.token,
      },
    });
    let maintenance_id = this.$route.params.id;
    let maint_result = await axios.get(`${maintenanceURL}/${maintenance_id}`, {
      headers: {
        "Content-Type": "application/json",
        "Access-Control-Allow-Origin": "*",
        Authorization: "Bearer " + localStorage.token,
      },
    });
    console.log(maint_result);
    this.selectedRecord = maint_result.data.data;
    this.inventory_arr = inv_result.data.data.parts;
    // console.log(this.maintenance_arr)
  },
};
</script>
