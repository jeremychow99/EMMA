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
        
        <v-row>
            <select name="" id="" style="margin-left: 20px; padding: 0px 10px; border: lightgrey 1px solid;">
                <option value="" disabled selected >Please select</option>
                <option value="volvo">Volvo</option>
                <option value="saab">Saab</option>
                <option value="mercedes">Mercedes</option>
                <option value="audi">Audi</option>
            </select>
        </v-row>

        <v-row style="padding-left: 20px;">
            <v-table>
                <tr>
                    <td>ID</td>
                    <td>Description</td>
                </tr>
            </v-table>
        </v-row>


        <v-btn style="position: fixed; bottom: 5; right: 0;">
            Start
        </v-btn>
    </v-container>

</template>



<script>
import navbar from "../components/navbar.vue"

export default {
    components: {
        navbar
    },

    data() {
        return {
        }
    }
};

</script>