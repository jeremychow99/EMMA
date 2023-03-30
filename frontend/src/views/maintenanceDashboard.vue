<template>
    <navbar></navbar>

    <v-container fluid class="mt-2">
        <!-- {{ maintenance_arr }} -->
        <div class="text-center my-3">
            <h1>Maintenance Tasks</h1>
        </div>
        <v-table fixed-header class="border rounded-sm" height="400px">
                        <thead>
                            <tr style="background-color: red;">
                                <th class="text-center">
                                    EquipmentID
                                </th>
                                <th class="text-center">
                                    Location
                                </th>
                                <th class="text-center">
                                    Scheduled Date
                                </th>
                                <th class="text-center">
                                    Select
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr
                                v-for="maintenance in maintenance_arr"
                                :key="maintenance['_id']"
                            >
                                <td class="text-center">{{ maintenance.equipment["equipment_name"] }}</td>
                                <td class="text-center">{{ maintenance.equipment["equipment_location"] }}</td>
                                <td class="text-center">{{ maintenance.schedule_date }}</td>
                                <td class="text-center">
                                    <router-link class="btn btn btn-success" :to="{ name: 'startMaintenance', params: {id: maintenance['_id']} }" style="text-decoration:none">
                                        <v-btn class="btn btn-success">View</v-btn>
                                    </router-link></td>
                            </tr>
                        </tbody>
                    </v-table>
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
            maintenance_arr: [],
        }
    },
    async mounted() {
        // let user_id = sessionStorage.getItem('userId')
        let user_id = '6421c02771a20121b7371d0f'

        let maintenance_result = await axios.get(`${maintenanceURL}/technician/${user_id}`)

        this.maintenance_arr = maintenance_result.data.data
    }
}

</script>