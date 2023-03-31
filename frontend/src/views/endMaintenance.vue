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
                    <select v-model="selectedRecord" style="margin-left: 20px; padding: 0px 10px; border: lightgrey 1px solid;">
                        <option disabled selected >Please select</option>
                        <option v-for="maintenance in maintenance_arr" :value="maintenance">{{ maintenance._id }}</option>
                    </select>
            </v-col>
        </v-row>

        <v-row>
            <v-col cols="4">
                <h2>Maintenance Record ID: {{ selectedRecord._id }}</h2>
            </v-col>
        </v-row>

        <v-row style="margin-top: 0">
            <v-col cols="4">
                <h3>Equipment ID: {{ selectedRecord.equipment_id }} </h3>
            </v-col>
        </v-row>

        <v-row style="margin-top: 0">
            <v-col cols="4">
                <h3>Current status: {{ selectedRecord.status }} </h3>
            </v-col>
        </v-row>

        <v-row style="margin-top: 0">
            <v-col cols="4">
                <h3>Scheduled for: {{ selectedRecord.last_maintained }}</h3>
            </v-col>
        </v-row>

        <v-row>
            <v-col cols="4">
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
                            <input type="text" placeholder="Number of unused parts">
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
            <input type="radio" v-model="status">
            <v-radio-group v-model="active">
                <v-radio label="Successful"></v-radio>
                <v-radio label="Not successful"></v-radio>
            </v-radio-group>
        </v-row>

        <v-row>
            <v-col cols="4">
                <h3>Additional Comments</h3>
                <v-text-field v-model="description"></v-text-field>
            </v-col>
        </v-row>


        <v-row class="justify-end">
            <v-btn style="background-color: #5D5FEF; color: white; margin-right: 10px;">
                Previous
            </v-btn>
            <v-btn style="background-color: #5D5FEF; color: white;" @click="end_of_maintenance()">
                End
            </v-btn>
        </v-row>
    </v-container>
</template>

<script>
import navbar from "../components/navbar.vue"
import { equipmentURL, inventoryURL , maintenanceURL} from '../../api'
import axios from "axios";


export default {
    components: {
        navbar
    },

    data(){
        return {
            status: "",
            selectedRecord: "",
            description: "",
            maintenance_arr: [],

        }
    }, 
    methods: {
        end_of_maintenance(){
            let data = {
                "equipment_id": this.selectedRecord.equipment_id,
                "schedule_date": this.selectedRecord.schedule_date,
                "start_datetime": this.selectedRecord.start_datetime,
                "partlist": this.selectedRecord.partlist,
                "return_partlist": [
                    {

                    }
                ],
                "end_datetime": "29/03/2023",
                "description": this.description,
                "maintenance_status": "COMPLETE - SUCCESSFUL"
            }
            axios.put(`${maintenanceControllerURL}/end_maintenance/${this.selectedRecord._id}`, data, {
                    headers: {
                        "Content-Type": "application/json",
                        "Access-Control-Allow-Origin": "*",
                        Authorization: "Bearer " + localStorage.token,
                    }
                }).then(response => {
                    console.log(response.data)
                    alert("Successfully ended maintenance")
                }).catch(error =>{
                    console.error(error)
                    alert("Failed to end maintenance")
                })
        }
    },

    async mounted() {
        let eqp_result = await axios.get(equipmentURL)
        let inv_result = await axios.get(inventoryURL)
        let maint_result = await axios.get(maintenanceURL)

        this.maintenance_arr = maint_result.data.data.maintenance
        this.equipment_arr = eqp_result.data.data.equipment
        this.inventory_arr = inv_result.data.data.parts
        console.log(this.maintenance_arr)
    }
}

</script>