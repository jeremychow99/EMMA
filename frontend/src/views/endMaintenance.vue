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
            <v-col cols="4">
                <h2>Maintenance Record ID: {{ recordId }}</h2>
            </v-col>
        </v-row>

        <v-row style="margin-top: 0">
            <v-col cols="4">
                <h3>Equipment ID: {{ eqId }} </h3>
            </v-col>
        </v-row>

        <v-row style="margin-top: 0">
            <v-col cols="4">
                <h3>Current status: {{ status }} </h3>
            </v-col>
        </v-row>

        <v-row style="margin-top: 0">
            <v-col cols="4">
                <h3>Scheduled for: {{ date }}</h3>
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
                        ID
                        </th>
                        <th class="text-left">
                        Description
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr
                     v-for="parts in inventory"
                     :key="parts"
                    >
                        <td>{{ parts.part }}</td>
                        <td>{{ parts.qty }}</td>
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
            <v-radio-group>
                <v-radio label="Successful"></v-radio>
                <v-radio label="Not successful"></v-radio>
            </v-radio-group>
        </v-row>

        <v-row>
            <v-col cols="4">
                <h3>Additional Comments</h3>
                <v-text-field></v-text-field>
            </v-col>
        </v-row>


        <v-row class="justify-end">
            <v-btn style="background-color: #5D5FEF; color: white; margin-right: 10px;">
                Previous
            </v-btn>
            <v-btn style="background-color: #5D5FEF; color: white;">
                Start
            </v-btn>
        </v-row>
    </v-container>
</template>

<script>
import navbar from "../components/navbar.vue"

export default {
    components: {
        navbar
    },

    data(){
        return {
            eqId: "1234",
            recordId: "#12345H5",
            status: "In progress",
            date: "24 Mar 2023, 4:44pm",
            inventory: [
                {
                    part: "Screw",
                    qty: 1
                },
                {
                    part: "Battery",
                    qty: 4
                }
            ]
        }
    }
}

</script>