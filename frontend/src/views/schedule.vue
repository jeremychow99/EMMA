<template>
    <navbar></navbar>

    <v-container>

        <v-row>
            <v-col cols="4">
                <h2 class="text-start mt-3">Schedule</h2>
                <h2 class="text-start mb-3">Maintenance</h2>
            </v-col>
            <v-col cols="6" style="background-color: lavender;">
                <v-row class="mt-0 mx-2">
                    <v-col cols="5" style="padding:8px 16px">
                        <p class="text-right font-weight-regular">EquipmentID: </p>
                    </v-col>
                    <v-col cols="7" style="padding:8px 16px">
                        <p>{{ equipment_id }}</p>
                    </v-col>
                </v-row>
                <v-row class="mt-0 mx-2">
                    <v-col cols="5" style="padding:8px 16px">
                        <p class="text-right font-weight-regular">Parts: </p>
                    </v-col>
                    <v-col cols="7" style="padding:8px 16px">
                        <p v-if="Object.keys(part_list).length != 0">Reserved</p>
                        <p v-else>No Parts Reserved</p>
                    </v-col>
                </v-row>
                <v-row class="mt-0 mx-2">
                    <v-col cols="5" style="padding: 8px 16px 8px 8px">
                        <p class="text-right font-weight-regular">Maintenance Datetime: </p>
                    </v-col>
                    <v-col cols="7" style="padding:8px 16px">
                        <input type="datetime-local" id="maintenance-datetime"
                        name="maintenance-datetime" v-model="maintenance_datetime">
                    </v-col>
                </v-row>
            </v-col>
            <v-col cols="2" class="position-relative" style="background-color: lavender;">
                <v-btn class="btn btn-success position-absolute" style="right: 5px; bottom: 5px;" @click="ScheduleMaintenance" :disabled="!formValid">Submit</v-btn>
            </v-col>
        </v-row>

        <!-- Testing Purpose -->
        <!-- <div>
            {{ part_list }}
        </div> -->
        

        <v-tabs class="mx-auto mt-3" align-tabs="center" v-model="tab">
            <v-tab value="equipment">
                Equipment
            </v-tab>
            <v-tab value="parts">
                Parts
            </v-tab>
        </v-tabs>

        <v-window v-model="tab">

            <v-window-item key="equipment" value="equipment">
                <v-container class="rounded-lg" style="background-color: lightblue; padding: 5px">
                
                    <v-text-field
                        :loading="loading"
                        density="compact"
                        variant="solo"
                        label="Search equipment name"
                        single-line
                        hide-details
                        @click:append-inner="onClick"
                        class="my-2"
                        style="margin:auto; width: 50%;"
                    ></v-text-field>
                   
                        

                    <v-table fixed-header height="400px">
                        <thead>
                            <tr>
                                <th class="text-center">
                                    Name
                                </th>
                                <th class="text-center">
                                    Status
                                </th>
                                <th class="text-center">
                                    Select
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr
                                v-for="eqp in equipment_arr"
                                :key="eqp['_id']"
                            >
                                <td class="text-center">{{ eqp.equipment_name }}</td>
                                <td class="text-center">{{ eqp.equipment_status }}</td>
                                <td class="text-center"><input type="radio" :id="eqp['_id']" :value="eqp['_id']" name="eqp_radios" v-model="equipment_id"></td>
                            </tr>
                        </tbody>
                    </v-table>
                </v-container>
            </v-window-item>

            <v-window-item key="parts" value="parts">
                <v-container class="rounded-lg" style="background-color: lightgray; padding: 5px">
                
                    <v-text-field
                        :loading="loading"
                        density="compact"
                        variant="solo"
                        label="Search part name"
                        single-line
                        hide-details
                        @click:append-inner="onClick"
                        class="my-2"
                        style="margin:auto; width: 50%;"
                    ></v-text-field>
                
                        

                    <v-table fixed-header height="400px">
                        <thead>
                            <tr>
                                <th class="text-center">
                                    Part Name
                                </th>
                                <th class="text-center">
                                    Quantity
                                </th>
                                <th class="text-center" width="20%">
                                    Reserve Qty
                                </th>
                            </tr>
                        </thead>
                        <tbody>
                            <tr v-for="part in inventory_arr">
                                <td class="text-center">{{ part['PartName'] }}</td>
                                <td class="text-center">{{ part['Qty'] }}</td>
                                <td class="text-center">
                                    <!-- <v-text-field :id="part['_id']" type="number" width="50" value="0" min="0"></v-text-field> -->
                                    <input type="number" :id="part['_id']" min="0" value="0" style="border: 1px solid gray; width:50%;" class="text-center" @change="UpdatePartList($event.currentTarget)">
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
import navbar from "../components/navbar.vue"
import { equipmentURL, inventoryURL, maintenanceControllerURL } from '../../api'
import axios from "axios"

export default {
    components: {
        navbar
    },

    data() {
        return {
            tab: null,
            date: null,
            equipment_arr: [],
            inventory_arr: [],

            // Form Inputs
            equipment_id: "",
            part_list: {},
            maintenance_datetime: "",
        }
    },

    computed: {
        formValid() {
            return (this.equipment_id != "" && this.maintenance_datetime != "")
        }
    },

    methods: {
        UpdatePartList(e) {
            // console.log(e)
            // console.log(e.id)
            // console.log(e.value)
            let part_id = e.id
            let reserve_qty = e.value
            let temp_part_list = this.part_list

            // Remove part from part list
            if (reserve_qty === "0" && part_id in temp_part_list) {
                delete temp_part_list[part_id]
            } else {
                // console.log("Added Parts")
                temp_part_list[part_id] = reserve_qty
            }
            
            this.part_list = temp_part_list
        },
        
        ScheduleMaintenance() {
            // console.log(this.equipment_id)
            // console.log(this.part_list)
            // console.log(this.maintenance_datetime)

            var datetime_arr = this.maintenance_datetime.split("T");
            var date = datetime_arr[0]
            var time = datetime_arr[1]
            var date_arr = date.split("-")

            var month = date_arr[1]
            var day = date_arr[2]
            var year = date_arr[0]

            // console.log(month)
            // console.log(day)
            // console.log(year)

            var formatted_datetime = day + '-' + month + '-' + year + " " + time + ":00"
            console.log(formatted_datetime)


            if (this.equipment_id != "" && this.maintenance_datetime != "") {
                console.log("Submitted schedule maintenance request")
                let data = {
                    equipment_id: this.equipment_id,
                    schedule_datetime: formatted_datetime,
                    partlist: this.part_list
                    // Need to include creator_id from sessionStorage
                }

                // console.log(this.part_list)
                // console.log(data)

                axios.post(`${maintenanceControllerURL}/schedule_maintenance`, data).then(response => {
                    console.log(response.data)
                    

                    // Reset Inputs
                    this.equipment_id = ""
                    this.part_list = {}
                    this.maintenance_datetime= ""
                }).catch(error => {

                })
            }
        }
    },
    async mounted() {
        let eqp_result = await axios.get(equipmentURL)
        let inv_result = await axios.get(inventoryURL)

        this.equipment_arr = eqp_result.data.data.equipment
        this.inventory_arr = inv_result.data.data.parts
    }
};
</script>