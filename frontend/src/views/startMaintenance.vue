<template>
    <navbar></navbar>

    <v-container>
        <v-row>
            <v-col cols="4">
                <h2 class="text-start mt-3">View</h2>
                <h2 class="text-start mb-3">Maintenance</h2>
            </v-col>
            <v-col cols="8">
            </v-col>
        </v-row>
        
            <v-div v-if="selectedRecord.status == 'SCHEDULED' ">
                <v-row>
                    <v-col>
                        <h2>Maintenance Record ID: {{ selectedRecord._id }}</h2>
                    </v-col>
                </v-row>

                <v-row style="margin-top: 0">
                    <v-col>
                        <h3>Equipment ID: {{ selectedRecord.equipment.equipment_name}} </h3>
                    </v-col>
                </v-row>

                <v-row style="margin-top: 0">
                    <v-col>
                        <h3>Current status: {{ selectedRecord.status }} </h3>
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
                            <th>
                                Parts
                            </th>
                            <th>
                                Quantity Reserved
                            </th>
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
                    <v-btn @click="update_maintenance_status" style="background-color: #5D5FEF; color: white;">
                        Start
                    </v-btn>
                </v-row>  
            </v-div>

            <v-container v-else-if="selectedRecord.status =='STARTED'">
                <v-row>
                    <v-col>
                        <h2>Maintenance Record ID: {{ selectedRecord._id }}</h2>
                    </v-col>
                </v-row>

                <v-row style="margin-top: 0">
                    <v-col>
                        <h3>Equipment ID: {{ selectedRecord.equipment.equipment_name}} </h3>
                    </v-col>
                </v-row>

                <v-row style="margin-top: 0">
                    <v-col>
                        <h3>Current status: {{ selectedRecord.status }} </h3>
                    </v-col>
                </v-row>

                <v-row style="margin-top: 0">
                    <v-col>
                        <h3>Scheduled for: {{ selectedRecord.last_maintained }}</h3>
                    </v-col>
                </v-row>

                <v-row>
                    <v-col>
                        <h3>Unused Parts</h3>
                    </v-col>
                </v-row>

                <v-row>
                    <v-table>
                        <thead>
                            <tr>
                                <th class="text-left">
                                Part Name
                                </th>
                                <th class="text-left">
                                Qty
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr
                            v-for="parts in selectedRecord.partlist"
                            :key="parts"
                            >
                                <td>{{ parts.PartName }}</td>
                                <td>
                                    <input type="text" :id="parts['_id']" :name="parts.PartName" placeholder="Number of unused parts" @change="update_unused_partlist_array($event.currentTarget)">
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
                    <input type="radio" :value="'COMPLETE - SUCCESSFUL'" v-model="status"> Completed Successfully
                </v-row>
                
                <v-row>
                    <input type="radio" :value="'COMPLETE - UNSUCCESSFUL'" v-model="status"> Not Successful
                </v-row>

                <v-row>
                    <v-col cols="4">
                        <h3>Additional Comments</h3>
                        <v-text-field v-model="description"></v-text-field>
                    </v-col>
                </v-row>
                <v-row class="justify-end">
                    <v-btn @click="end_of_maintenance" style="background-color: #5D5FEF; color: white;">
                        End
                    </v-btn>
                </v-row>  
            </v-container>

            <v-div v-else>
                
            </v-div>
    </v-container>

</template>



<script>
import navbar from "../components/navbar.vue"
import { maintenanceURL, maintenanceControllerURL } from '../../api'
import axios from "axios";

export default {
    components: {
        navbar
    },

    data() {
        return {
            selectedRecord: "",
            description: "",
            status: "",
            unusedPartList: [],
            current_date: new Date().toISOString(),
        }
    },

    methods: {
        update_maintenance_status(){
            console.log("=== START update_maintenance_status() ===")

            var datetime_arr = this.current_date.split("T");
            var date = datetime_arr[0]
            var time = datetime_arr[1]
            var date_arr = date.split("-")
            var time_arr = time.split(":")

            var month = date_arr[1]
            var day = date_arr[2]
            var year = date_arr[0]

            var hour = time_arr[0]
            var minute = time_arr[1]

            console.log(month)
            console.log(day)
            console.log(year)

            var formatted_datetime = day + '-' + month + '-' + year + " " + hour + ":" + minute + ":00"
            console.log(formatted_datetime)

            // console.log(datetime)
            let data = {
                "equipment": this.selectedRecord.equipment,
                // doesnt work when i put current time but hardcoded is ok
                "start_datetime": formatted_datetime
            }
            console.log(data);
            axios.put(`${maintenanceControllerURL}/start_maintenance/${this.selectedRecord._id}`, data, {
                    headers: {
                        "Content-Type": "application/json",
                        "Access-Control-Allow-Origin": "*",
                        Authorization: "Bearer " + localStorage.token,
                    }
                }).then(response => {
                    console.log(response.data)
                    this.$router.go()
                }).catch(error =>{
                    console.error(error)
                })
        }, 

        update_unused_partlist_array(e){
            console.log(e)
            let part_id = e.id
            let unused_qty = e.value
            let part_name = e.name
            
            let partObj = {
                "PartName": part_name,
                "Qty": unused_qty,
                "_id": part_id
            }

            let temp_part_list = this.unusedPartList
            let part_idx = temp_part_list.findIndex(part => part["_id"] === part_id)

            console.log(part_idx)

            if (unused_qty === "0") {
                temp_part_list = temp_part_list.filter(part => part["_id"] != part_id)
            } else if (part_idx == -1) {
                console.log("Added Parts")
                temp_part_list.push(partObj)
            } else {
                temp_part_list[part_idx]["Qty"] = reserve_qty
            }

            this.unusedPartList = temp_part_list
        }, 

        async end_of_maintenance(){
            console.log("=== START end_of_maintenance ===")
            var datetime_arr = this.current_date.split("T");
            var date = datetime_arr[0]
            var time = datetime_arr[1]
            var date_arr = date.split("-")
            var time_arr = time.split(":")

            var month = date_arr[1]
            var day = date_arr[2]
            var year = date_arr[0]

            var hour = time_arr[0]
            var minute = time_arr[1]

            console.log(month)
            console.log(day)
            console.log(year)

            var formatted_datetime = day + '-' + month + '-' + year + " " + hour + ":" + minute + ":00"
            console.log(formatted_datetime)

            let data = {
                "equipment": this.selectedRecord.equipment,
                "schedule_date": this.selectedRecord.schedule_date,
                "start_datetime": this.selectedRecord.equipment.start_datetime,
                "partlist": this.selectedRecord.partlist,
                "return_partlist": this.unusedPartList,
                "end_datetime": formatted_datetime,
                "description": this.description,
                "maintenance_status": this.status
            }
            console.log(data)
            axios.put(`${maintenanceControllerURL}/end_maintenance/${this.selectedRecord._id}`, data, {
                    headers: {
                        "Content-Type": "application/json",
                        "Access-Control-Allow-Origin": "*",
                        Authorization: "Bearer " + localStorage.token,
                    }
                }).then(response => {
                    console.log(response.data)
                    alert("Successfully ended maintenance")
                    this.$router.go()
                }).catch(error =>{
                    console.error(error)
                    alert("Failed to end maintenance")
                })
        }
    },

    
    async mounted() {
        let maintenance_id = this.$route.params.id
        let maint_result = await axios.get(`${maintenanceURL}/${maintenance_id}`)
        console.log(maint_result)
        this.selectedRecord = maint_result.data.data
        // console.log(this.maintenance_arr)
    }

    
};

</script>