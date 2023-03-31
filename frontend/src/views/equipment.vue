<template>
    <navbar></navbar>
    
    <v-container>
      <div class="text-h4 text-center font-weight-black my-5">Equipment</div>
    </v-container>
    
    <v-container>
        <v-row>
            <v-col class="text-h6" cols="4">
                All Equipment Overview
            </v-col>
            <v-col cols="4">
            </v-col>
            <v-col style="margin-bottom: 10px;">
                <!-- <v-btn @click="show_modal = true" style="background-color: #5D5FEF; color: white;">
                    Add Equipment
                </v-btn> -->
                <v-btn
        color="primary"
        >
        Add equipment

        <v-dialog
            v-model="dialog"
            activator="parent"
            width="auto"
            height="auto"
        >
            <v-card>
                <v-card-text>
                    <v-row style="margin-bottom: 20px">
                        <h3>New Equipment</h3>
                    </v-row>
                    <v-div>
                        <form @submit.prevent="handleSubmit">
                            <v-row style="margin-bottom: 5px">
                                <label for="name">Name: </label>      
                                <input type="text" v-model="new_equipment_name">
                            </v-row>
                            <v-row style="margin-bottom: 5px">
                                <label for="location">Location:</label>
                                <input type="text" v-model="new_equipment_location">
                            </v-row>
                            <!-- <v-row style="margin-bottom: 5px">
                                <label for="status">Status:</label>
                                <input type="text" v-model="new_equipment_status">
                            </v-row> -->
                            <!-- <v-row style="margin-bottom: 5px">
                                <label for="last maintained">Last Maintained:</label>
                                <input type="text" v-model="new_equipment_last_maintained">
                            </v-row> -->
                            <v-card-actions style="justify-content: center">
                                <v-btn color="primary" @click="dialog = false">Close</v-btn>
                                <v-btn color="primary" type="submit">Add</v-btn>
                            </v-card-actions>
                        </form>
                    </v-div>
                </v-card-text>
            </v-card>
        </v-dialog>
    </v-btn>
                
            </v-col>
        </v-row>

            <v-table fixed-header height="300px">
                <thead>
                    <tr>
                        <th>
                            Equipment Name
                        </th>
                        <th>
                            Equipment Location
                        </th>                        
                        <th>
                            Last Maintained
                        </th>                        
                        <th>
                            Status
                        </th>
                    </tr>
                </thead>
                <tbody>
                    <tr v-for="equipment of eq_arr">
                        <td>
                            {{ equipment.equipment_name }}
                        </td>
                        <td>
                            {{ equipment.equipment_location }}
                        </td>
                        <td>
                            {{ equipment.last_maintained }}
                        </td>
                        <td>
                            {{ equipment.equipment_status }}
                        </td>
                    </tr>
                </tbody>
            </v-table>
    </v-container>
    
</template>

<script>
import navbar from "../components/navbar.vue"
import { equipmentURL  } from '../../api'
import axios from "axios";
import modal from "../components/modal.vue"

export default {
    components: {
        navbar,
        modal
    },

    data(){ 
        return {
            eq_arr: [],
            dialog: false,
            new_equipment_name: "",
            new_equipment_location: "",
            new_equipment_status: "",
            new_equipment_last_maintained: ""
            
        }
    },

    methods: {
        async handleSubmit(){
            var current_date = new Date().toISOString()
            var datetime_arr = current_date.split("T");
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
                "equipment_name": this.new_equipment_name,
                "equipment_location": this.new_equipment_location,
                "equipment_status": "Healthy",
                "last_maintained": formatted_datetime,
            }
            console.log(data)
            axios.post(`${equipmentURL}`, data, {
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

        showModalForm() {
            this.show_modal = true;
        },
    },

    async mounted() {
        let eqp_result = await axios.get(equipmentURL)
        this.eq_arr = eqp_result.data.data.equipment
        console.log(this.eq_arr)
    }

}

</script>