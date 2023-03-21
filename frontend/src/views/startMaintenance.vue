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
            <select name="" id="" style="margin-left: 20px; padding: 0px 10px; border: lightgrey 1px solid;">
                <option value="" disabled selected >Please select</option>
                <v-for>
                    <option value=""></option>
                </v-for>
            </select>
        </v-row>

        <v-row style="padding-left: 20px;">
            <v-table>
                <thead>
                <tr>
                    <th class="text-left">
                    ID
                    </th>
                    <th class="text-left">
                    Description
                    </th>
                </tr>
                </thead>
                <tbody>
                <tr
                    v-for="maintenance in maintenance_arr"
                    :key="maintenance.id"
                >
                    <td>{{ maintenance.id }}</td>
                    <td>{{ maintenance.status }}</td>
                </tr>
                </tbody>
            </v-table>
        </v-row>

        <v-btn style="position: fixed; bottom: 0; right: 0;">
            Start
        </v-btn>
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

    data() {
        return {
            maintenance_arr: [],
            equipment_arr: [],
            inventory_arr: [],
        }
    },

    methods: {

    },

    async mounted() {
        let eqp_result = await axios.get(equipmentURL)
        let inv_result = await axios.get(inventoryURL)
        let maint_result = await axios.get(maintenanceURL)

        this.maintenance_arr = maint_result.data.data.maintenance
        this.equipment_arr = eqp_result.data.data.equipment
        this.inventory_arr = inv_result.data.data.parts
    }
};

</script>