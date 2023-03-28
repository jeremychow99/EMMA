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
        
        <v-row>
            <v-col cols="4">
                <span style="color: #5D5FEF; font-weight: bold">Select a maintenance record:</span>
            </v-col>
        </v-row>
        
        <v-row>
            <v-col style="padding-left: 20px;">
                <h5>Maintenance Record ID</h5>
            </v-col>
        </v-row>
        
        <v-row style="margin-bottom: 10px;">
            <select v-model="selectedRecord" style="margin-left: 20px; padding: 0px 10px; border: lightgrey 1px solid;">
                <option disabled selected >Please select</option>
                <option v-for="maintenance in maintenance_arr" :value="maintenance">{{ maintenance._id }}</option>
            </select>
        </v-row>
        <!-- <h1>{{selectedRecord}}</h1> -->
        

        <v-row style="padding-left: 20px;">
            <v-table v-if="selectedRecord == ''">
                <thead>
                <tr>
                    <th class="text-left">
                    ID
                    </th>
                    <th class="text-left">
                    Status
                    </th>
                </tr>
                </thead>
                <tbody>
                <tr
                v-for="(maintenance, index) in maintenance_arr" :key="maintenance._id" 
                >
                    <td v-if="index < 5">{{ maintenance._id }}</td>
                    <td v-if="maintenance.status && index < 5">{{ maintenance.status }}</td>
                    <td v-else v-if="index <5"> NA </td>
                </tr>
                </tbody>
            </v-table>

            
            <!--DISPLAY PARTS INVENTORY-->
            <v-div v-else>
                <h2>Maintenance Record ID: {{ selectedRecord._id }} </h2>
                <h3>Equipment ID: {{ selectedRecord.equipment_id }}</h3>
                <!-- <h3>Current Status: {{ status }}</h3>
                <h3>Scheduled for: {{ }}</h3>  -->

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
                                {{ detail.ReservedQty }}
                            </td>
                        </tr>
                    </tbody>
                </v-table>
            </v-div>
            
        </v-row>

        <v-row class="justify-end">
            <!-- <v-btn style="background-color: #5D5FEF; color: white; margin-right: 10px;">
                Previous
            </v-btn> -->
            <v-btn @click="update_maintenance_status" style="background-color: #5D5FEF; color: white;">
                Start
            </v-btn>
        </v-row>
    </v-container>

</template>



<script>
import navbar from "../components/navbar.vue"
import { equipmentURL, maintenanceURL, maintenanceControllerURL} from '../../api'
import axios from "axios";

export default {
    components: {
        navbar
    },

    data() {
        return {
            selectedRecord: "",
            status: "NA",
            maintenance_arr: [],


        }
    },

    methods: {
        update_maintenance_status(){
            let datetime = new Date().toLocaleDateString("en-GB")
            console.log(datetime)
            let data = {
                "equipment_id": this.selectedRecord.equipment_id,
                // doesnt work when i put current time but hardcoded is ok
                "start_datetime": datetime
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
                    alert("Success")
                }).catch(error =>{
                    console.error(error)
                    alert("Failed")
                })
        }
    },

    
    async mounted() {
        let maint_result = await axios.get(maintenanceURL)
        this.maintenance_arr = maint_result.data.data.maintenance
        // console.log(this.maintenance_arr)
    }

    
};

</script>